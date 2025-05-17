package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
	"github.com/tranChiVi1412/goraemon/internal/domain/project"
	"github.com/tranChiVi1412/goraemon/internal/domain/registry"
	"github.com/tranChiVi1412/goraemon/internal/domain/service"
	"github.com/tranChiVi1412/goraemon/internal/domain/template"
)

type initProjectImpl struct {
	ProjectRepo      project.ProjectRepository
	ModuleRepo       module.ModuleRepository
	ModuleRegistry   registry.ModuleRegistry
	Validator        service.ProjectValidator
	TemplateRenderer template.TemplateRenderer
}

func NewInitProjectInteractor(
	projectRepo project.ProjectRepository,
	moduleRepo module.ModuleRepository,
	moduleRegistry *registry.ModuleRegistry,
	validator service.ProjectValidator,
	templateRenderer template.TemplateRenderer,
) (*initProjectImpl, error) {
	if projectRepo == nil {
		return nil, fmt.Errorf("project repository is required")
	}
	if moduleRepo == nil {
		return nil, fmt.Errorf("module repository is required")
	}
	if moduleRegistry == nil {
		return nil, fmt.Errorf("module registry is required")
	}
	if validator == nil {
		return nil, fmt.Errorf("validator is required")
	}
	if templateRenderer == nil {
		return nil, fmt.Errorf("template renderer is required")
	}

	return &initProjectImpl{
		ProjectRepo:      projectRepo,
		ModuleRepo:       moduleRepo,
		ModuleRegistry:   *moduleRegistry,
		Validator:        validator,
		TemplateRenderer: templateRenderer,
	}, nil
}

func (uc *initProjectImpl) Execute(ctx context.Context, input InitProjectInput) (InitProjectOutput, error) {
	// 1. Validate project name is exist
	exists, err := uc.ProjectRepo.Exists(ctx, input.Name)
	if err != nil {
		return InitProjectOutput{}, fmt.Errorf("failed to check project existence: %w", err)
	}
	if exists {
		return InitProjectOutput{}, fmt.Errorf("project %s already exists", input.Name)
	}

	// 2. Create project object
	prj := project.NewProject(input.Name, input.ModulePath, input.Transports, input.Modules, time.Now())

	// 3. Get module from repository and add into project
	for _, mod := range input.Modules {
		// Get module from registry
		regMod, err := uc.ModuleRegistry.Get(mod.Name)
		if err != nil {
			return InitProjectOutput{}, fmt.Errorf("module %s not found in registry: %w", mod.Name, err)
		}

		// Convert registry.Module to module.Module
		mod := &module.Module{
			Name:         regMod.Name,
			Type:         regMod.Type,
			Dependencies: regMod.Dependencies,
			Enabled:      true,
		}

		prj.AddModule(mod)
	}

	// 4. Validate project (domain service)
	if err := uc.Validator.Validate(prj, &uc.ModuleRegistry); err != nil {
		return InitProjectOutput{}, fmt.Errorf("validation failed: %w", err)
	}

	// 5. Save project
	if err := uc.ProjectRepo.Save(ctx, prj); err != nil {
		return InitProjectOutput{}, fmt.Errorf("failed to save project: %w", err)
	}

	// 6. Render templates
	// Get all unique module types from the project
	moduleTypes := make(map[module.ModuleType]struct{})
	for _, mod := range prj.Modules {
		moduleTypes[mod.Type] = struct{}{}
	}

	// Convert modules to map for template data
	moduleData := make(map[string]interface{})
	for _, mod := range prj.Modules {
		moduleData[mod.Name] = mod
	}

	// Render templates for each module type
	for modType := range moduleTypes {
		templates, err := uc.TemplateRenderer.ListTemplates(modType)
		if err != nil {
			return InitProjectOutput{}, fmt.Errorf("failed to list templates for module type %s: %w", modType, err)
		}

		// Render each template
		for _, tmpl := range templates {
			_, err := uc.TemplateRenderer.Render(tmpl, moduleData)
			if err != nil {
				return InitProjectOutput{}, fmt.Errorf("failed to render template %s: %w", tmpl.Name, err)
			}
		}
	}

	// 7. Return success
	return InitProjectOutput{
		ProjectID: prj.ID,
		Message:   "Project created successfully",
	}, nil
}
