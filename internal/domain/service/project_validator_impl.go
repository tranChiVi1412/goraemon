package service

import (
	"github.com/tranChiVi1412/goraemon/internal/domain/project"
	"github.com/tranChiVi1412/goraemon/internal/domain/registry"
)

type projectValidator struct{}

func NewProjectValidator() ProjectValidator {
	return &projectValidator{}
}

func (v *projectValidator) Validate(prj *project.Project, registry *registry.ModuleRegistry) error {
	// TODO: Implement validation logic
	return nil
}
