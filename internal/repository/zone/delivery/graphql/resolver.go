package graphql

import (
	"context"
	"fmt"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/internal/repository/zone"
	"github.com/sspserver/api/internal/repository/zone/usecase"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qlmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc zone.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get is the resolver for the zone field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qlmodels.ZonePayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &qlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		Zone:             qlmodels.FromZoneModel(obj),
	}, nil
}

// List Zones is the resolver for the listApplications field.
func (r *QueryResolver) List(ctx context.Context, filter *qlmodels.ZoneListFilter, order *qlmodels.ZoneListOrder, page *qlmodels.Page) (*connectors.ZoneConnection, error) {
	return connectors.NewZoneConnection(ctx, r.uc, filter, order, page), nil
}

// Create Zone is the resolver for the createApplication field.
func (r *QueryResolver) Create(ctx context.Context, input qlmodels.ZoneInput) (*qlmodels.ZonePayload, error) {
	var obj models.Zone
	input.FillModel(&obj)

	id, err := r.uc.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}

	return &qlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             qlmodels.FromZoneModel(&obj),
	}, nil
}

// Update Zone is the resolver for the updateApplication field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qlmodels.ZoneInput) (*qlmodels.ZonePayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if obj == nil {
		return nil, fmt.Errorf("zone not found")
	}

	input.FillModel(obj)
	if err = r.uc.Update(ctx, id, obj); err != nil {
		return nil, err
	}

	return &qlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             qlmodels.FromZoneModel(obj),
	}, nil
}

// Delete Zone is the resolver for the deleteApplication field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qlmodels.ZonePayload, error) {
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

	return &qlmodels.ZonePayload{
		ClientMutationID: requestid.Get(ctx),
		ZoneID:           id,
		Zone:             qlmodels.FromZoneModel(obj),
	}, nil
}

// Run Zone is the resolver for the runApplication field.
func (r *QueryResolver) Run(ctx context.Context, id uint64, msg *string) (*qlmodels.StatusResponse, error) {
	err := r.uc.Run(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	return &qlmodels.StatusResponse{
		ClientMutationID: requestid.Get(ctx),
		Status:           qlmodels.ResponseStatusSuccess,
	}, nil
}

// Pause Zone is the resolver for the pauseApplication field.
func (r *QueryResolver) Pause(ctx context.Context, id uint64, msg *string) (*qlmodels.StatusResponse, error) {
	err := r.uc.Pause(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	return &qlmodels.StatusResponse{
		ClientMutationID: requestid.Get(ctx),
		Status:           qlmodels.ResponseStatusSuccess,
	}, nil
}

// Approve Zone is the resolver for the approveApplication field.
func (r *QueryResolver) Approve(ctx context.Context, id uint64, msg *string) (*qlmodels.StatusResponse, error) {
	err := r.uc.Approve(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	return &qlmodels.StatusResponse{
		ClientMutationID: requestid.Get(ctx),
		Status:           qlmodels.ResponseStatusSuccess,
	}, nil
}

// Reject Zone is the resolver for the rejectApplication field.
func (r *QueryResolver) Reject(ctx context.Context, id uint64, msg *string) (*qlmodels.StatusResponse, error) {
	err := r.uc.Reject(ctx, id, gocast.PtrAsValue(msg, ""))
	if err != nil {
		return nil, err
	}
	return &qlmodels.StatusResponse{
		ClientMutationID: requestid.Get(ctx),
		Status:           qlmodels.ResponseStatusSuccess,
	}, nil
}
