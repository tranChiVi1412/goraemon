package project

import (
	"time"

	"github.com/tranChiVi1412/goraemon/internal/domain/module"
)

type Project struct {
	ID         string          // ID of project
	Name       string          // Name of project
	ModulePath string          // Path to module (ex: github.com/example/project)
	Transports []Transport     // Transports of project(ex: rest, grpc, mq)
	Modules    []module.Module // Modules of project (ex: DB, Cache, Logging, etc)
	CreatedAt  time.Time       // Created at
	UpdatedAt  time.Time       // Updated at
}

type Transport string

const (
	TransportREST Transport = "rest"
	TransportGRPC Transport = "grpc"
	TransportMQ   Transport = "mq"
)
