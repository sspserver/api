// Package adformat present full API functionality of the specific object
package adformat

import (
	"context"

	"github.com/sspserver/api/models"
)

// Repository of access to the ad format
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/repository.go
type Repository interface {
	Get(ctx context.Context, id uint64) (*models.Format, error)
	GetByCodename(ctx context.Context, codename string) (*models.Format, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.Format, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, source *models.Format) (uint64, error)
	Update(ctx context.Context, id uint64, source *models.Format) error
	Delete(ctx context.Context, id uint64) error
	DeleteByCodename(ctx context.Context, codename string) error
}
