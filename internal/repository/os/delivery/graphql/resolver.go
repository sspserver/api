package graphql

import (
	"context"
	"fmt"

	"github.com/geniusrabbit/blaze-api/pkg/requestid"

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
	obj, err := r.uc.Get(ctx, id)
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
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.OSListFilter, order *qmodels.OSListOrder, page *qmodels.Page) (*connectors.OSConnection, error) {
	return connectors.NewOSConnection(ctx, r.uc, filter, order, page), nil
}

// Create OS is the resolver for the createOS field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.OSInput) (*qmodels.OSPayload, error) {
	var obj models.OS
	input.FillModel(&obj)

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
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.OSInput) (*qmodels.OSPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, fmt.Errorf("OS not found")
	}
	input.FillModel(obj)
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
