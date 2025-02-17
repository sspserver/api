package graphql

import (
	"context"
	"fmt"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/internal/context/ctxcache"
	"github.com/sspserver/api/internal/repository/rtbsource"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc rtbsource.Usecase
}

func NewQueryResolver(uc rtbsource.Usecase) *QueryResolver {
	return &QueryResolver{uc: uc}
}

// Get RTBSource is the resolver for the RTBSource field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qmodels.RTBSourcePayload, error) {
	source, err := ctxcache.GetCache(ctx, "RTBSource").GetOrCache(id, func(key any) (any, error) {
		return r.uc.Get(ctx, id)
	})
	if err != nil {
		return nil, err
	}
	src := gocast.IfThenExec(source != nil,
		func() *models.RTBSource { return source.(*models.RTBSource) },
		func() *models.RTBSource { return nil })
	return &qmodels.RTBSourcePayload{
		ClientMutationID: requestid.Get(ctx),
		SourceID:         gocast.IfThenExec(src != nil, func() uint64 { return src.ID }, func() uint64 { return 0 }),
		Source:           qmodels.FromRTBSourceModel(src),
	}, nil
}

// List RTBSources is the resolver for the listRTBSources field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.RTBSourceListFilter, order *qmodels.RTBSourceListOrder, page *qmodels.Page) (*connectors.RTBSourceConnection, error) {
	return connectors.NewRTBSourceConnection(ctx, r.uc, filter, order, page), nil
}

// Create RTBSource is the resolver for the createRTBSource field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.RTBSourceInput) (*qmodels.RTBSourcePayload, error) {
	var source models.RTBSource
	input.FillModel(&source)

	id, err := r.uc.Create(ctx, &source)
	if err != nil {
		return nil, err
	}

	return &qmodels.RTBSourcePayload{
		ClientMutationID: requestid.Get(ctx),
		SourceID:         id,
		Source:           qmodels.FromRTBSourceModel(&source),
	}, nil
}

// Update RTBSource is the resolver for the updateRTBSource field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.RTBSourceInput) (*qmodels.RTBSourcePayload, error) {
	source, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if source == nil {
		return nil, fmt.Errorf("RTBSource not found")
	}

	input.FillModel(source)
	if err = r.uc.Update(ctx, id, source); err != nil {
		return nil, err
	}

	return &qmodels.RTBSourcePayload{
		ClientMutationID: requestid.Get(ctx),
		SourceID:         gocast.IfThenExec(source != nil, func() uint64 { return source.ID }, func() uint64 { return 0 }),
		Source:           qmodels.FromRTBSourceModel(source),
	}, nil
}

// Delete RTBSource is the resolver for the deleteRTBSource field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qmodels.RTBSourcePayload, error) {
	source, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if source == nil {
		return nil, fmt.Errorf("RTBSource not found")
	}
	if err := r.uc.Delete(ctx, id, msg); err != nil {
		return nil, err
	}
	return &qmodels.RTBSourcePayload{
		ClientMutationID: requestid.Get(ctx),
		SourceID:         source.ID,
		Source:           qmodels.FromRTBSourceModel(source),
	}, nil
}

// Run RTBSource is the resolver for the runRTBSource field.
func (r *QueryResolver) Run(ctx context.Context, id uint64) (*qmodels.StatusResponse, error) {
	err := r.uc.Run(ctx, id, "")
	if err != nil {
		return nil, err
	}
	return &qmodels.StatusResponse{
		ClientMutationID: requestid.Get(ctx),
		Status:           qmodels.ResponseStatusSuccess,
	}, nil
}

// PauseRTBSource is the resolver for the pauseRTBSource field.
func (r *QueryResolver) Pause(ctx context.Context, id uint64) (*qmodels.StatusResponse, error) {
	err := r.uc.Pause(ctx, id, "")
	if err != nil {
		return nil, err
	}
	return &qmodels.StatusResponse{
		ClientMutationID: requestid.Get(ctx),
		Status:           qmodels.ResponseStatusSuccess,
	}, nil
}

// Approve RTBSource is the resolver for the approveRTBSource field.
func (r *QueryResolver) Approve(ctx context.Context, id uint64, msg *string) (*qmodels.StatusResponse, error) {
	err := r.uc.Approve(ctx, id,
		gocast.IfThenExec(msg != nil,
			func() string { return *msg },
			func() string { return "" },
		))
	if err != nil {
		return nil, err
	}
	return &qmodels.StatusResponse{
		ClientMutationID: requestid.Get(ctx),
		Status:           qmodels.ResponseStatusSuccess,
	}, nil
}

// Reject RTBSource is the resolver for the rejectRTBSource field.
func (r *QueryResolver) Reject(ctx context.Context, id uint64, msg *string) (*qmodels.StatusResponse, error) {
	err := r.uc.Reject(ctx, id,
		gocast.IfThenExec(msg != nil,
			func() string { return *msg },
			func() string { return "" },
		))
	if err != nil {
		return nil, err
	}
	return &qmodels.StatusResponse{
		ClientMutationID: requestid.Get(ctx),
		Status:           qmodels.ResponseStatusSuccess,
	}, nil
}
