package graphql

import (
	"context"
	"fmt"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/trafficrouter"
	"github.com/sspserver/api/pkg/server/graphql/connectors"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type Resolver struct {
	uc trafficrouter.Usecase
}

func NewResolver() *Resolver {
	return &Resolver{
		uc: trafficrouter.NewUsecaseImpl(
			trafficrouter.NewRepositoryImpl(),
		),
	}
}

// Get is the resolver for the trafficRouter field.
func (r *Resolver) Get(ctx context.Context, id uint64) (*gqlmodels.TrafficRouterPayload, error) {
	tr, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &gqlmodels.TrafficRouterPayload{
		ClientMutationID: requestid.Get(ctx),
		RouterID:         tr.ID,
		Router:           gqlmodels.FromTrafficRouterModel(tr),
	}, nil
}

// List is the resolver for the trafficRouters field.
func (r *Resolver) List(ctx context.Context, filter *gqlmodels.TrafficRouterListFilter, order []*gqlmodels.TrafficRouterListOrder, page *gqlmodels.Page) (*connectors.TrafficRouterConnection, error) {
	return connectors.NewTrafficRouterConnection(ctx, r.uc, filter, order, page), nil
}

// Create TrafficRouter is the resolver for the createTrafficRouter field.
func (r *Resolver) Create(ctx context.Context, input gqlmodels.TrafficRouterCreateInput) (*gqlmodels.TrafficRouterPayload, error) {
	var object models.TrafficRouter
	input.FillModel(&object)

	id, err := r.uc.Create(ctx, &object)
	if err != nil {
		return nil, err
	}

	return &gqlmodels.TrafficRouterPayload{
		ClientMutationID: requestid.Get(ctx),
		RouterID:         id,
		Router:           gqlmodels.FromTrafficRouterModel(&object),
	}, nil
}

// Update TrafficRouter is the resolver for the updateTrafficRouter field.
func (r *Resolver) Update(ctx context.Context, id uint64, input gqlmodels.TrafficRouterUpdateInput) (*gqlmodels.TrafficRouterPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if object == nil {
		return nil, fmt.Errorf("TrafficRouter not found")
	}

	input.FillModel(object)
	if err = r.uc.Update(ctx, id, object); err != nil {
		return nil, err
	}

	return &gqlmodels.TrafficRouterPayload{
		ClientMutationID: requestid.Get(ctx),
		RouterID:         gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Router:           gqlmodels.FromTrafficRouterModel(object),
	}, nil
}

// Delete TrafficRouter is the resolver for the deleteTrafficRouter field.
func (r *Resolver) Delete(ctx context.Context, id uint64) (*gqlmodels.TrafficRouterPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, fmt.Errorf("TrafficRouter not found")
	}
	if err := r.uc.Delete(ctx, id); err != nil {
		return nil, err
	}
	return &gqlmodels.TrafficRouterPayload{
		ClientMutationID: requestid.Get(ctx),
		RouterID:         object.ID,
		Router:           gqlmodels.FromTrafficRouterModel(object),
	}, nil
}

// Run TrafficRouter is the resolver for the runTrafficRouter field.
func (r *Resolver) Run(ctx context.Context, id uint64, message string) (*gqlmodels.TrafficRouterPayload, error) {
	if err := r.uc.Run(ctx, id, message); err != nil {
		return nil, err
	}
	router, _ := r.uc.Get(ctx, id)
	return &gqlmodels.TrafficRouterPayload{
		ClientMutationID: requestid.Get(ctx),
		RouterID:         id,
		Router:           gqlmodels.FromTrafficRouterModel(router),
	}, nil
}

// Pause TrafficRouter is the resolver for the pauseTrafficRouter field.
func (r *Resolver) Pause(ctx context.Context, id uint64, message string) (*gqlmodels.TrafficRouterPayload, error) {
	if err := r.uc.Pause(ctx, id, message); err != nil {
		return nil, err
	}
	router, _ := r.uc.Get(ctx, id)
	return &gqlmodels.TrafficRouterPayload{
		ClientMutationID: requestid.Get(ctx),
		RouterID:         id,
		Router:           gqlmodels.FromTrafficRouterModel(router),
	}, nil
}
