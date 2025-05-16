package registry

import "github.com/tranChiVi1412/goraemon/internal/domain/module"

type ModuleRegistry struct {
	Modules map[string]*ModuleMetaData
}

type ModuleMetaData struct {
	Name         string
	Type         module.ModuleType
	Version      string
	Description  string
	Dependencies []string
	Templates    []string // Path for template
	DocsURL      string   // URL for documentation
}
