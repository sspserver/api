package usecase

import (
	"context"

	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/repository/historylog"
	"github.com/go-faster/errors"

	"github.com/sspserver/api/internal/repository/devicemodel"
	"github.com/sspserver/api/internal/repository/devicemodel/repository"
	"github.com/sspserver/api/models"
)

type Usecase struct {
	repo devicemodel.Repository
}

// New create new usecase
func New() *Usecase {
	return &Usecase{
		repo: repository.New(),
	}
}

func (u *Usecase) Get(ctx context.Context, id uint64) (*models.DeviceModel, error) {
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if !acl.HaveAccessView(ctx, obj) {
		return nil, errors.Wrap(acl.ErrNoPermissions, "view")
	}
	return obj, nil
}

func (u *Usecase) GetByCodename(ctx context.Context, codename string) (*models.DeviceModel, error) {
	obj, err := u.repo.GetByCodename(ctx, codename)
	if err != nil {
		return nil, err
	}
	if !acl.HaveAccessView(ctx, obj) {
		return nil, errors.Wrap(acl.ErrNoPermissions, "view")
	}
	return obj, nil
}

func (u *Usecase) FetchList(ctx context.Context, qops ...devicemodel.Option) ([]*models.DeviceModel, error) {
	if !acl.HaveAccessList(ctx, &models.DeviceModel{}) {
		return nil, errors.Wrap(acl.ErrNoPermissions, "fetch list")
	}
	return u.repo.FetchList(ctx, qops...)
}

func (u *Usecase) Count(ctx context.Context, qops ...devicemodel.Option) (int64, error) {
	if !acl.HaveAccessList(ctx, &models.DeviceModel{}) {
		return 0, errors.Wrap(acl.ErrNoPermissions, "count")
	}
	return u.repo.Count(ctx, qops...)
}

func (u *Usecase) Create(ctx context.Context, object *models.DeviceModel) (uint64, error) {
	if !acl.HaveAccessCreate(ctx, object) {
		return 0, errors.Wrap(acl.ErrNoPermissions, "create")
	}
	return u.repo.Create(ctx, object)
}

func (u *Usecase) Update(ctx context.Context, id uint64, object *models.DeviceModel) error {
	if !acl.HaveAccessUpdate(ctx, object) {
		return errors.Wrap(acl.ErrNoPermissions, "update")
	}
	return u.repo.Update(ctx, id, object)
}

func (u *Usecase) Delete(ctx context.Context, id uint64, msg *string) error {
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessDelete(ctx, obj) {
		return errors.Wrap(acl.ErrNoPermissions, "delete")
	}
	if msg != nil {
		ctx = historylog.WithMessage(ctx, *msg)
	}
	return u.repo.Delete(ctx, id)
}
