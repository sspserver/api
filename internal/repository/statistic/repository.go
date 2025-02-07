// Package statistic present full API functionality of the specific object
package statistic

import (
	"context"

	"github.com/sspserver/api/models"
)

// Repository of access to the statistic
//
//go:generate mockgen -source $GOFILE -package mocks -destination mocks/repository.go
type Repository interface {
	Statistic(ctx context.Context, opts ...Option) ([]*models.StatisticAdItem, error)
	Count(ctx context.Context, opts ...Option) (int64, error)
}
