package usecase

import (
	"context"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
	"github.com/tranChiVi1412/goraemon/internal/domain/project"
)

type InitProjectInput struct {
	Name       string
	ModulePath string
	Transports []project.Transport
	Modules    []*module.Module
}

type InitProjectOutput struct {
	ProjectID string
	Message   string
	Error     error
}

type NewProjectUsecase interface {
	Execute(ctx context.Context, input InitProjectInput) (InitProjectOutput, error)
}
