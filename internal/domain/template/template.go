package template

import (
	"github.com/tranChiVi1412/goraemon/internal/domain/module"
)

type Template struct {
	Name        string
	Path        string
	ModuleType  module.ModuleType
	Description string
	Vars        []TemplateVariable
	Content     string
}

type TemplateVariable struct {
	Name        string
	Type        string
	Description string
}
