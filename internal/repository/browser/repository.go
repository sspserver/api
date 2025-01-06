// Package browser present full API functionality of the specific object
package browser

import (
	"context"

	"github.com/sspserver/api/models"
)

// Repository of access to the browser
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/repository.go
type Repository interface {
	Get(ctx context.Context, id uint64) (*models.Browser, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.Browser, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, object *models.Browser) (uint64, error)
	Update(ctx context.Context, id uint64, object *models.Browser) error
	Delete(ctx context.Context, id uint64) error
}
