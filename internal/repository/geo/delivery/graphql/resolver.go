package graphql

import (
	"context"

	"github.com/geniusrabbit/gogeo"

	"github.com/sspserver/api/internal/server/graphql/models"
)

type QueryResolver struct {
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{}
}

// Continents is the resolver for the continents field.
func (r *QueryResolver) Continents(ctx context.Context) ([]*models.Continent, error) {
	return models.FromGeoContinentModelList(gogeo.Continents), nil
}

// Countries is the resolver for the countries field.
func (r *QueryResolver) Countries(ctx context.Context) ([]*models.Country, error) {
	return models.FromGeoCountryModelList(gogeo.Countries), nil
}
