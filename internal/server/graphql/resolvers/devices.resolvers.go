package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"

	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	models1 "github.com/geniusrabbit/blaze-api/server/graphql/models"
	"github.com/sspserver/api/internal/server/graphql/models"
)

// CreateDeviceMaker is the resolver for the createDeviceMaker field.
func (r *mutationResolver) CreateDeviceMaker(ctx context.Context, input models.DeviceMakerInput) (*models.DeviceMakerPayload, error) {
	return r.device_makers.Create(ctx, input)
}

// UpdateDeviceMaker is the resolver for the updateDeviceMaker field.
func (r *mutationResolver) UpdateDeviceMaker(ctx context.Context, id uint64, input models.DeviceMakerInput) (*models.DeviceMakerPayload, error) {
	return r.device_makers.Update(ctx, id, input)
}

// DeleteDeviceMaker is the resolver for the deleteDeviceMaker field.
func (r *mutationResolver) DeleteDeviceMaker(ctx context.Context, id uint64, msg *string) (*models.DeviceMakerPayload, error) {
	return r.device_makers.Delete(ctx, id, msg)
}

// CreateDeviceModel is the resolver for the createDeviceModel field.
func (r *mutationResolver) CreateDeviceModel(ctx context.Context, input models.DeviceModelInput) (*models.DeviceModelPayload, error) {
	return r.device_models.Create(ctx, input)
}

// UpdateDeviceModel is the resolver for the updateDeviceModel field.
func (r *mutationResolver) UpdateDeviceModel(ctx context.Context, id uint64, input models.DeviceModelInput) (*models.DeviceModelPayload, error) {
	return r.device_models.Update(ctx, id, input)
}

// DeleteDeviceModel is the resolver for the deleteDeviceModel field.
func (r *mutationResolver) DeleteDeviceModel(ctx context.Context, id uint64, msg *string) (*models.DeviceModelPayload, error) {
	return r.device_models.Delete(ctx, id, msg)
}

// ListDeviceTypes is the resolver for the listDeviceTypes field.
func (r *queryResolver) ListDeviceTypes(ctx context.Context) ([]*models.DeviceType, error) {
	return r.device_types.List(ctx)
}

// DeviceMaker is the resolver for the deviceMaker field.
func (r *queryResolver) DeviceMaker(ctx context.Context, id uint64) (*models.DeviceMakerPayload, error) {
	return r.device_makers.Get(ctx, id)
}

// ListDeviceMakers is the resolver for the listDeviceMakers field.
func (r *queryResolver) ListDeviceMakers(ctx context.Context, filter *models.DeviceMakerListFilter, order *models.DeviceMakerListOrder, page *models1.Page) (*connectors.CollectionConnection[models.DeviceMaker, models.DeviceMakerEdge], error) {
	return r.device_makers.List(ctx, filter, order, page)
}

// DeviceModel is the resolver for the deviceModel field.
func (r *queryResolver) DeviceModel(ctx context.Context, id uint64) (*models.DeviceModelPayload, error) {
	return r.device_models.Get(ctx, id)
}

// ListDeviceModels is the resolver for the listDeviceModels field.
func (r *queryResolver) ListDeviceModels(ctx context.Context, filter *models.DeviceModelListFilter, order *models.DeviceModelListOrder, page *models1.Page) (*connectors.CollectionConnection[models.DeviceModel, models.DeviceModelEdge], error) {
	return r.device_models.List(ctx, filter, order, page)
}
