package module

import "context"

type ModuleRepository interface {
	FindByName(ctx context.Context, name string) (*Module, error)
	ListAll(ctx context.Context) ([]*Module, error)
	Exists(ctx context.Context, name string) (bool, error)
}
