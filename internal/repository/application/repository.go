// Package application present full API functionality of the specific object
package application

import (
	"context"

	"github.com/sspserver/api/models"
)

// Repository of access to the application
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/repository.go
type Repository interface {
	Get(ctx context.Context, id uint64) (*models.Application, error)
	FetchList(ctx context.Context, qops ...Option) ([]*models.Application, error)
	Count(ctx context.Context, qops ...Option) (int64, error)

	Create(ctx context.Context, source *models.Application) (uint64, error)
	Update(ctx context.Context, id uint64, source *models.Application) error
	Delete(ctx context.Context, id uint64, msg string) error

	Run(ctx context.Context, id uint64, msg string) error
	Pause(ctx context.Context, id uint64, msg string) error
	Approve(ctx context.Context, id uint64, msg string) error
	Reject(ctx context.Context, id uint64, msg string) error
}
