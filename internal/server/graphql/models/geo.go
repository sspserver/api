package models

import (
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/gogeo"
)

func FromGeoContinentByCode(code string) *Continent {
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

func FromGeoContinentModel(c gogeo.Continent) *Continent {
	return &Continent{
		ID:    uint64(c.ID),
		Code2: c.Code2,
		Name:  c.Name,
	}
}

func FromGeoCountryModel(c gogeo.Country) *Country {
	return &Country{
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
		TimeZones: xtypes.SliceApply(c.TimeZones, func(tz gogeo.TimeZone) *TimeZone {
			return &TimeZone{
				Name: tz.ZoneName,
				Lon:  tz.Lon,
			}
		}),
		Coordinates: &Coordinates{
			Lat: c.Coordinates.Lat,
			Lon: c.Coordinates.Lon,
		},
	}
}

func FromGeoCountryModelList(c []gogeo.Country) []*Country {
	return xtypes.SliceApply(c, FromGeoCountryModel)
}

func FromGeoContinentModelList(c []gogeo.Continent) []*Continent {
	return xtypes.SliceApply(c, FromGeoContinentModel)
}
