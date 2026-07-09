package resolvers

import (
	"context"
	"slices"

	"github.com/sspserver/api/pkg/server/graphql/accounts"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type generalResolver struct{ *Resolver }

// Account is the resolver for the account field.
func (r *generalResolver) Account(ctx context.Context, obj *accounts.Account, id uint64) (*accounts.Account, error) {
	if obj != nil || id == 0 {
		return obj, nil
	}
	acc, err := r.accounts.Account(ctx, id)
	if err != nil {
		return nil, err
	}
	return acc.Account, nil
}

// Formats is the resolver for the formats field.
func (r *generalResolver) Formats(ctx context.Context, list []*gqlmodels.AdFormat, codes []string) ([]*gqlmodels.AdFormat, error) {
	if len(codes) == 0 || len(list) > 0 {
		// If no format codes are provided, or formats are already loaded, return the existing formats
		return list, nil
	}
	return r.adformat.ListByCodes(ctx, codes)
}

// DeviceTypes is the resolver for the deviceTypes field.
func (r *generalResolver) DeviceTypes(ctx context.Context, list []*gqlmodels.DeviceType, ids []uint64) ([]*gqlmodels.DeviceType, error) {
	if len(ids) == 0 || len(list) > 0 {
		// If no device type IDs are provided, or device types are already loaded, return the existing device types
		return list, nil
	}
	deviceTypes, err := r.device_types.List(ctx)
	if err != nil {
		return nil, err
	}
	newList := []*gqlmodels.DeviceType{}
	for _, deviceType := range deviceTypes {
		if slices.Contains(ids, deviceType.ID) {
			newList = append(newList, deviceType)
		}
	}
	return newList, nil
}

// Devices is the resolver for the devices field.
func (r *generalResolver) Devices(ctx context.Context, list []*gqlmodels.DeviceModel, ids []uint64) ([]*gqlmodels.DeviceModel, error) {
	if len(ids) == 0 || len(list) > 0 {
		// If no device IDs are provided, or devices are already loaded, return the existing devices
		return list, nil
	}
	return r.device_models.ListByIDs(ctx, ids)
}

// Os is the resolver for the OS field.
func (r *generalResolver) Os(ctx context.Context, list []*gqlmodels.Os, ids []uint64) ([]*gqlmodels.Os, error) {
	if len(ids) == 0 || len(list) > 0 {
		// If no OS codes are provided, or OS are already loaded, return the existing OS
		return list, nil
	}
	return r.os.ListByIDs(ctx, ids)
}

// Browsers is the resolver for the browsers field.
func (r *generalResolver) Browsers(ctx context.Context, list []*gqlmodels.Browser, ids []uint64) ([]*gqlmodels.Browser, error) {
	if len(ids) == 0 || len(list) > 0 {
		// If no browser IDs are provided, or browsers are already loaded, return the existing browsers
		return list, nil
	}
	return r.browsers.ListByIDs(ctx, ids)
}

// Categories is the resolver for the categories field.
func (r *generalResolver) Categories(ctx context.Context, list []*gqlmodels.Category, ids []uint64) ([]*gqlmodels.Category, error) {
	if len(ids) == 0 || len(list) > 0 {
		// If no category IDs are provided, or categories are already loaded, return the existing categories
		return list, nil
	}
	return r.categories.ListByIDs(ctx, ids)
}

// Countries is the resolver for the countries field.
func (r *generalResolver) Countries(ctx context.Context, list []*gqlmodels.Country, codes []string) ([]*gqlmodels.Country, error) {
	if len(codes) == 0 || len(list) > 0 {
		// If no country codes are provided, or countries are already loaded, return the existing countries
		return list, nil
	}
	countries, err := r.geo.Countries(ctx)
	if err != nil {
		return nil, err
	}
	for _, country := range countries {
		if slices.Contains(codes, country.Code2) {
			list = append(list, country)
		}
	}
	return list, nil
}

// Languages is the resolver for the languages field.
func (r *generalResolver) Languages(ctx context.Context, list []*gqlmodels.Lang, codes []string) ([]*gqlmodels.Lang, error) {
	if len(codes) == 0 || len(list) > 0 {
		// If no language codes are provided, or languages are already loaded, return the existing languages
		return list, nil
	}
	return r.langs.Languages(ctx, &gqlmodels.LangListFilter{Iso2: codes})
}

// Applications is the resolver for the applications field.
func (r *generalResolver) Applications(ctx context.Context, list []*gqlmodels.Application, ids []uint64) ([]*gqlmodels.Application, error) {
	if len(ids) == 0 || len(list) > 0 {
		// If no application IDs are provided, or applications are already loaded, return the existing applications
		return list, nil
	}
	return r.app.ListByIDs(ctx, ids)
}

// Zones is the resolver for the zones field.
func (r *generalResolver) Zones(ctx context.Context, list []*gqlmodels.Zone, ids []uint64) ([]*gqlmodels.Zone, error) {
	if len(ids) == 0 || len(list) > 0 {
		// If no zone IDs are provided, or zones are already loaded, return the existing zones
		return list, nil
	}
	return r.zone.ListByIDs(ctx, ids)
}
