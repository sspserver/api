package usecase

import (
	"context"

	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/pkg/context/session"

	"github.com/sspserver/api/internal/repository/application"
	"github.com/sspserver/api/internal/repository/application/repository"
	"github.com/sspserver/api/internal/sysops"
	"github.com/sspserver/api/models"
)

type Usecase struct {
	repo application.Repository
}

// New create new usecase
func New() *Usecase {
	return &Usecase{
		repo: repository.New(),
	}
}

// Get application object by id
func (u *Usecase) Get(ctx context.Context, id uint64) (*models.Application, error) {
	obj, err := u.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if !acl.HaveAccessView(ctx, obj) {
		return nil, acl.ErrNoPermissions.WithMessage("view")
	}
	return obj, nil
}

// FetchList application objects
func (u *Usecase) FetchList(ctx context.Context, qops ...application.Option) ([]*models.Application, error) {
	if !acl.HaveAccessList(ctx, &models.Application{}) {
		if !acl.HaveAccessList(ctx, &models.Application{AccountID: session.Account(ctx).ID}) {
			return nil, acl.ErrNoPermissions.WithMessage("list::account")
		}
		qops = append(qops, &application.Filter{
			AccountID: []uint64{session.Account(ctx).ID},
		})
	}
	return u.repo.FetchList(ctx, qops...)
}

// Count application objects
func (u *Usecase) Count(ctx context.Context, qops ...application.Option) (int64, error) {
	if !acl.HaveAccessCount(ctx, &models.Application{}) {
		if !acl.HaveAccessCount(ctx, &models.Application{AccountID: session.Account(ctx).ID}) {
			return 0, acl.ErrNoPermissions.WithMessage("count::account")
		}
		qops = append(qops, &application.Filter{
			AccountID: []uint64{session.Account(ctx).ID},
		})
	}
	return u.repo.Count(ctx, qops...)
}

// Create application object by id
func (u *Usecase) Create(ctx context.Context, object *models.Application) (uint64, error) {
	if object.AccountID == 0 {
		object.AccountID = session.Account(ctx).ID
	}
	if object.CreatorID == 0 {
		object.CreatorID = session.User(ctx).ID
	}
	if !acl.HaveAccessCreate(ctx, object) {
		return 0, acl.ErrNoPermissions.WithMessage("create")
	}
	object.Status = models.ApproveStatus(
		sysops.Get(`logic.crud.default.approval`, models.StatusPending).
			Int())
	return u.repo.Create(ctx, object)
}

// Update application object by id
func (u *Usecase) Update(ctx context.Context, id uint64, object *models.Application) error {
	if !acl.HaveAccessUpdate(ctx, object) {
		return acl.ErrNoPermissions.WithMessage("update")
	}
	return u.repo.Update(ctx, id, object)
}

// Delete remove object by id
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

// Run application by id
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

// Pause application
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

// Approve application by id
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

// Reject application by id
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
