package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
	"github.com/tranChiVi1412/goraemon/internal/domain/project"
	"github.com/tranChiVi1412/goraemon/internal/domain/registry"
	"github.com/tranChiVi1412/goraemon/internal/domain/service"
	"github.com/tranChiVi1412/goraemon/internal/domain/template"
	"github.com/tranChiVi1412/goraemon/internal/infrastructure/persistence"
	"github.com/tranChiVi1412/goraemon/internal/usecase"
)

func main() {
	// Create a directory to save metadata if not exist
	repoDir := ".goraemon"
	if err := os.MkdirAll(repoDir, 0755); err != nil {
		log.Fatalf("Failed to create repository directory: %v", err)
	}

	// Initialize repositories
	projectRepo := persistence.NewProjectFileRepo(repoDir)
	moduleRepo := persistence.NewModuleFileRepo(repoDir)
	moduleRegistry := registry.NewModuleRegistry()

	// Register default modules
	defaultModules := []struct {
		name    string
		modType string
		deps    []string
	}{
		{
			name:    "mysql",
			modType: string(module.ModuleTypeDatabase),
			deps:    []string{},
		},
		{
			name:    "redis",
			modType: string(module.ModuleTypeCache),
			deps:    []string{},
		},
	}

	for _, m := range defaultModules {
		if err := moduleRegistry.Register(&module.Module{
			Name:         m.name,
			Type:         module.ModuleType(m.modType),
			Dependencies: m.deps,
		}); err != nil {
			log.Fatalf("Failed to register module %s: %v", m.name, err)
		}
	}

	validator := service.NewProjectValidator()
	templateRenderer := template.NewTemplateRenderer()

	// Initialize use cases
	initProjectUC, err := usecase.NewInitProjectInteractor(
		projectRepo,
		moduleRepo,
		&moduleRegistry,
		validator,
		templateRenderer,
	)
	if err != nil {
		log.Fatalf("Failed to initialize project interactor: %v", err)
	}

	input := usecase.InitProjectInput{
		Name:       "my-project",
		ModulePath: "github.com/tranChiVi1412/user-service",
		Transports: []project.Transport{project.TransportREST},
		Modules: []*module.Module{
			{
				Name: "mysql",
				Type: module.ModuleTypeDatabase,
			},
			{
				Name: "redis",
				Type: module.ModuleTypeCache,
			},
		},
	}

	output, err := initProjectUC.Execute(context.Background(), input)
	if err != nil {
		log.Fatalf("Failed to initialize project: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Project initialized successfully: %s\n", output.ProjectID)
}
