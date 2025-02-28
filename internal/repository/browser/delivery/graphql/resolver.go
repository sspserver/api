package graphql

import (
	"context"
	"fmt"

	"github.com/geniusrabbit/blaze-api/pkg/requestid"

	"github.com/sspserver/api/internal/repository/browser"
	"github.com/sspserver/api/internal/repository/browser/usecase"
	"github.com/sspserver/api/internal/server/graphql/connectors"
	qmodels "github.com/sspserver/api/internal/server/graphql/models"
	"github.com/sspserver/api/models"
)

type QueryResolver struct {
	uc browser.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{uc: usecase.New()}
}

// Get Browser is the resolver for the Browser field.
func (r *QueryResolver) Get(ctx context.Context, id uint64) (*qmodels.BrowserPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &qmodels.BrowserPayload{
		ClientMutationID: requestid.Get(ctx),
		BrowserID:        obj.ID,
		Browser:          qmodels.FromBrowserModel(obj),
	}, nil
}

// List Browser is the resolver for the listBrowser field.
func (r *QueryResolver) List(ctx context.Context, filter *qmodels.BrowserListFilter, order []*qmodels.BrowserListOrder, page *qmodels.Page) (*connectors.BrowserConnection, error) {
	return connectors.NewBrowserConnection(ctx, r.uc, filter, order, page), nil
}

// Create Browser is the resolver for the createBrowser field.
func (r *QueryResolver) Create(ctx context.Context, input qmodels.BrowserCreateInput) (*qmodels.BrowserPayload, error) {
	var obj models.Browser
	if err := input.FillModel(&obj); err != nil {
		return nil, err
	}

	id, err := r.uc.Create(ctx, &obj)
	if err != nil {
		return nil, err
	}

	return &qmodels.BrowserPayload{
		ClientMutationID: requestid.Get(ctx),
		BrowserID:        id,
		Browser:          qmodels.FromBrowserModel(&obj),
	}, nil
}

// Update Browser is the resolver for the updateBrowser field.
func (r *QueryResolver) Update(ctx context.Context, id uint64, input qmodels.BrowserUpdateInput) (*qmodels.BrowserPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, fmt.Errorf("browser not found")
	}
	if err = input.FillModel(obj); err != nil {
		return nil, err
	}
	if err := r.uc.Update(ctx, id, obj); err != nil {
		return nil, err
	}
	return &qmodels.BrowserPayload{
		ClientMutationID: requestid.Get(ctx),
		BrowserID:        id,
		Browser:          qmodels.FromBrowserModel(obj),
	}, nil
}

// Delete Browser is the resolver for the deleteBrowser field.
func (r *QueryResolver) Delete(ctx context.Context, id uint64, msg *string) (*qmodels.BrowserPayload, error) {
	obj, err := r.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if obj == nil {
		return nil, fmt.Errorf("browser not found")
	}
	if err := r.uc.Delete(ctx, id, msg); err != nil {
		return nil, err
	}
	return &qmodels.BrowserPayload{
		ClientMutationID: requestid.Get(ctx),
		BrowserID:        obj.ID,
		Browser:          qmodels.FromBrowserModel(obj),
	}, nil
}
