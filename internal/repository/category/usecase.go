// Package category present full API functionality of the specific object
package category

import (
	"context"

	"github.com/sspserver/api/models"
)

// usecase of access to the category
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/usecase.go
type Usecase interface {
	Get(ctx context.Context, id uint64) (*models.Category, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.Category, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, object *models.Category) (uint64, error)
	Update(ctx context.Context, id uint64, object *models.Category) error
	Delete(ctx context.Context, id uint64, msg *string) error
}
