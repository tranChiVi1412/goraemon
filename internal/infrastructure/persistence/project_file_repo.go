package persistence

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/tranChiVi1412/goraemon/internal/domain/project"
	"gopkg.in/yaml.v3"
)

type ProjectFileRepo struct {
	mu  sync.Mutex
	dir string
}

func NewProjectFileRepo(dir string) *ProjectFileRepo {
	return &ProjectFileRepo{dir: dir}
}

func (r *ProjectFileRepo) Save(ctx context.Context, project *project.Project) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Create project directory if not exists
	projectDir := filepath.Join(r.dir, project.ID)
	if err := os.MkdirAll(projectDir, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create project file
	path := filepath.Join(projectDir, ".yaml")
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create project file: %w", err)
	}
	defer f.Close()

	// Encode project to YAML
	enc := yaml.NewEncoder(f)
	if err := enc.Encode(project); err != nil {
		return fmt.Errorf("failed to encode project: %w", err)
	}

	return nil
}

func (r *ProjectFileRepo) Load(ctx context.Context, projectID string) (*project.Project, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	path := filepath.Join(r.dir, projectID, ".yaml")
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open project file: %w", err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	var prj project.Project
	if err := dec.Decode(&prj); err != nil {
		return nil, fmt.Errorf("failed to decode project file: %w", err)
	}
	return &prj, nil
}

func (r *ProjectFileRepo) Exists(ctx context.Context, projectID string) (bool, error) {
	files, err := os.ReadDir(r.dir)
	if err != nil {
		return false, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".yaml" {
			continue
		}
		f, err := os.Open(filepath.Join(r.dir, file.Name()))
		if err != nil {
			continue
		}
		var prj project.Project
		if err := yaml.NewDecoder(f).Decode(&prj); err != nil {
			f.Close()
			continue
		}
		if prj.ID == projectID {
			f.Close()
			return true, nil
		}
		f.Close()
	}
	return false, nil
}

func (r *ProjectFileRepo) List(ctx context.Context) ([]*project.Project, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	files, err := os.ReadDir(r.dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var projects []*project.Project
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".yaml" {
			continue
		}
		f, err := os.Open(filepath.Join(r.dir, file.Name()))
		if err != nil {
			continue
		}
		var prj project.Project
		if err := yaml.NewDecoder(f).Decode(&prj); err != nil {
			f.Close()
			continue
		}
		projects = append(projects, &prj)
		f.Close()
	}
	return projects, nil
}
