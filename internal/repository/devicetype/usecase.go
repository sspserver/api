// Package devicetype present full API functionality of the specific object
package devicetype

import (
	"context"

	"github.com/sspserver/api/models"
)

// usecase of access to the devicetype
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/usecase.go
type Usecase interface {
	Get(ctx context.Context, id uint64) (*models.DeviceType, error)
	GetByCodename(ctx context.Context, codename string) (*models.DeviceType, error)
	FetchList(ctx context.Context) ([]*models.DeviceType, error)
	FetchListByIDs(ctx context.Context, ids []uint64) ([]*models.DeviceType, error)
	Count(ctx context.Context) (int64, error)
}
