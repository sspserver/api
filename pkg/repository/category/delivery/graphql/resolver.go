package graphql

import (
	"context"
	"fmt"
	"slices"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"
	"github.com/geniusrabbit/blaze-api/repository"

	"github.com/sspserver/api/pkg/context/ctxcache"
	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/category"
	"github.com/sspserver/api/pkg/repository/category/usecase"
	qmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
	uc category.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get Category is the resolver for the Category field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qmodels.CategoryPayload, error) {
	object, err := r.cachedItemByID(ctx, id, true)
	if err != nil {
		return nil, err
	}
	return &qmodels.CategoryPayload{
		ClientMutationID: requestid.Get(ctx),
		CategoryID:       gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Category:         FromCategoryModel(object),
	}, nil
}

// List Categorys is the resolver for the listCategorys field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.CategoryListFilter, order *qmodels.CategoryListOrder, page *qmodels.Page) (*CategoryConnection, error) {
	return NewCategoryConnection(ctx, r.uc, filter, order, page), nil
}

// ListByIDs returns a list of DeviceModels by their IDs.
func (r *QueryResolver) ListByIDs(ctx context.Context, ids []uint64) ([]*qmodels.Category, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var (
		list   = make([]*models.Category, 0, len(ids))
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
		newList, err := r.uc.FetchList(ctx, &category.Filter{ID: newIDs})
		if err != nil {
			return nil, err
		}
		cache := r.cacheObj(ctx)
		for _, obj := range newList {
			cache.Set(obj.ID, obj)
		}
		list = append(list, newList...)
	}
	return FromCategoryModelList(list), nil
}

// Childrens is the resolver for the childrens field.
func (r *QueryResolver) Childrens(ctx context.Context, obj *qmodels.Category) ([]*qmodels.Category, error) {
	if obj == nil {
		return nil, nil
	}
	if len(obj.Childrens) > 0 {
		return obj.Childrens, nil
	}
	list, err := r.uc.FetchList(ctx,
		&repository.PreloadOption{
			Fields: []string{`Childrens`},
		},
		&category.Filter{
			ParentID: []uint64{obj.ID},
			Active:   &[]models.ActiveStatus{models.StatusActive}[0],
		},
		&category.ListOrder{Position: models.OrderAsc},
	)
	if err != nil {
		return nil, err
	}
	return FromCategoryModelList(list), nil
}

// Create Category is the resolver for the createCategory field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.CategoryInput) (*qmodels.CategoryPayload, error) {
	var object models.Category
	FillCategoryInputModel(input, &object)

	id, err := r.uc.Create(ctx, &object)
	if err != nil {
		return nil, err
	}

	return &qmodels.CategoryPayload{
		ClientMutationID: requestid.Get(ctx),
		CategoryID:       id,
		Category:         FromCategoryModel(&object),
	}, nil
}

// Update Category is the resolver for the updateCategory field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.CategoryInput) (*qmodels.CategoryPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if object == nil {
		return nil, fmt.Errorf("category not found")
	}

	FillCategoryInputModel(input, object)
	if err = r.uc.Update(ctx, id, object); err != nil {
		return nil, err
	}

	return &qmodels.CategoryPayload{
		ClientMutationID: requestid.Get(ctx),
		CategoryID:       gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Category:         FromCategoryModel(object),
	}, nil
}

// Delete Category is the resolver for the deleteCategory field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qmodels.CategoryPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, fmt.Errorf("category not found")
	}
	if err := r.uc.Delete(ctx, id, msg); err != nil {
		return nil, err
	}
	return &qmodels.CategoryPayload{
		ClientMutationID: requestid.Get(ctx),
		CategoryID:       object.ID,
		Category:         FromCategoryModel(object),
	}, nil
}

func (r *QueryResolver) cachedItemByID(ctx context.Context, id uint64, orLoad bool) (*models.Category, error) {
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
	return obj.(*models.Category), nil
}

func (r *QueryResolver) cacheObj(ctx context.Context) ctxcache.Cacher {
	return ctxcache.GetCache(ctx, "Category")
}
