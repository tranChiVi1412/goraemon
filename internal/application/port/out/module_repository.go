package out

import (
	"context"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
)

type ModuleRepositoryPort interface {
	Save(ctx context.Context, module *module.Module) error
	FindByName(ctx context.Context, name string) (*module.Module, error)
	ListByType(ctx context.Context, moduleType module.ModuleType) ([]*module.Module, error)
	ListAll(ctx context.Context) ([]*module.Module, error)
	Exists(ctx context.Context, name string) (bool, error)
	Delete(ctx context.Context, name string) error
}
