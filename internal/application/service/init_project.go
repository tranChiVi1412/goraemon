package service

import (
	"context"
	"fmt"
	"time"

	"github.com/tranChiVi1412/goraemon/internal/application/port/in"
	"github.com/tranChiVi1412/goraemon/internal/application/port/out"
	"github.com/tranChiVi1412/goraemon/internal/domain/module"
	"github.com/tranChiVi1412/goraemon/internal/domain/project"
	"github.com/tranChiVi1412/goraemon/internal/domain/registry"
	"github.com/tranChiVi1412/goraemon/internal/domain/service"
	"github.com/tranChiVi1412/goraemon/internal/domain/template"
)

// initProjectUseCase implements InitProjectUseCase
type initProjectUseCase struct {
	projectRepo      out.ProjectRepositoryPort
	moduleRepo       out.ModuleRepositoryPort
	moduleRegistry   registry.ModuleRegistry
	validator        service.ProjectService
	templateRenderer template.TemplateRenderer
}

// NewInitProjectUseCase creates a new init project use case
func NewInitProjectUseCase(
	projectRepo out.ProjectRepositoryPort,
	moduleRepo out.ModuleRepositoryPort,
	moduleRegistry registry.ModuleRegistry,
	validator service.ProjectService,
	templateRenderer template.TemplateRenderer,
) in.InitProjectUseCase {
	return &initProjectUseCase{
		projectRepo:      projectRepo,
		moduleRepo:       moduleRepo,
		moduleRegistry:   moduleRegistry,
		validator:        validator,
		templateRenderer: templateRenderer,
	}
}

// Execute executes the use case
func (uc *initProjectUseCase) Execute(ctx context.Context, input in.InitProjectInput) (in.InitProjectOutput, error) {
	// Check if project exists
	exists, err := uc.projectRepo.Exists(ctx, input.Name)
	if err != nil {
		return in.InitProjectOutput{}, fmt.Errorf("failed to check project existence: %w", err)
	}
	if exists {
		return in.InitProjectOutput{}, fmt.Errorf("project %s already exists", input.Name)
	}

	// Create project
	project := project.NewProject(input.Name, input.ModulePath, input.Transports, input.Modules, time.Now())

	// Validate project
	if err := uc.validator.Validate(project, uc.moduleRegistry); err != nil {
		return in.InitProjectOutput{}, fmt.Errorf("failed to validate project: %w", err)
	}

	// Save project
	if err := uc.projectRepo.Save(ctx, project); err != nil {
		return in.InitProjectOutput{}, fmt.Errorf("failed to save project: %w", err)
	}

	// Render templates
	moduleTypes := make(map[module.ModuleType]bool)
	for _, module := range project.Modules {
		moduleTypes[module.Type] = true
	}

	for modType := range moduleTypes {
		templates, err := uc.templateRenderer.ListTemplates(modType)
		if err != nil {
			return in.InitProjectOutput{}, fmt.Errorf("failed to list templates for module type %s: %w", modType, err)
		}

		for _, tmpl := range templates {
			data := map[string]interface{}{
				"ProjectName": project.Name,
				"ModulePath":  project.ModulePath,
			}

			if _, err := uc.templateRenderer.Render(tmpl, data); err != nil {
				return in.InitProjectOutput{}, fmt.Errorf("failed to render template %s: %w", tmpl.Name, err)
			}
		}
	}

	return in.InitProjectOutput{
		ProjectID: project.ID,
		Message:   fmt.Sprintf("Project %s initialized successfully", project.Name),
	}, nil
}
