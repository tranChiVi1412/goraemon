package registry

import (
	"fmt"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
)

// ModuleRegistry defines the interface for module registration
type ModuleRegistry interface {
	// Register registers a module
	Register(module *module.Module) error

	// Get gets a module by name
	Get(name string) (*module.Module, error)

	// List lists all modules
	List() []*module.Module
}

// moduleRegistry implements ModuleRegistry
type moduleRegistry struct {
	modules map[string]*module.Module
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

// NewModuleRegistry creates a new module registry
func NewModuleRegistry() ModuleRegistry {
	return &moduleRegistry{
		modules: make(map[string]*module.Module),
	}
}

// Register registers a module
func (r *moduleRegistry) Register(module *module.Module) error {
	if module == nil {
		return fmt.Errorf("module cannot be nil")
	}

	if _, exists := r.modules[module.Name]; exists {
		return fmt.Errorf("module %s already exists", module.Name)
	}

	r.modules[module.Name] = module
	return nil
}

// Get gets a module by name
func (r *moduleRegistry) Get(name string) (*module.Module, error) {
	module, exists := r.modules[name]
	if !exists {
		return nil, fmt.Errorf("module %s not found", name)
	}
	return module, nil
}

// List lists all modules
func (r *moduleRegistry) List() []*module.Module {
	modules := make([]*module.Module, 0, len(r.modules))
	for _, module := range r.modules {
		modules = append(modules, module)
	}
	return modules
}
