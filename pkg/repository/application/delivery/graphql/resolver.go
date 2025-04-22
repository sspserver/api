package graphql

import (
	"context"
	"fmt"
	"slices"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/sspserver/api/pkg/context/ctxcache"
	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/application"
	"github.com/sspserver/api/pkg/repository/application/usecase"
	"github.com/sspserver/api/pkg/server/graphql/connectors"
	qlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
	uc application.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get is the resolver for the application field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qlmodels.ApplicationPayload, error) {
	obj, err := r.cachedItemByID(ctx, id, true)
	if err != nil {
		return nil, err
	}
	return &qlmodels.ApplicationPayload{
		ClientMutationID: requestid.Get(ctx),
		ApplicationID:    gocast.IfThenExec(obj != nil, func() uint64 { return obj.ID }, func() uint64 { return 0 }),
		Application:      qlmodels.FromApplicationModel(obj),
	}, nil
}

// List Applications is the resolver for the listApplications field.
func (r *QueryResolver) List(ctx context.Context, filter *qlmodels.ApplicationListFilter, order *qlmodels.ApplicationListOrder, page *qlmodels.Page) (*connectors.ApplicationConnection, error) {
	return connectors.NewApplicationConnection(ctx, r.uc, filter, order, page), nil
}

func (r *QueryResolver) ListByIDs(ctx context.Context, ids []uint64) ([]*qlmodels.Application, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var (
		list   = make([]*models.Application, 0, len(ids))
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
		newList, err := r.uc.FetchList(ctx, &application.Filter{ID: newIDs}, nil, nil)
		if err != nil {
			return nil, err
		}
		cache := r.cacheObj(ctx)
		for _, obj := range newList {
			cache.Set(obj.ID, obj)
		}
		list = append(list, newList...)
	}
	return qlmodels.FromApplicationModelList(list), nil
}

// Create Application is the resolver for the createApplication field.
func (r *QueryResolver) Create(ctx context.Context, input qlmodels.ApplicationCreateInput) (*qlmodels.ApplicationPayload, error) {
	var obj models.Application
	if err := input.FillModel(&obj); err != nil {
		return nil, err
	}

	if obj.URI == "" {
		return nil, &gqlerror.Error{
			Message: "URI is required",
			Extensions: map[string]any{
				"code":     "validation",
				"required": true,
				"field":    "URI",
			},
		}
	}

	id, err := r.uc.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}

	return &qlmodels.ApplicationPayload{
		ClientMutationID: requestid.Get(ctx),
		ApplicationID:    id,
		Application:      qlmodels.FromApplicationModel(&obj),
	}, nil
}

// Update Application is the resolver for the updateApplication field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qlmodels.ApplicationUpdateInput) (*qlmodels.ApplicationPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, fmt.Errorf("application not found")
	}

	if err = input.FillModel(obj); err != nil {
		return nil, err
	}

	if err = r.uc.Update(ctx, id, obj); err != nil {
		return nil, err
	}

	return &qlmodels.ApplicationPayload{
		ClientMutationID: requestid.Get(ctx),
		ApplicationID:    id,
		Application:      qlmodels.FromApplicationModel(obj),
	}, nil
}

// Delete Application is the resolver for the deleteApplication field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qlmodels.ApplicationPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, fmt.Errorf("application not found")
	}

	if err = r.uc.Delete(ctx, id, gocast.PtrAsValue(msg, "")); err != nil {
		return nil, err
	}

	return &qlmodels.ApplicationPayload{
		ClientMutationID: requestid.Get(ctx),
		ApplicationID:    id,
		Application:      qlmodels.FromApplicationModel(obj),
	}, nil
}

// Run Application is the resolver for the runApplication field.
func (r *QueryResolver) Run(ctx context.Context, id uint64, msg *string) (*qlmodels.ApplicationPayload, error) {
	err := r.uc.Run(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	app, _ := r.uc.Get(ctx, id)
	return &qlmodels.ApplicationPayload{
		ClientMutationID: requestid.Get(ctx),
		ApplicationID:    id,
		Application:      qlmodels.FromApplicationModel(app),
	}, nil
}

// Pause Application is the resolver for the pauseApplication field.
func (r *QueryResolver) Pause(ctx context.Context, id uint64, msg *string) (*qlmodels.ApplicationPayload, error) {
	err := r.uc.Pause(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	app, _ := r.uc.Get(ctx, id)
	return &qlmodels.ApplicationPayload{
		ClientMutationID: requestid.Get(ctx),
		ApplicationID:    id,
		Application:      qlmodels.FromApplicationModel(app),
	}, nil
}

// Approve Application is the resolver for the approveApplication field.
func (r *QueryResolver) Approve(ctx context.Context, id uint64, msg *string) (*qlmodels.ApplicationPayload, error) {
	err := r.uc.Approve(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	app, _ := r.uc.Get(ctx, id)
	return &qlmodels.ApplicationPayload{
		ClientMutationID: requestid.Get(ctx),
		ApplicationID:    id,
		Application:      qlmodels.FromApplicationModel(app),
	}, nil
}

// Reject Application is the resolver for the rejectApplication field.
func (r *QueryResolver) Reject(ctx context.Context, id uint64, msg *string) (*qlmodels.ApplicationPayload, error) {
	err := r.uc.Reject(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	app, _ := r.uc.Get(ctx, id)
	return &qlmodels.ApplicationPayload{
		ClientMutationID: requestid.Get(ctx),
		ApplicationID:    id,
		Application:      qlmodels.FromApplicationModel(app),
	}, nil
}

func (r *QueryResolver) cachedItemByID(ctx context.Context, id uint64, orLoad bool) (*models.Application, error) {
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
	return obj.(*models.Application), nil
}

func (r *QueryResolver) cacheObj(ctx context.Context) ctxcache.Cacher {
	return ctxcache.GetCache(ctx, "Application")
}
