package graphql

import (
	"context"
	"fmt"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/internal/repository/category"
	"github.com/sspserver/api/internal/repository/category/usecase"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc category.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get Category is the resolver for the Category field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qmodels.CategoryPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &qmodels.CategoryPayload{
		ClientMutationID: requestid.Get(ctx),
		CategoryID:       gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Category:         qmodels.FromCategoryModel(object),
	}, nil
}

// List Categorys is the resolver for the listCategorys field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.CategoryListFilter, order *qmodels.CategoryListOrder, page *qmodels.Page) (*connectors.CategoryConnection, error) {
	return connectors.NewCategoryConnection(ctx, r.uc, filter, order, page), nil
}

// Create Category is the resolver for the createCategory field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.CategoryInput) (*qmodels.CategoryPayload, error) {
	var object models.Category
	input.FillModel(&object)

	id, err := r.uc.Create(ctx, &object)
	if err != nil {
		return nil, err
	}

	return &qmodels.CategoryPayload{
		ClientMutationID: requestid.Get(ctx),
		CategoryID:       id,
		Category:         qmodels.FromCategoryModel(&object),
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

	input.FillModel(object)
	if err = r.uc.Update(ctx, id, object); err != nil {
		return nil, err
	}

	return &qmodels.CategoryPayload{
		ClientMutationID: requestid.Get(ctx),
		CategoryID:       gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Category:         qmodels.FromCategoryModel(object),
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
		Category:         qmodels.FromCategoryModel(object),
	}, nil
}
