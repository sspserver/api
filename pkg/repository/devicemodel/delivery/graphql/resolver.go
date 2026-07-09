package graphql

import (
	"context"
	"fmt"
	"slices"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/pkg/context/ctxcache"
	"github.com/sspserver/api/pkg/repository/devicemodel"
	"github.com/sspserver/api/pkg/repository/devicemodel/models"
	"github.com/sspserver/api/pkg/repository/devicemodel/usecase"
	qmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
	uc devicemodel.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get DeviceModel is the resolver for the DeviceModel field.
func (r *QueryResolver) Get(ctx context.Context, id uint64, codename string) (*qmodels.DeviceModelPayload, error) {
	var (
		object *models.DeviceModel
		err    error
	)
	if codename != "" {
		object, err = r.uc.GetByCodename(ctx, codename)
	} else {
		object, err = r.uc.Get(ctx, id)
	}
	if err != nil {
		return nil, err
	}
	return &qmodels.DeviceModelPayload{
		ClientMutationID: requestid.Get(ctx),
		ModelID:          gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Model:            FromDeviceModelModel(object),
	}, nil
}

// List DeviceModel is the resolver for the listDeviceModels field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.DeviceModelListFilter, order []*qmodels.DeviceModelListOrder, page *qmodels.Page) (*DeviceModelConnection, error) {
	return NewDeviceModelConnection(ctx, r.uc, filter, order, page), nil
}

// ListByIDs returns a list of DeviceModels by their IDs.
func (r *QueryResolver) ListByIDs(ctx context.Context, ids []uint64) ([]*qmodels.DeviceModel, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var (
		list   = make([]*models.DeviceModel, 0, len(ids))
		cached []uint64
	)
	for _, id := range ids {
		if obj, err := r.cachedItemByID(ctx, id, false); err != nil {
			return nil, err
		} else if obj != nil {
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
		newList, err := r.uc.FetchList(ctx, &devicemodel.Filter{ID: newIDs})
		if err != nil {
			return nil, err
		}
		cache := r.cacheObj(ctx)
		for _, obj := range newList {
			if obj != nil {
				cache.Set(obj.ID, obj)
			}
		}
		list = append(list, newList...)
	}
	return FromDeviceModelModelList(list), nil
}

// Create DeviceModel is the resolver for the createDeviceModel field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.DeviceModelCreateInput) (*qmodels.DeviceModelPayload, error) {
	var object models.DeviceModel
	if err := FillDeviceModelCreateInputModel(input, &object); err != nil {
		return nil, err
	}

	id, err := r.uc.Create(ctx, &object)
	if err != nil {
		return nil, err
	}

	// Get object to return
	nObject, err := r.uc.Get(ctx, id)
	if err != nil {
		nObject = &object
	}

	return &qmodels.DeviceModelPayload{
		ClientMutationID: requestid.Get(ctx),
		ModelID:          id,
		Model:            FromDeviceModelModel(nObject),
	}, nil
}

// Update DeviceModel is the resolver for the updateDeviceModel field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.DeviceModelUpdateInput) (*qmodels.DeviceModelPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if object == nil {
		return nil, fmt.Errorf("DeviceModel not found")
	}

	FillDeviceModelUpdateInputModel(input, object)
	if err = r.uc.Update(ctx, id, object); err != nil {
		return nil, err
	}

	// Get object to return
	nObject, err := r.uc.Get(ctx, id)
	if err != nil {
		nObject = object
	}

	return &qmodels.DeviceModelPayload{
		ClientMutationID: requestid.Get(ctx),
		ModelID:          gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Model:            FromDeviceModelModel(nObject),
	}, nil
}

// Delete DeviceModel is the resolver for the deleteDeviceModel field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qmodels.DeviceModelPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, fmt.Errorf("DeviceModel not found")
	}
	if err := r.uc.Delete(ctx, id, msg); err != nil {
		return nil, err
	}
	return &qmodels.DeviceModelPayload{
		ClientMutationID: requestid.Get(ctx),
		ModelID:          object.ID,
		Model:            FromDeviceModelModel(object),
	}, nil
}

func (r *QueryResolver) cachedItemByID(ctx context.Context, id uint64, orLoad bool) (*models.DeviceModel, error) {
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
	return obj.(*models.DeviceModel), nil
}

func (r *QueryResolver) cacheObj(ctx context.Context) ctxcache.Cacher {
	return ctxcache.GetCache(ctx, "DeviceModel")
}
