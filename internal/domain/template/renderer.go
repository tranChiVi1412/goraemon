package template

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
)

type TemplateRenderer interface {
	Render(template *Template, data map[string]interface{}) (string, error)
	ListTemplates(moduleType module.ModuleType) ([]*Template, error)
}

type templateRenderer struct {
	templates map[module.ModuleType][]*Template
}

func NewTemplateRenderer() TemplateRenderer {
	// Initialize default templates
	templates := make(map[module.ModuleType][]*Template)

	// Database templates
	templates[module.ModuleTypeDatabase] = []*Template{
		{
			Name:        "database.go",
			Path:        "internal/infrastructure/database/database.go",
			ModuleType:  module.ModuleTypeDatabase,
			Description: "Database connection setup",
			Vars: []TemplateVariable{
				{
					Name:        "dsn",
					Type:        "string",
					Description: "Database connection string",
				},
			},
			Content: `package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB(dsn string) (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}`,
		},
		{
			Name:        "config.go",
			Path:        "internal/config/database.go",
			ModuleType:  module.ModuleTypeDatabase,
			Description: "Database configuration",
			Vars: []TemplateVariable{
				{
					Name:        "dsn",
					Type:        "string",
					Description: "Database connection string",
				},
			},
			Content: `package config

type DatabaseConfig struct {
	DSN string
}`,
		},
	}

	// Cache templates
	templates[module.ModuleTypeCache] = []*Template{
		{
			Name:        "cache.go",
			Path:        "internal/infrastructure/cache/cache.go",
			ModuleType:  module.ModuleTypeCache,
			Description: "Cache connection setup",
			Vars: []TemplateVariable{
				{
					Name:        "addr",
					Type:        "string",
					Description: "Redis server address",
				},
			},
			Content: `package cache

import (
	"github.com/redis/go-redis/v9"
)

func NewRedis(addr string) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	}), nil
}`,
		},
		{
			Name:        "config.go",
			Path:        "internal/config/cache.go",
			ModuleType:  module.ModuleTypeCache,
			Description: "Cache configuration",
			Vars: []TemplateVariable{
				{
					Name:        "addr",
					Type:        "string",
					Description: "Redis server address",
				},
			},
			Content: `package config

type CacheConfig struct {
	Addr string
}`,
		},
	}

	return &templateRenderer{
		templates: templates,
	}
}

func (r *templateRenderer) Render(template *Template, data map[string]interface{}) (string, error) {
	if template == nil {
		return "", fmt.Errorf("template cannot be nil")
	}

	// Create directory if not exists
	dir := filepath.Dir(template.Path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	// Write template content to file
	if err := os.WriteFile(template.Path, []byte(template.Content), 0644); err != nil {
		return "", fmt.Errorf("failed to write template file %s: %w", template.Path, err)
	}

	return template.Path, nil
}

func (r *templateRenderer) ListTemplates(moduleType module.ModuleType) ([]*Template, error) {
	templates, exists := r.templates[moduleType]
	if !exists {
		return []*Template{}, nil
	}
	return templates, nil
}
