package usecase

import (
	"context"

	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/pkg/context/session"
	"github.com/geniusrabbit/blaze-api/repository/historylog"
	"github.com/go-faster/errors"

	"github.com/sspserver/api/internal/repository/rtbsource"
	"github.com/sspserver/api/internal/repository/rtbsource/repository"
	"github.com/sspserver/api/models"
)

type Usecase struct {
	repo rtbsource.Repository
}

// New create new usecase
func New() *Usecase {
	return &Usecase{
		repo: repository.New(),
	}
}

func (u *Usecase) Get(ctx context.Context, id uint64) (*models.RTBSource, error) {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if !acl.HaveAccessView(ctx, src) {
		return nil, errors.Wrap(acl.ErrNoPermissions, "view")
	}
	return src, nil
}

func (u *Usecase) FetchList(ctx context.Context, qops ...rtbsource.Option) ([]*models.RTBSource, error) {
	if !acl.HaveAccessList(ctx, &models.RTBSource{}) {
		accountID := session.Account(ctx).ID
		if acl.HaveAccessList(ctx, &models.RTBSource{AccountID: accountID}) {
			qops = rtbsource.Options(qops).With(&rtbsource.Filter{
				AccountID: accountID,
			})
		} else {
			return nil, errors.Wrap(acl.ErrNoPermissions, "fetch list")
		}
	}
	return u.repo.FetchList(ctx, qops...)
}

func (u *Usecase) Count(ctx context.Context, qops ...rtbsource.Option) (int64, error) {
	if !acl.HaveAccessList(ctx, &models.RTBSource{}) {
		accountID := session.Account(ctx).ID
		if acl.HaveAccessList(ctx, &models.RTBSource{AccountID: accountID}) {
			qops = rtbsource.Options(qops).With(&rtbsource.Filter{
				AccountID: accountID,
			})
		} else {
			return 0, errors.Wrap(acl.ErrNoPermissions, "fetch list")
		}
	}
	return u.repo.Count(ctx, qops...)
}

func (u *Usecase) Create(ctx context.Context, source *models.RTBSource) (uint64, error) {
	if source.AccountID == 0 {
		source.AccountID = session.Account(ctx).ID
	}
	if !acl.HaveAccessCreate(ctx, source) {
		return 0, errors.Wrap(acl.ErrNoPermissions, "create")
	}
	return u.repo.Create(ctx, source)
}

func (u *Usecase) Update(ctx context.Context, id uint64, source *models.RTBSource) error {
	srcObj := *source
	srcObj.ID = id
	if !acl.HaveAccessUpdate(ctx, srcObj) {
		return errors.Wrap(acl.ErrNoPermissions, "update")
	}
	return u.repo.Update(ctx, id, source)
}

func (u *Usecase) Run(ctx context.Context, id uint64, message string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessUpdate(ctx, src) {
		return errors.Wrap(acl.ErrNoPermissions, "update::run")
	}
	return u.repo.Run(ctx, id, message)
}

func (u *Usecase) Pause(ctx context.Context, id uint64, message string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessUpdate(ctx, src) {
		return errors.Wrap(acl.ErrNoPermissions, "update::pause")
	}
	return u.repo.Pause(ctx, id, message)
}

func (u *Usecase) Approve(ctx context.Context, id uint64, message string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveObjectPermissions(ctx, src, acl.PermApprove+`.*`) {
		return errors.Wrap(acl.ErrNoPermissions, "approve")
	}
	return u.repo.Approve(ctx, id, message)
}

func (u *Usecase) Reject(ctx context.Context, id uint64, message string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveObjectPermissions(ctx, src, acl.PermReject+`.*`) {
		return errors.Wrap(acl.ErrNoPermissions, "reject")
	}
	return u.repo.Reject(ctx, id, message)
}

func (u *Usecase) Delete(ctx context.Context, id uint64, msg *string) error {
	src, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessDelete(ctx, src) {
		return errors.Wrap(acl.ErrNoPermissions, "delete")
	}
	if msg != nil {
		ctx = historylog.WithMessage(ctx, *msg)
	}
	return u.repo.Delete(ctx, id)
}
