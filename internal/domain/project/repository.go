package project

import (
	"context"
)

type ProjectRepository interface {
	Save(ctx context.Context, project *Project) error
	Load(ctx context.Context, id string) (*Project, error)
	Exists(ctx context.Context, id string) (bool, error)
	List(ctx context.Context) ([]*Project, error)
}
