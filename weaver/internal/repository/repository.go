package repository

import (
	"context"
	"errors"

	"github.com/codecflow/fabric/weaver/internal/namespace"
	"github.com/codecflow/fabric/weaver/internal/secret"
	"github.com/codecflow/fabric/weaver/internal/workload"
)

var ErrNotFound = errors.New("resource not found")

type Repository struct {
	Workload  workload.Repository
	Namespace namespace.Repository
	Secret    secret.Repository
}

// HealthCheck checks the health of the repository
func (r *Repository) HealthCheck(ctx context.Context) error {
	// This would be implemented by the concrete repository
	return nil
}

// Close closes the repository connections
func (r *Repository) Close() error {
	// This would be implemented by the concrete repository
	return nil
}
