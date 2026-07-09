// Package adformat present full API functionality of the specific object
package adformat

import (
	"context"
)

// Repository of access to the ad format
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/repository.go
type Repository interface {
	Get(ctx context.Context, id uint64) (*Format, error)
	GetByCodename(ctx context.Context, codename string) (*Format, error)
	FetchList(ctx context.Context, qops ...Option) ([]*Format, error)
	Count(ctx context.Context, qops ...Option) (int64, error)
	Create(ctx context.Context, source *Format) (uint64, error)
	Update(ctx context.Context, id uint64, source *Format) error
	Delete(ctx context.Context, id uint64) error
	DeleteByCodename(ctx context.Context, codename string) error
}
