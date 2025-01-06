// Package adformat present full API functionality of the specific object
package adformat

import (
	"context"

	"github.com/geniusrabbit/blaze-api/repository"

	"github.com/sspserver/api/models"
)

// usecase of access to the ad format
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/usecase.go
type Usecase interface {
	Get(ctx context.Context, id uint64) (*models.Format, error)
	GetByCodename(ctx context.Context, codename string) (*models.Format, error)
	FetchList(ctx context.Context, qops ...repository.QOption) ([]*models.Format, error)
	Count(ctx context.Context, qops ...repository.QOption) (int64, error)
	Create(ctx context.Context, source *models.Format) (uint64, error)
	Update(ctx context.Context, id uint64, source *models.Format) error
	Delete(ctx context.Context, id uint64, msg *string) error
	DeleteByCodename(ctx context.Context, codename string, msg *string) error
}
