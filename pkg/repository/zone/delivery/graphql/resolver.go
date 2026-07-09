package graphql

import (
	"context"
	"fmt"
	"slices"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/pkg/context/ctxcache"
	"github.com/sspserver/api/pkg/repository/zone"
	"github.com/sspserver/api/pkg/repository/zone/models"
	"github.com/sspserver/api/pkg/repository/zone/usecase"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
	uc zone.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get is the resolver for the zone field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*gqlmodels.ZonePayload, error) {
	obj, err := r.cachedZoneByID(ctx, id, true)
	if err != nil {
		return nil, err
	}
	return &gqlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		Zone:             FromZoneModel(obj),
	}, nil
}

// List Zones is the resolver for the listApplications field.
func (r *QueryResolver) List(ctx context.Context, filter *gqlmodels.ZoneListFilter, order *gqlmodels.ZoneListOrder, page *gqlmodels.Page) (*ZoneConnection, error) {
	return NewZoneConnection(ctx, r.uc, filter, order, page), nil
}

func (r *QueryResolver) ListByIDs(ctx context.Context, ids []uint64) ([]*gqlmodels.Zone, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var (
		list   = make([]*models.Zone, 0, len(ids))
		cached []uint64
	)

	for _, id := range ids {
		if id == 0 {
			continue
		}
		if obj, err := r.cachedZoneByID(ctx, id, false); err != nil {
			return nil, err
		} else if obj != nil {
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
		zoneList, err := r.uc.FetchList(ctx, &zone.Filter{ID: newIDs})
		if err != nil {
			return nil, err
		}
		r.cacheList(ctx, zoneList)
		list = append(list, zoneList...)
	}

	return FromZoneModelList(list), nil
}

// Create Zone is the resolver for the createApplication field.
func (r *QueryResolver) Create(ctx context.Context, input gqlmodels.ZoneCreateInput) (*gqlmodels.ZonePayload, error) {
	var obj models.Zone
	FillZoneCreateInputModel(input, &obj)

	id, err := r.uc.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}

	return &gqlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             FromZoneModel(&obj),
	}, nil
}

// Update Zone is the resolver for the updateApplication field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input gqlmodels.ZoneUpdateInput) (*gqlmodels.ZonePayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, fmt.Errorf("zone not found")
	}

	FillZoneUpdateInputModel(input, obj)
	if err = r.uc.Update(ctx, id, obj); err != nil {
		return nil, err
	}

	return &gqlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             FromZoneModel(obj),
	}, nil
}

// Delete Zone is the resolver for the deleteApplication field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*gqlmodels.ZonePayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, fmt.Errorf("zone not found")
	}

	if err = r.uc.Delete(ctx, id, gocast.PtrAsValue(msg, "")); err != nil {
		return nil, err
	}

	return &gqlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             FromZoneModel(obj),
	}, nil
}

// Run Zone is the resolver for the runApplication field.
func (r *QueryResolver) Run(ctx context.Context, id uint64, msg *string) (*gqlmodels.ZonePayload, error) {
	err := r.uc.Run(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	zone, _ := r.uc.Get(ctx, id)
	return &gqlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             FromZoneModel(zone),
	}, nil
}

// Pause Zone is the resolver for the pauseApplication field.
func (r *QueryResolver) Pause(ctx context.Context, id uint64, msg *string) (*gqlmodels.ZonePayload, error) {
	err := r.uc.Pause(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	zone, _ := r.uc.Get(ctx, id)
	return &gqlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             FromZoneModel(zone),
	}, nil
}

// Approve Zone is the resolver for the approveApplication field.
func (r *QueryResolver) Approve(ctx context.Context, id uint64, msg *string) (*gqlmodels.ZonePayload, error) {
	err := r.uc.Approve(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	zone, _ := r.uc.Get(ctx, id)
	return &gqlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             FromZoneModel(zone),
	}, nil
}

// Reject Zone is the resolver for the rejectApplication field.
func (r *QueryResolver) Reject(ctx context.Context, id uint64, msg *string) (*gqlmodels.ZonePayload, error) {
	err := r.uc.Reject(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	zone, _ := r.uc.Get(ctx, id)
	return &gqlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             FromZoneModel(zone),
	}, nil
}

func (r *QueryResolver) cachedZoneByID(ctx context.Context, id uint64, orLoad bool) (*models.Zone, error) {
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
	return obj.(*models.Zone), nil
}

func (r *QueryResolver) cacheList(ctx context.Context, list []*models.Zone) {
	if len(list) > 0 {
		cache := r.cacheObj(ctx)
		for _, obj := range list {
			cache.Set(obj.ID, obj)
		}
	}
}

func (r *QueryResolver) cacheObj(ctx context.Context) ctxcache.Cacher {
	return ctxcache.GetCache(ctx, "Zone")
}
