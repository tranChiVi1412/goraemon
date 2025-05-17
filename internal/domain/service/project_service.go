package service

import (
	"fmt"

	"github.com/tranChiVi1412/goraemon/internal/domain/project"
	"github.com/tranChiVi1412/goraemon/internal/domain/registry"
)

// ProjectService defines the interface for project-related business logic
type ProjectService interface {
	// Validate validates a project
	Validate(project *project.Project, moduleRegistry registry.ModuleRegistry) error
}

// projectService implements ProjectService
type projectService struct{}

// NewProjectService creates a new project service
func NewProjectService() ProjectService {
	return &projectService{}
}

// Validate validates a project
func (s *projectService) Validate(project *project.Project, moduleRegistry registry.ModuleRegistry) error {
	// Validate project name
	if project.Name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	// Validate project path
	if project.ModulePath == "" {
		return fmt.Errorf("project path cannot be empty")
	}

	// Validate transports
	if len(project.Transports) == 0 {
		return fmt.Errorf("project must have at least one transport")
	}

	// Validate modules
	if len(project.Modules) == 0 {
		return fmt.Errorf("project must have at least one module")
	}

	// Validate each module
	for _, module := range project.Modules {
		// Check if module exists in registry
		regModule, err := moduleRegistry.Get(module.Name)
		if err != nil {
			return fmt.Errorf("module %s not found in registry: %w", module.Name, err)
		}

		// Check module type
		if module.Type != regModule.Type {
			return fmt.Errorf("module %s has invalid type: expected %s, got %s", module.Name, regModule.Type, module.Type)
		}

		// Check dependencies
		for _, dep := range module.Dependencies {
			if !project.HasModule(dep) {
				return fmt.Errorf("module %s depends on %s, but it is not added to the project", module.Name, dep)
			}
		}
	}

	return nil
}
