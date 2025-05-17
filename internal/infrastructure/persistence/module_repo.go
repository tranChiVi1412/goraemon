package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
)

type ModuleFileRepo struct {
	repoDir string
}

func NewModuleFileRepo(repoDir string) *ModuleFileRepo {
	return &ModuleFileRepo{
		repoDir: repoDir,
	}
}

func (r *ModuleFileRepo) Exists(ctx context.Context, name string) (bool, error) {
	modulePath := filepath.Join(r.repoDir, "modules", name+".json")
	_, err := os.Stat(modulePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (r *ModuleFileRepo) FindByName(ctx context.Context, name string) (*module.Module, error) {
	modulePath := filepath.Join(r.repoDir, "modules", name+".json")
	data, err := os.ReadFile(modulePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("module %s not found", name)
		}
		return nil, fmt.Errorf("failed to read module file: %w", err)
	}

	var mod module.Module
	if err := json.Unmarshal(data, &mod); err != nil {
		return nil, fmt.Errorf("failed to unmarshal module data: %w", err)
	}

	return &mod, nil
}

func (r *ModuleFileRepo) Save(ctx context.Context, mod *module.Module) error {
	// Ensure modules directory exists
	modulesDir := filepath.Join(r.repoDir, "modules")
	if err := os.MkdirAll(modulesDir, 0755); err != nil {
		return fmt.Errorf("failed to create modules directory: %w", err)
	}

	// Marshal module to JSON
	data, err := json.Marshal(mod)
	if err != nil {
		return fmt.Errorf("failed to marshal module data: %w", err)
	}

	// Write to file
	modulePath := filepath.Join(modulesDir, mod.Name+".json")
	if err := os.WriteFile(modulePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write module file: %w", err)
	}

	return nil
}

func (r *ModuleFileRepo) ListAll(ctx context.Context) ([]*module.Module, error) {
	modulesDir := filepath.Join(r.repoDir, "modules")

	// Check if directory exists
	if _, err := os.Stat(modulesDir); os.IsNotExist(err) {
		return []*module.Module{}, nil
	}

	// Read all files in the modules directory
	files, err := os.ReadDir(modulesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read modules directory: %w", err)
	}

	var modules []*module.Module
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		// Read and parse each module file
		data, err := os.ReadFile(filepath.Join(modulesDir, file.Name()))
		if err != nil {
			return nil, fmt.Errorf("failed to read module file %s: %w", file.Name(), err)
		}

		var mod module.Module
		if err := json.Unmarshal(data, &mod); err != nil {
			return nil, fmt.Errorf("failed to unmarshal module data from %s: %w", file.Name(), err)
		}

		modules = append(modules, &mod)
	}

	return modules, nil
}
