package graphql

import (
	"context"

	"github.com/geniusrabbit/gogeo"

	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{}
}

// Continents is the resolver for the continents field.
func (r *QueryResolver) Continents(ctx context.Context) ([]*gqlmodels.Continent, error) {
	return FromGeoContinentModelList(gogeo.Continents), nil
}

// Countries is the resolver for the countries field.
func (r *QueryResolver) Countries(ctx context.Context) ([]*gqlmodels.Country, error) {
	return FromGeoCountryModelList(gogeo.Countries), nil
}
