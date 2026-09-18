package graphql

import (
	"context"
	"testing"
)

func TestRegionQuery(t *testing.T) {
	r := NewQueryResolver()
	got, err := r.Region(context.Background(), "us-ca")
	if err != nil {
		t.Fatal(err)
	}
	if got == nil || got.Code != "US-CA" {
		t.Fatalf("region = %+v, want US-CA", got)
	}
	country, err := r.RegionCountry(context.Background(), got)
	if err != nil {
		t.Fatal(err)
	}
	if country == nil || country.Code2 != "US" {
		t.Fatalf("country = %+v, want US", country)
	}
}

func TestRegionUKAlias(t *testing.T) {
	r := NewQueryResolver()
	got, err := r.Region(context.Background(), "UK-ENG")
	if err != nil {
		t.Fatal(err)
	}
	if got == nil || got.Code != "GB-ENG" {
		t.Fatalf("region = %+v, want GB-ENG", got)
	}
}

func TestRegionsByCountry(t *testing.T) {
	r := NewQueryResolver()
	cc := "US"
	list, err := r.Regions(context.Background(), &cc)
	if err != nil {
		t.Fatal(err)
	}
	if len(list) == 0 {
		t.Fatal("US regions must not be empty")
	}
}

func TestContinentCountries(t *testing.T) {
	r := NewQueryResolver()
	continents, err := r.Continents(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, c := range continents {
		if c.Code2 != "EU" {
			continue
		}
		found = true
		countries, err := r.ContinentCountries(context.Background(), c)
		if err != nil {
			t.Fatal(err)
		}
		if len(countries) == 0 {
			t.Fatal("EU countries must not be empty")
		}
	}
	if !found {
		t.Fatal("EU continent missing")
	}
}

func TestUnknownRegion(t *testing.T) {
	r := NewQueryResolver()
	got, err := r.Region(context.Background(), "ZZ-ZZ")
	if err != nil {
		t.Fatal(err)
	}
	if got != nil {
		t.Fatalf("unknown region = %+v, want nil", got)
	}
}
