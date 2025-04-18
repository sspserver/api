package graphql

import (
	"context"
	"fmt"
	"slices"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/internal/context/ctxcache"
	"github.com/sspserver/api/internal/repository/browser"
	"github.com/sspserver/api/internal/repository/browser/usecase"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc browser.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get Browser is the resolver for the Browser field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qmodels.BrowserPayload, error) {
	obj, err := r.cachedItemByID(ctx, id, true)
	if err != nil {
		return nil, err
	}
	return &qmodels.BrowserPayload{
		ClientMutationID: requestid.Get(ctx),
		BrowserID:        obj.ID,
		Browser:          qmodels.FromBrowserModel(obj),
	}, nil
}

// List Browser is the resolver for the listBrowser field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.BrowserListFilter, order []*qmodels.BrowserListOrder, page *qmodels.Page) (*connectors.BrowserConnection, error) {
	return connectors.NewBrowserConnection(ctx, r.uc, filter, order, page), nil
}

func (r *QueryResolver) ListByIDs(ctx context.Context, ids []uint64) ([]*qmodels.Browser, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var (
		list   = make([]*models.Browser, 0, len(ids))
		cached []uint64
	)
	for _, id := range ids {
		if obj, err := r.cachedItemByID(ctx, id, false); err != nil {
			return nil, err
		} else {
			list = append(list, obj)
			cached = append(cached, id)
		}
	}
	if len(cached) < len(ids) {
		// not all formats are cached, so we need to load them from the database
		// and cache them
		newIDs := gocast.IfThenExec(len(cached) == 0,
			func() []uint64 { return ids },
			func() []uint64 {
				return slices.DeleteFunc(ids, func(id uint64) bool {
					return slices.Contains(cached, id)
				})
			})
		newList, err := r.uc.FetchList(ctx, &browser.Filter{ID: newIDs}, nil, nil)
		if err != nil {
			return nil, err
		}
		cache := r.cacheObj(ctx)
		for _, obj := range newList {
			cache.Set(obj.ID, obj)
		}
		list = append(list, newList...)
	}
	return qmodels.FromBrowserModelList(list), nil
}

// Create Browser is the resolver for the createBrowser field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.BrowserCreateInput) (*qmodels.BrowserPayload, error) {
	var obj models.Browser
	if err := input.FillModel(&obj); err != nil {
		return nil, err
	}

	id, err := r.uc.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}

	return &qmodels.BrowserPayload{
		ClientMutationID: requestid.Get(ctx),
		BrowserID:        id,
		Browser:          qmodels.FromBrowserModel(&obj),
	}, nil
}

// Update Browser is the resolver for the updateBrowser field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.BrowserUpdateInput) (*qmodels.BrowserPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, fmt.Errorf("browser not found")
	}
	if err = input.FillModel(obj); err != nil {
		return nil, err
	}
	if err := r.uc.Update(ctx, id, obj); err != nil {
		return nil, err
	}
	return &qmodels.BrowserPayload{
		ClientMutationID: requestid.Get(ctx),
		BrowserID:        id,
		Browser:          qmodels.FromBrowserModel(obj),
	}, nil
}

// Delete Browser is the resolver for the deleteBrowser field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qmodels.BrowserPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, fmt.Errorf("browser not found")
	}
	if err := r.uc.Delete(ctx, id, msg); err != nil {
		return nil, err
	}
	return &qmodels.BrowserPayload{
		ClientMutationID: requestid.Get(ctx),
		BrowserID:        obj.ID,
		Browser:          qmodels.FromBrowserModel(obj),
	}, nil
}

func (r *QueryResolver) cachedItemByID(ctx context.Context, id uint64, orLoad bool) (*models.Browser, error) {
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
	return obj.(*models.Browser), nil
}

func (r *QueryResolver) cacheObj(ctx context.Context) ctxcache.Cacher {
	return ctxcache.GetCache(ctx, "Browser")
}
