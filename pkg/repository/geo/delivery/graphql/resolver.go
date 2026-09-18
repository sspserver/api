package graphql

import (
	"context"
	"strings"

	"github.com/geniusrabbit/adcorelib/geo"

	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{}
}

// Continents is the resolver for the continents field.
func (r *QueryResolver) Continents(ctx context.Context) ([]*gqlmodels.Continent, error) {
	return FromGeoContinentModelList(geo.Continents), nil
}

// Countries is the resolver for the countries field.
func (r *QueryResolver) Countries(ctx context.Context) ([]*gqlmodels.Country, error) {
	return FromGeoCountryModelList(geo.Countries), nil
}

// Regions is the resolver for the regions field.
func (r *QueryResolver) Regions(ctx context.Context, countryCode *string) ([]*gqlmodels.Region, error) {
	if countryCode != nil && *countryCode != "" {
		c := geo.CountryByCode2(strings.ToUpper(*countryCode))
		if c.ISO2() == geo.UndefinedCountryCodeISO2 {
			return nil, nil
		}
		return FromGeoRegionModelList(c.Regions()), nil
	}
	return FromGeoRegionModelList(geo.Regions), nil
}

// Region is the resolver for the region field.
func (r *QueryResolver) Region(ctx context.Context, code string) (*gqlmodels.Region, error) {
	reg := geo.RegionByCode(strings.ToUpper(code))
	if reg.ID == 0 {
		return nil, nil
	}
	return FromGeoRegionModel(*reg), nil
}

// ContinentCountries is the resolver for Continent.countries.
func (r *QueryResolver) ContinentCountries(ctx context.Context, obj *gqlmodels.Continent) ([]*gqlmodels.Country, error) {
	return FromGeoCountryModelList(geo.CountriesByContinent(obj.Code2)), nil
}

// CountryRegions is the resolver for Country.regions.
func (r *QueryResolver) CountryRegions(ctx context.Context, obj *gqlmodels.Country) ([]*gqlmodels.Region, error) {
	return FromGeoRegionModelList(geo.CountryByID(uint8(obj.ID)).Regions()), nil
}

// RegionCountry is the resolver for Region.country.
func (r *QueryResolver) RegionCountry(ctx context.Context, obj *gqlmodels.Region) (*gqlmodels.Country, error) {
	return FromGeoCountryModel(*geo.RegionByID(uint16(obj.ID)).Country()), nil
}

// RegionParent is the resolver for Region.parent.
func (r *QueryResolver) RegionParent(ctx context.Context, obj *gqlmodels.Region) (*gqlmodels.Region, error) {
	p := geo.RegionByID(uint16(obj.ID)).Parent()
	if p == nil {
		return nil, nil
	}
	return FromGeoRegionModel(*p), nil
}
