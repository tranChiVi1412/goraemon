package command

import "github.com/tranChiVi1412/goraemon/internal/domain/project"

type InitProjectCommand struct {
	Name       string
	ModulePath string
	Transports []project.Transport
	Modules    []string // Name of modules user want to add
}

type AddModuleCommand struct {
	ProjectID  string
	ModuleName string
	Config     map[string]string
}
