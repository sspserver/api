// Package devicemaker present full API functionality of the specific object
package devicemaker

import (
	"context"

	"github.com/sspserver/api/models"
)

// usecase of access to the devicemaker
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/usecase.go
type Usecase interface {
	Get(ctx context.Context, id uint64) (*models.DeviceMaker, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.DeviceMaker, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, object *models.DeviceMaker) (uint64, error)
	Update(ctx context.Context, id uint64, object *models.DeviceMaker) error
	Delete(ctx context.Context, id uint64, msg *string) error
}
