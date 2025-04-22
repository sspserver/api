package usecase

import (
	"context"

	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/pkg/context/session"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/zone"
	"github.com/sspserver/api/pkg/repository/zone/repository"
	"github.com/sspserver/api/pkg/sysops"
)

type Usecase struct {
	repo zone.Repository
}

// New create new usecase
func New() *Usecase {
	return &Usecase{
		repo: repository.New(),
	}
}

func (u *Usecase) Get(ctx context.Context, id uint64) (*models.Zone, error) {
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if !acl.HaveAccessView(ctx, obj) {
		return nil, acl.ErrNoPermissions.WithMessage("view")
	}
	return obj, nil
}

func (u *Usecase) GetByCodename(ctx context.Context, codename string) (*models.Zone, error) {
	obj, err := u.repo.GetByCodename(ctx, codename)
	if err != nil {
		return nil, err
	}
	if !acl.HaveAccessView(ctx, obj) {
		return nil, acl.ErrNoPermissions.WithMessage("view")
	}
	return obj, nil
}

func (u *Usecase) FetchList(ctx context.Context, qops ...zone.Option) ([]*models.Zone, error) {
	if !acl.HaveAccessList(ctx, &models.Application{}) {
		return nil, acl.ErrNoPermissions.WithMessage("fetch list")
	}
	return u.repo.FetchList(ctx, qops...)
}

func (u *Usecase) Count(ctx context.Context, qops ...zone.Option) (int64, error) {
	if !acl.HaveAccessList(ctx, &models.Application{}) {
		return 0, acl.ErrNoPermissions.WithMessage("count")
	}
	return u.repo.Count(ctx, qops...)
}

func (u *Usecase) Create(ctx context.Context, object *models.Zone) (uint64, error) {
	if object.AccountID == 0 {
		object.AccountID = session.Account(ctx).ID
	}
	if !acl.HaveAccessCreate(ctx, object) {
		return 0, acl.ErrNoPermissions.WithMessage("create")
	}
	object.Status = models.ApproveStatus(
		sysops.Get(`logic.crud.default.approval`, models.StatusPending).
			Int())
	return u.repo.Create(ctx, object)
}

func (u *Usecase) Update(ctx context.Context, id uint64, object *models.Zone) error {
	if !acl.HaveAccessUpdate(ctx, object) {
		return acl.ErrNoPermissions.WithMessage("update")
	}
	return u.repo.Update(ctx, id, object)
}

func (u *Usecase) Delete(ctx context.Context, id uint64, msg string) error {
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessDelete(ctx, obj) {
		return acl.ErrNoPermissions.WithMessage("delete")
	}
	return u.repo.Delete(ctx, id, msg)
}

func (u *Usecase) Run(ctx context.Context, id uint64, message string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessUpdate(ctx, src) {
		return acl.ErrNoPermissions.WithMessage("update::run")
	}
	return u.repo.Run(ctx, id, message)
}

func (u *Usecase) Pause(ctx context.Context, id uint64, message string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessUpdate(ctx, src) {
		return acl.ErrNoPermissions.WithMessage("update::pause")
	}
	return u.repo.Pause(ctx, id, message)
}

func (u *Usecase) Approve(ctx context.Context, id uint64, message string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveObjectPermissions(ctx, src, acl.PermApprove+`.*`) {
		return acl.ErrNoPermissions.WithMessage("approve")
	}
	return u.repo.Approve(ctx, id, message)
}

func (u *Usecase) Reject(ctx context.Context, id uint64, message string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveObjectPermissions(ctx, src, acl.PermReject+`.*`) {
		return acl.ErrNoPermissions.WithMessage("reject")
	}
	return u.repo.Reject(ctx, id, message)
}
