package service

import (
	"github.com/tranChiVi1412/goraemon/internal/domain/module"
	"github.com/tranChiVi1412/goraemon/internal/domain/registry"
)

type DependencyResolver interface {
	Resolve(module *module.Module, registry *registry.ModuleRegistry) ([]string, error) // Return list of valid module and dependencies
}
