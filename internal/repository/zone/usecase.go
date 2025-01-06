package zone

import (
	"context"

	"github.com/sspserver/api/models"
)

// Usecase of access to the zone
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/repository.go
type Usecase interface {
	Get(ctx context.Context, id uint64) (*models.Zone, error)
	GetByCodename(ctx context.Context, codename string) (*models.Zone, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.Zone, error)
	Count(ctx context.Context, qops ...Option) (int64, error)

	Create(ctx context.Context, obj *models.Zone) (uint64, error)
	Update(ctx context.Context, id uint64, obj *models.Zone) error
	Delete(ctx context.Context, id uint64, msg string) error

	Run(ctx context.Context, id uint64, msg string) error
	Pause(ctx context.Context, id uint64, msg string) error
	Approve(ctx context.Context, id uint64, msg string) error
	Reject(ctx context.Context, id uint64, msg string) error
}
