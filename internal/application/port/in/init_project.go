package in

import (
	"context"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
	"github.com/tranChiVi1412/goraemon/internal/domain/project"
)

// InitProjectInput represents the input for initializing a project
type InitProjectInput struct {
	Name       string
	ModulePath string
	Transports []project.Transport
	Modules    []*module.Module
}

// InitProjectOutput represents the output of initializing a project
type InitProjectOutput struct {
	ProjectID string
	Message   string
}

// InitProjectUseCase defines the interface for initializing a project
type InitProjectUseCase interface {
	// Execute executes the use case
	Execute(ctx context.Context, input InitProjectInput) (InitProjectOutput, error)
}
