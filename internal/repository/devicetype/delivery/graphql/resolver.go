package graphql

import (
	"context"

	"github.com/sspserver/api/internal/repository/devicetype"
	"github.com/sspserver/api/internal/repository/devicetype/repository"
	qlmodels "github.com/sspserver/api/internal/server/graphql/models"
)

type QueryResolver struct {
	repo devicetype.Usecase
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{
		repo: repository.NewStaticRepository(),
	}
}

// List DeviceTypes is the resolver for the listDeviceTypes field.
func (r *QueryResolver) List(ctx context.Context) ([]*qlmodels.DeviceType, error) {
	list, err := r.repo.FetchList(ctx)
	if err != nil {
		return nil, err
	}
	return qlmodels.FromDeviceTypeModelList(list), nil
}
