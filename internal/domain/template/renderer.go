package template

import "github.com/tranChiVi1412/goraemon/internal/domain/module"

type TemplateRenderer interface {
	Render(template *Template, data map[string]interface{}) (string, error)
	ListTemplates(moduleType module.ModuleType) ([]*Template, error)
}
