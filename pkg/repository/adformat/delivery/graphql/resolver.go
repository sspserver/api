package graphql

import (
	"context"
	"fmt"
	"slices"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/pkg/context/ctxcache"
	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/adformat"
	"github.com/sspserver/api/pkg/repository/adformat/usecase"
	"github.com/sspserver/api/pkg/server/graphql/connectors"
	qmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
	uc adformat.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get RTBAccessPoint is the resolver for the RTBAccessPoint field.
func (r *QueryResolver) Get(ctx context.Context, id uint64, codename string) (*qmodels.AdFormatPayload, error) {
	var (
		err error
		obj *models.Format
	)
	if id == 0 {
		obj, err = r.cachedAdFormat(ctx, codename, true)
	} else {
		obj, err = r.cachedAdFormatByID(ctx, id, true)
	}
	if err != nil {
		return nil, err
	}
	return &qmodels.AdFormatPayload{
		ClientMutationID: requestid.Get(ctx),
		FormatID:         gocast.IfThenExec(obj != nil, func() uint64 { return obj.ID }, func() uint64 { return 0 }),
		Format:           qmodels.FromAdFormatModel(obj),
	}, nil
}

// List RTBAccessPoints is the resolver for the listRTBAccessPoints field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.AdFormatListFilter, order *qmodels.AdFormatListOrder, page *qmodels.Page) (*connectors.AdFormatConnection, error) {
	return connectors.NewAdFormatConnection(ctx, r.uc, filter, order, page), nil
}

// ListByCodes returns a list of ad formats by their codes.
func (r *QueryResolver) ListByCodes(ctx context.Context, codes []string) ([]*qmodels.AdFormat, error) {
	if len(codes) == 0 {
		return nil, nil
	}
	var (
		list   = make([]*models.Format, 0, len(codes))
		cached []string
	)
	for _, code := range codes {
		if obj, err := r.cachedAdFormat(ctx, code, false); err != nil {
			return nil, err
		} else {
			list = append(list, obj)
			cached = append(cached, code)
		}
	}
	if len(cached) < len(codes) {
		// not all formats are cached, so we need to load them from the database
		// and cache them
		newCodes := gocast.IfThenExec(len(cached) == 0,
			func() []string { return codes },
			func() []string {
				return slices.DeleteFunc(codes, func(code string) bool {
					return slices.Contains(cached, code)
				})
			})
		newList, err := r.uc.FetchList(ctx, &adformat.Filter{Codename: newCodes}, nil, nil)
		if err != nil {
			return nil, err
		}
		cache := r.cacheObj(ctx)
		for _, obj := range newList {
			cache.Set("code_"+obj.Codename, obj)
			cache.Set(obj.ID, obj)
		}
		list = append(list, newList...)
	}
	return qmodels.FromAdFormatModelList(list), nil
}

// Create RTBAccessPoint is the resolver for the createRTBAccessPoint field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.AdFormatInput) (*qmodels.AdFormatPayload, error) {
	var object models.Format
	input.FillModel(&object)

	id, err := r.uc.Create(ctx, &object)
	if err != nil {
		return nil, err
	}

	return &qmodels.AdFormatPayload{
		ClientMutationID: requestid.Get(ctx),
		FormatID:         id,
		Format:           qmodels.FromAdFormatModel(&object),
	}, nil
}

// Update RTBAccessPoint is the resolver for the updateRTBAccessPoint field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.AdFormatInput) (*qmodels.AdFormatPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if object == nil {
		return nil, fmt.Errorf("AdFromat not found")
	}

	input.FillModel(object)

	if err = r.uc.Update(ctx, id, object); err != nil {
		return nil, err
	}

	return &qmodels.AdFormatPayload{
		ClientMutationID: requestid.Get(ctx),
		FormatID:         gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Format:           qmodels.FromAdFormatModel(object),
	}, nil
}

// Delete RTBAccessPoint is the resolver for the deleteRTBAccessPoint field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, codename string, msg *string) (*qmodels.AdFormatPayload, error) {
	var (
		err    error
		object *models.Format
	)
	if id == 0 {
		object, err = r.uc.GetByCodename(ctx, codename)
	} else {
		object, err = r.uc.Get(ctx, id)
	}
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, fmt.Errorf("AdFormat not found")
	}
	if id == 0 {
		err = r.uc.DeleteByCodename(ctx, codename, msg)
	} else {
		err = r.uc.Delete(ctx, id, msg)
	}
	if err != nil {
		return nil, err
	}
	return &qmodels.AdFormatPayload{
		ClientMutationID: requestid.Get(ctx),
		FormatID:         object.ID,
		Format:           qmodels.FromAdFormatModel(object),
	}, nil
}

func (r *QueryResolver) cachedAdFormat(ctx context.Context, codename string, orLoad bool) (*models.Format, error) {
	var (
		cache = r.cacheObj(ctx)
		obj   any
		err   error
	)
	if orLoad {
		obj, err = cache.GetOrCache("code_"+codename, func(any) (any, error) {
			return r.uc.GetByCodename(ctx, codename)
		})
	} else {
		obj = cache.Get("code_" + codename)
	}
	if err != nil || obj == nil {
		return nil, err
	}
	return obj.(*models.Format), nil
}

func (r *QueryResolver) cachedAdFormatByID(ctx context.Context, id uint64, orLoad bool) (*models.Format, error) {
	var (
		cache = r.cacheObj(ctx)
		obj   any
		err   error
	)
	if orLoad {
		obj, err = cache.GetOrCache(id, func(any) (any, error) {
			return r.uc.Get(ctx, id)
		})
	} else {
		obj = cache.Get(id)
	}
	if err != nil || obj == nil {
		return nil, err
	}
	return obj.(*models.Format), nil
}

func (r *QueryResolver) cacheObj(ctx context.Context) ctxcache.Cacher {
	return ctxcache.GetCache(ctx, "AdFormat")
}
