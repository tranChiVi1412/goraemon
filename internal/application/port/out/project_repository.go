package out

import (
	"context"

	"github.com/tranChiVi1412/goraemon/internal/domain/project"
)

// ProjectRepositoryPort defines the interface for project persistence
type ProjectRepositoryPort interface {
	// Save saves a project
	Save(ctx context.Context, project *project.Project) error

	// FindByID finds a project by ID
	FindByID(ctx context.Context, id string) (*project.Project, error)

	// FindByName finds a project by name
	FindByName(ctx context.Context, name string) (*project.Project, error)

	// List lists all projects
	List(ctx context.Context) ([]*project.Project, error)

	// Exists checks if a project exists
	Exists(ctx context.Context, name string) (bool, error)

	// Delete deletes a project
	Delete(ctx context.Context, id string) error
}
