// Package rtbsource present full API functionality of the specific object
package rtbsource

import (
	"context"

	"github.com/sspserver/api/models"
)

// usecase of access to the source
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/usecase.go
type Usecase interface {
	Get(ctx context.Context, id uint64) (*models.RTBSource, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.RTBSource, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, source *models.RTBSource) (uint64, error)
	Update(ctx context.Context, id uint64, source *models.RTBSource) error
	Run(ctx context.Context, id uint64, message string) error
	Pause(ctx context.Context, id uint64, message string) error
	Approve(ctx context.Context, id uint64, message string) error
	Reject(ctx context.Context, id uint64, message string) error
	Delete(ctx context.Context, id uint64, msg *string) error
}
