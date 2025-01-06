package graphql

import (
	"context"
	"fmt"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/internal/repository/devicemodel"
	"github.com/sspserver/api/internal/repository/devicemodel/usecase"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc devicemodel.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get DeviceModel is the resolver for the DeviceModel field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qmodels.DeviceModelPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &qmodels.DeviceModelPayload{
		ClientMutationID: requestid.Get(ctx),
		ModelID:          gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Model:            qmodels.FromDeviceModelModel(object),
	}, nil
}

// List DeviceModel is the resolver for the listDeviceModels field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.DeviceModelListFilter, order *qmodels.DeviceModelListOrder, page *qmodels.Page) (*connectors.DeviceModelConnection, error) {
	return connectors.NewDeviceModelConnection(ctx, r.uc, filter, order, page), nil
}

// Create DeviceModel is the resolver for the createDeviceModel field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.DeviceModelInput) (*qmodels.DeviceModelPayload, error) {
	var object models.DeviceModel
	input.FillModel(&object)

	id, err := r.uc.Create(ctx, &object)
	if err != nil {
		return nil, err
	}

	// Get object to return
	nObject, err := r.uc.Get(ctx, id)
	fmt.Println(">>> CREATE", nObject, err)
	if err != nil {
		nObject = &object
	}

	return &qmodels.DeviceModelPayload{
		ClientMutationID: requestid.Get(ctx),
		ModelID:          id,
		Model:            qmodels.FromDeviceModelModel(nObject),
	}, nil
}

// Update DeviceModel is the resolver for the updateDeviceModel field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.DeviceModelInput) (*qmodels.DeviceModelPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if object == nil {
		return nil, fmt.Errorf("DeviceModel not found")
	}

	input.FillModel(object)
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
		Model:            qmodels.FromDeviceModelModel(nObject),
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
		Model:            qmodels.FromDeviceModelModel(object),
	}, nil
}
