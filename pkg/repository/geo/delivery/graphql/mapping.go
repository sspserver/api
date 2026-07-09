package graphql

import (
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/gogeo"

	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromGeoContinentByCode(code string) *gqlmodels.Continent {
	if code == "" {
		return nil
	}
	for _, c := range gogeo.Continents {
		if c.Code2 == code {
			return FromGeoContinentModel(c)
		}
	}
	return nil
}

func FromGeoContinentModel(c gogeo.Continent) *gqlmodels.Continent {
	return &gqlmodels.Continent{
		ID:    uint64(c.ID),
		Code2: c.Code2,
		Name:  c.Name,
	}
}

func FromGeoCountryModel(c gogeo.Country) *gqlmodels.Country {
	return &gqlmodels.Country{
		ID:            uint64(c.ID),
		Code2:         c.Code2,
		Code3:         c.Code3,
		Name:          c.Name,
		NativeName:    c.Native,
		ContinentCode: c.Continent,
		Continent:     FromGeoContinentByCode(c.Continent),
		Capital:       c.Capital,
		Languages:     c.Languages,
		PhoneCodes:    c.Phones,
		Currency:      c.Currency,
		TimeZones: xtypes.SliceApply(c.TimeZones, func(tz gogeo.TimeZone) *gqlmodels.TimeZone {
			return &gqlmodels.TimeZone{
				Name: tz.ZoneName,
				Lon:  tz.Lon,
			}
		}),
		Coordinates: &gqlmodels.Coordinates{
			Lat: c.Coordinates.Lat,
			Lon: c.Coordinates.Lon,
		},
	}
}

func FromGeoCountryModelList(c []gogeo.Country) []*gqlmodels.Country {
	return xtypes.SliceApply(c, FromGeoCountryModel)
}

func FromGeoContinentModelList(c []gogeo.Continent) []*gqlmodels.Continent {
	return xtypes.SliceApply(c, FromGeoContinentModel)
}
