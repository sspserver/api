package graphql

import (
	"github.com/demdxx/xtypes"

	"github.com/geniusrabbit/adcorelib/geo"

	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromGeoContinentByCode(code string) *gqlmodels.Continent {
	c := geo.ContinentByCode2(code)
	if c == nil {
		return nil
	}
	return FromGeoContinentModel(*c)
}

func FromGeoContinentModel(c geo.Continent) *gqlmodels.Continent {
	return &gqlmodels.Continent{
		ID:    uint64(c.ID),
		Code2: c.Code2,
		Name:  c.Name,
	}
}

func FromGeoCountryModel(c geo.Country) *gqlmodels.Country {
	return &gqlmodels.Country{
		ID:            uint64(c.ID),
		Code2:         c.ISO2(),
		Code3:         c.Code3(),
		Name:          c.Name,
		NativeName:    c.Native,
		ContinentCode: c.Continent(),
		Continent:     FromGeoContinentByCode(c.Continent()),
		Capital:       c.Capital,
		Languages:     c.Languages(),
		PhoneCodes:    c.Phones(),
		Currency:      c.Currency(),
		TimeZones: xtypes.SliceApply(c.TimeZones(), func(tz geo.TimeZone) *gqlmodels.TimeZone {
			return &gqlmodels.TimeZone{
				Name: tz.ZoneName,
				Lon:  float64(tz.Lon),
			}
		}),
		Coordinates: &gqlmodels.Coordinates{
			Lat: float64(c.Coordinates.Lat),
			Lon: float64(c.Coordinates.Lon),
		},
	}
}

func FromGeoRegionModel(r geo.Region) *gqlmodels.Region {
	out := &gqlmodels.Region{
		ID:              uint64(r.ID),
		Code:            r.Code(),
		Name:            r.Name,
		Names:           r.Names(),
		SubdivisionType: r.Type(),
	}
	if r.HasCoordinates() {
		out.Coordinates = &gqlmodels.Coordinates{
			Lat: float64(r.Coordinates.Lat),
			Lon: float64(r.Coordinates.Lon),
		}
	}
	return out
}

func FromGeoCountryModelList(c []geo.Country) []*gqlmodels.Country {
	return xtypes.SliceApply(c, FromGeoCountryModel)
}

func FromGeoContinentModelList(c []geo.Continent) []*gqlmodels.Continent {
	return xtypes.SliceApply(c, FromGeoContinentModel)
}

func FromGeoRegionModelList(r []geo.Region) []*gqlmodels.Region {
	return xtypes.SliceApply(r, FromGeoRegionModel)
}
