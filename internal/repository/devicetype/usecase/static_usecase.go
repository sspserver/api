package usecase

import (
	"context"

	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/pkg/errors"

	"github.com/sspserver/api/internal/repository/devicetype"
	"github.com/sspserver/api/internal/repository/devicetype/repository"
	"github.com/sspserver/api/models"
)

type StaticUsecase struct {
	repo devicetype.Repository
}

func NewStaticUsecase() *StaticUsecase {
	return &StaticUsecase{
		repo: repository.NewStaticRepository(),
	}
}

func (u *StaticUsecase) Get(ctx context.Context, id uint64) (*models.DeviceType, error) {
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if !acl.HaveAccessView(ctx, obj) {
		return nil, errors.Wrap(acl.ErrNoPermissions, "view")
	}
	return obj, nil
}

func (u *StaticUsecase) FetchList(ctx context.Context) ([]*models.DeviceType, error) {
	if !acl.HaveAccessList(ctx, &models.DeviceType{}) {
		return nil, errors.Wrap(acl.ErrNoPermissions, "fetch list")
	}
	return u.repo.FetchList(ctx)
}

func (u *StaticUsecase) FetchListByIDs(ctx context.Context, ids []uint64) ([]*models.DeviceType, error) {
	if !acl.HaveAccessList(ctx, &models.DeviceType{}) {
		return nil, errors.Wrap(acl.ErrNoPermissions, "fetch list")
	}
	return u.repo.FetchListByIDs(ctx, ids)
}

func (u *StaticUsecase) Count(ctx context.Context) (int64, error) {
	if !acl.HaveAccessCount(ctx, &models.DeviceType{}) {
		return 0, errors.Wrap(acl.ErrNoPermissions, "count")
	}
	return u.repo.Count(ctx)
}
