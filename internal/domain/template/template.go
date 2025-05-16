package template

import "github.com/tranChiVi1412/goraemon/internal/domain/module"

type Template struct {
	Name        string
	Path        string
	ModuleType  module.ModuleType
	Vars        []TemplateVariable
	Description string
}

type TemplateVariable struct {
	Key         string
	Description string
	Required    bool
	Default     string
}
