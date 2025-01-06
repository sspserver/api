// Package os present full API functionality of the specific object
package os

import (
	"context"

	"github.com/sspserver/api/models"
)

// Repository of access to the os
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/repository.go
type Repository interface {
	Get(ctx context.Context, id uint64) (*models.OS, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.OS, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, object *models.OS) (uint64, error)
	Update(ctx context.Context, id uint64, object *models.OS) error
	Delete(ctx context.Context, id uint64) error
}
