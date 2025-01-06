// Package devicemodel present full API functionality of the specific object
package devicemodel

import (
	"context"

	"github.com/sspserver/api/models"
)

// Repository of access to the devicemodel
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/repository.go
type Repository interface {
	Get(ctx context.Context, id uint64) (*models.DeviceModel, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.DeviceModel, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, object *models.DeviceModel) (uint64, error)
	Update(ctx context.Context, id uint64, object *models.DeviceModel) error
	Delete(ctx context.Context, id uint64) error
}
