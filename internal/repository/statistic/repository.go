// Package statistic present full API functionality of the specific object
package statistic

import (
	"context"

	"github.com/geniusrabbit/archivarius/client"

	"github.com/sspserver/api/internal/server/graphql/models"
)

// Repository of access to the statistic
type Repository interface {
	Statistic(
		ctx context.Context,
		group []models.StatisticKey,
		order []*models.StatisticAdKeyOrder,
		filter *models.StatisticAdListFilter,
		page *models.Page,
	) (*client.StatisticResponse, error)
}
