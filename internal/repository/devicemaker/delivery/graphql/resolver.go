package graphql

import (
	"context"
	"fmt"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/internal/repository/devicemaker"
	"github.com/sspserver/api/internal/repository/devicemaker/usecase"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc devicemaker.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get DeviceMaker is the resolver for the DeviceMaker field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qmodels.DeviceMakerPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &qmodels.DeviceMakerPayload{
		ClientMutationID: requestid.Get(ctx),
		MakerID:          gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Maker:            qmodels.FromDeviceMakerModel(object),
	}, nil
}

// List DeviceMakers is the resolver for the listDeviceMakers field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.DeviceMakerListFilter, order []*qmodels.DeviceMakerListOrder, page *qmodels.Page) (*connectors.DeviceMakerConnection, error) {
	return connectors.NewDeviceMakerConnection(ctx, r.uc, filter, order, page), nil
}

// Create DeviceMaker is the resolver for the createDeviceMaker field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.DeviceMakerCreateInput) (*qmodels.DeviceMakerPayload, error) {
	var object models.DeviceMaker
	input.FillModel(&object)

	id, err := r.uc.Create(ctx, &object)
	if err != nil {
		return nil, err
	}
	return &qmodels.DeviceMakerPayload{
		ClientMutationID: requestid.Get(ctx),
		MakerID:          id,
		Maker:            qmodels.FromDeviceMakerModel(&object),
	}, nil
}

// Update DeviceMaker is the resolver for the updateDeviceMaker field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.DeviceMakerUpdateInput) (*qmodels.DeviceMakerPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, fmt.Errorf("DeviceMaker not found")
	}
	models := object.Models
	object.Models = nil
	input.FillModel(object)
	if err = r.uc.Update(ctx, id, object); err != nil {
		return nil, err
	}
	object.Models = models
	return &qmodels.DeviceMakerPayload{
		ClientMutationID: requestid.Get(ctx),
		MakerID:          object.ID,
		Maker:            qmodels.FromDeviceMakerModel(object),
	}, nil
}

// Delete DeviceMaker is the resolver for the deleteDeviceMaker field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qmodels.DeviceMakerPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, fmt.Errorf("DeviceMaker not found")
	}
	if err := r.uc.Delete(ctx, id, msg); err != nil {
		return nil, err
	}
	return &qmodels.DeviceMakerPayload{
		ClientMutationID: requestid.Get(ctx),
		MakerID:          object.ID,
		Maker:            qmodels.FromDeviceMakerModel(object),
	}, nil
}
