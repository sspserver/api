package archivarius

import (
	"context"
)

// StatisticResponse is a statistic response with AdItems
type StatisticResponse struct {
	Items []*AdItem
}

// Usecase of access to the statistic
type Usecase interface {
	// Statistic returns a list of items
	Statistic(ctx context.Context, opts ...Option) (*StatisticResponse, error)
	// Count returns a count of items
	Count(ctx context.Context, opts ...Option) (int64, error)
}

// UsecaseImpl is a default implementation of Usecase
type UsecaseImpl struct {
	repo Repository
}

// NewUsecase creates a new UsecaseImpl
func NewUsecase(repo Repository) *UsecaseImpl {
	return &UsecaseImpl{repo: repo}
}

// Statistic returns a list of items
func (s *UsecaseImpl) Statistic(ctx context.Context, opts ...Option) (*StatisticResponse, error) {
	stats, err := s.repo.Statistic(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &StatisticResponse{
		Items: stats,
	}, nil
}

// Count returns a count of items
func (s *UsecaseImpl) Count(ctx context.Context, opts ...Option) (int64, error) {
	return s.repo.Count(ctx, opts...)
}
