package project

import (
	"time"

	"github.com/google/uuid"
	"github.com/tranChiVi1412/goraemon/internal/domain/module"
)

type Project struct {
	ID         string           // ID of project
	Name       string           // Name of project
	ModulePath string           // Path to module (ex: github.com/example/project)
	Transports []Transport      // Transports of project(ex: rest, grpc, mq)
	Modules    []*module.Module // Modules of project (ex: DB, Cache, Logging, etc)
	CreatedAt  time.Time        // Created at
	UpdatedAt  time.Time        // Updated at
}

type Transport string

const (
	TransportREST Transport = "rest"
	TransportGRPC Transport = "grpc"
	TransportMQ   Transport = "mq"
)

func NewProject(name, modulePath string, transports []Transport, modules []*module.Module, createdAt time.Time) *Project {
	return &Project{
		ID:         uuid.New().String(),
		Name:       name,
		ModulePath: modulePath,
		Transports: transports,
		Modules:    modules,
		CreatedAt:  createdAt,
		UpdatedAt:  createdAt,
	}
}

// AddModule adds a module to the project
func (p *Project) AddModule(module *module.Module) {
	p.Modules = append(p.Modules, module)
	p.UpdatedAt = time.Now()
}

// RemoveModule removes a module from the project
func (p *Project) RemoveModule(name string) {
	for i, module := range p.Modules {
		if module.Name == name {
			p.Modules = append(p.Modules[:i], p.Modules[i+1:]...)
			p.UpdatedAt = time.Now()
			return
		}
	}
}

// HasModule checks if the project has a specific module
func (p *Project) HasModule(name string) bool {
	for _, module := range p.Modules {
		if module.Name == name {
			return true
		}
	}
	return false
}

// GetModule returns a module by name
func (p *Project) GetModule(name string) *module.Module {
	for _, module := range p.Modules {
		if module.Name == name {
			return module
		}
	}
	return nil
}
