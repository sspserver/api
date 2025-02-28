package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"

	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	models1 "github.com/geniusrabbit/blaze-api/server/graphql/models"
	"github.com/sspserver/api/internal/server/graphql/generated"
	"github.com/sspserver/api/internal/server/graphql/models"
)

// Versions is the resolver for the versions field.
func (r *deviceModelResolver) Versions(ctx context.Context, obj *models.DeviceModel, filter *models.DeviceModelListFilter, order []*models.DeviceModelListOrder) ([]*models.DeviceModel, error) {
	if filter == nil {
		filter = &models.DeviceModelListFilter{
			Active: []models.ActiveStatus{models.ActiveStatusActive},
		}
	}
	filter.ParentID = []uint64{obj.ID}
	coll, err := r.device_models.List(ctx, filter, order, nil)
	if err != nil {
		return nil, err
	}
	return coll.List(), nil
}

// CreateDeviceModel is the resolver for the createDeviceModel field.
func (r *mutationResolver) CreateDeviceModel(ctx context.Context, input models.DeviceModelCreateInput) (*models.DeviceModelPayload, error) {
	return r.device_models.Create(ctx, input)
}

// UpdateDeviceModel is the resolver for the updateDeviceModel field.
func (r *mutationResolver) UpdateDeviceModel(ctx context.Context, id uint64, input models.DeviceModelUpdateInput) (*models.DeviceModelPayload, error) {
	return r.device_models.Update(ctx, id, input)
}

// DeleteDeviceModel is the resolver for the deleteDeviceModel field.
func (r *mutationResolver) DeleteDeviceModel(ctx context.Context, id uint64, msg *string) (*models.DeviceModelPayload, error) {
	return r.device_models.Delete(ctx, id, msg)
}

// DeviceModel is the resolver for the deviceModel field.
func (r *queryResolver) DeviceModel(ctx context.Context, id uint64, codename string) (*models.DeviceModelPayload, error) {
	return r.device_models.Get(ctx, id, codename)
}

// ListDeviceModels is the resolver for the listDeviceModels field.
func (r *queryResolver) ListDeviceModels(ctx context.Context, filter *models.DeviceModelListFilter, order []*models.DeviceModelListOrder, page *models1.Page) (*connectors.CollectionConnection[models.DeviceModel, models.DeviceModelEdge], error) {
	return r.device_models.List(ctx, filter, order, page)
}

// DeviceModel returns generated.DeviceModelResolver implementation.
func (r *Resolver) DeviceModel() generated.DeviceModelResolver { return &deviceModelResolver{r} }

type deviceModelResolver struct{ *Resolver }
