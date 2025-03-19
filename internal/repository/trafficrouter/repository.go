package trafficrouter

import (
	"context"

	"github.com/sspserver/api/models"
)

type Repository interface {
	Get(ctx context.Context, id uint64) (*models.TrafficRouter, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.TrafficRouter, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, router *models.TrafficRouter) (uint64, error)
	Update(ctx context.Context, id uint64, router *models.TrafficRouter) error
	Delete(ctx context.Context, id uint64) error
	Run(ctx context.Context, id uint64, message string) error
	Pause(ctx context.Context, id uint64, message string) error
}
