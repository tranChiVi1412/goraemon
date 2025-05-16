package template

import (
	"github.com/tranChiVi1412/goraemon/internal/domain/project"
	"github.com/tranChiVi1412/goraemon/internal/domain/registry"
)

type TemplateSelector interface {
	Select(project *project.Project, registry *registry.ModuleRegistry) (*Template, error)
}
