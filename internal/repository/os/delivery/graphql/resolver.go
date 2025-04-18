package graphql

import (
	"context"
	"fmt"
	"slices"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"
	"github.com/geniusrabbit/blaze-api/repository"

	"github.com/sspserver/api/internal/context/ctxcache"
	"github.com/sspserver/api/internal/repository/os"
	"github.com/sspserver/api/internal/repository/os/usecase"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc os.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get OS is the resolver for the OS field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qmodels.OSPayload, error) {
	obj, err := r.cachedOS(ctx, id, true)
	if err != nil {
		return nil, err
	}
	return &qmodels.OSPayload{
		ClientMutationID: requestid.Get(ctx),
		Osid:             obj.ID,
		Os:               qmodels.FromOSModel(obj),
	}, nil
}

// List OS is the resolver for the listOS field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.OSListFilter, order []*qmodels.OSListOrder, page *qmodels.Page) (*connectors.OSConnection, error) {
	return connectors.NewOSConnection(ctx, r.uc, filter, order, page), nil
}

// ListByIDs returns a list of OS by their IDs.
func (r *QueryResolver) ListByIDs(ctx context.Context, ids []uint64) ([]*qmodels.Os, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var (
		list   = make([]*models.OS, 0, len(ids))
		cached []uint64
	)
	for _, id := range ids {
		if obj, err := r.cachedOS(ctx, id, false); err != nil {
			return nil, err
		} else {
			list = append(list, obj)
			cached = append(cached, id)
		}
	}
	if len(cached) < len(ids) {
		newIDs := gocast.IfThenExec(len(cached) == 0,
			func() []uint64 { return ids },
			func() []uint64 {
				return slices.DeleteFunc(ids, func(code uint64) bool {
					return slices.Contains(cached, code)
				})
			})
		osList, err := r.uc.FetchList(ctx,
			&os.Filter{ID: newIDs},
			&repository.PreloadOption{Fields: []string{`Versions`}},
		)
		if err != nil {
			return nil, err
		}
		r.cacheList(ctx, osList)
		list = append(list, osList...)
	}
	return qmodels.FromOSModelList(list), nil
}

// Versions is the resolver for the versions field.
func (r *QueryResolver) Versions(ctx context.Context, obj *qmodels.Os) ([]*qmodels.Os, error) {
	if len(obj.Versions) > 0 {
		return obj.Versions, nil
	}
	osList, err := r.uc.FetchList(ctx,
		&os.Filter{ParentID: []uint64{obj.ID}},
		&repository.PreloadOption{Fields: []string{`Versions`}},
	)
	if err != nil {
		return nil, err
	}
	r.cacheList(ctx, osList)
	return qmodels.FromOSModelList(osList), nil
}

// Create OS is the resolver for the createOS field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.OSCreateInput) (*qmodels.OSPayload, error) {
	var obj models.OS
	if err := input.FillModel(&obj); err != nil {
		return nil, err
	}

	id, err := r.uc.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}

	return &qmodels.OSPayload{
		ClientMutationID: requestid.Get(ctx),
		Osid:             id,
		Os:               qmodels.FromOSModel(&obj),
	}, nil
}

// Update OS is the resolver for the updateOS field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.OSUpdateInput) (*qmodels.OSPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, fmt.Errorf("OS not found")
	}
	if err = input.FillModel(obj); err != nil {
		return nil, err
	}
	if err := r.uc.Update(ctx, id, obj); err != nil {
		return nil, err
	}
	return &qmodels.OSPayload{
		ClientMutationID: requestid.Get(ctx),
		Osid:             id,
		Os:               qmodels.FromOSModel(obj),
	}, nil
}

// Delete OS is the resolver for the deleteOS field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qmodels.OSPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, fmt.Errorf("OS not found")
	}
	if err := r.uc.Delete(ctx, id, msg); err != nil {
		return nil, err
	}
	return &qmodels.OSPayload{
		ClientMutationID: requestid.Get(ctx),
		Osid:             obj.ID,
		Os:               qmodels.FromOSModel(obj),
	}, nil
}

func (r *QueryResolver) cachedOS(ctx context.Context, id uint64, orLoad bool) (*models.OS, error) {
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
	if err != nil {
		return nil, err
	}
	return obj.(*models.OS), nil
}

func (r *QueryResolver) cacheList(ctx context.Context, list []*models.OS) {
	cache := r.cacheObj(ctx)
	for _, obj := range list {
		cache.Set(obj.ID, obj)
	}
}

func (r *QueryResolver) cacheObj(ctx context.Context) ctxcache.Cacher {
	return ctxcache.GetCache(ctx, "OS")
}
