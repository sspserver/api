package graphql

import (
	"context"
	"fmt"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/internal/context/ctxcache"
	"github.com/sspserver/api/internal/repository/adformat"
	"github.com/sspserver/api/internal/repository/adformat/usecase"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc adformat.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get RTBAccessPoint is the resolver for the RTBAccessPoint field.
func (r *QueryResolver) Get(ctx context.Context, id uint64, codename string) (*qmodels.AdFormatPayload, error) {
	var (
		err error
		obj any
	)
	if id == 0 {
		obj, err = ctxcache.GetCache(ctx, "AdFormat").GetOrCache("code_"+codename, func(key any) (any, error) {
			f, err := r.uc.GetByCodename(ctx, key.(string))
			if err != nil {
				return nil, err
			}
			return f, nil
		})
	} else {
		obj, err = ctxcache.GetCache(ctx, "AdFormat").GetOrCache(id, func(key any) (any, error) {
			f, err := r.uc.Get(ctx, key.(uint64))
			if err != nil {
				return nil, err
			}
			return f, nil
		})
	}
	if err != nil {
		return nil, err
	}
	object := gocast.IfThenExec(obj != nil,
		func() *models.Format { return obj.(*models.Format) },
		func() *models.Format { return nil })
	return &qmodels.AdFormatPayload{
		ClientMutationID: requestid.Get(ctx),
		FormatID:         gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Format:           qmodels.FromAdFormatModel(object),
	}, nil
}

// List RTBAccessPoints is the resolver for the listRTBAccessPoints field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.AdFormatListFilter, order *qmodels.AdFormatListOrder, page *qmodels.Page) (*connectors.AdFormatConnection, error) {
	return connectors.NewAdFormatConnection(ctx, r.uc, filter, order, page), nil
}

// Create RTBAccessPoint is the resolver for the createRTBAccessPoint field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.AdFormatInput) (*qmodels.AdFormatPayload, error) {
	var object models.Format
	input.FillModel(&object)

	id, err := r.uc.Create(ctx, &object)
	if err != nil {
		return nil, err
	}

	return &qmodels.AdFormatPayload{
		ClientMutationID: requestid.Get(ctx),
		FormatID:         id,
		Format:           qmodels.FromAdFormatModel(&object),
	}, nil
}

// Update RTBAccessPoint is the resolver for the updateRTBAccessPoint field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.AdFormatInput) (*qmodels.AdFormatPayload, error) {
	object, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if object == nil {
		return nil, fmt.Errorf("AdFromat not found")
	}

	input.FillModel(object)

	if err = r.uc.Update(ctx, id, object); err != nil {
		return nil, err
	}

	return &qmodels.AdFormatPayload{
		ClientMutationID: requestid.Get(ctx),
		FormatID:         gocast.IfThenExec(object != nil, func() uint64 { return object.ID }, func() uint64 { return 0 }),
		Format:           qmodels.FromAdFormatModel(object),
	}, nil
}

// Delete RTBAccessPoint is the resolver for the deleteRTBAccessPoint field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, codename string, msg *string) (*qmodels.AdFormatPayload, error) {
	var (
		err    error
		object *models.Format
	)
	if id == 0 {
		object, err = r.uc.GetByCodename(ctx, codename)
	} else {
		object, err = r.uc.Get(ctx, id)
	}
	if err != nil {
		return nil, err
	}
	if object == nil {
		return nil, fmt.Errorf("AdFormat not found")
	}
	if id == 0 {
		err = r.uc.DeleteByCodename(ctx, codename, msg)
	} else {
		err = r.uc.Delete(ctx, id, msg)
	}
	if err != nil {
		return nil, err
	}
	return &qmodels.AdFormatPayload{
		ClientMutationID: requestid.Get(ctx),
		FormatID:         object.ID,
		Format:           qmodels.FromAdFormatModel(object),
	}, nil
}
