package service

import (
	"github.com/tranChiVi1412/goraemon/internal/domain/project"
	"github.com/tranChiVi1412/goraemon/internal/domain/registry"
)

type ProjectValidator interface {
	Validate(project *project.Project, registry *registry.ModuleRegistry) error
}
