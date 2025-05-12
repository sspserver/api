package trafficrouter

import (
	"context"

	"github.com/geniusrabbit/blaze-api/pkg/context/session"
	"github.com/sspserver/api/pkg/acl"
	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/sysops"
)

type UsecaseImpl struct {
	repo Repository
}

// NewUsecaseImpl creates a new instance of UsecaseImpl.
func NewUsecaseImpl(repo Repository) *UsecaseImpl {
	return &UsecaseImpl{
		repo: repo,
	}
}

// Get returns the object by id
func (uc *UsecaseImpl) Get(ctx context.Context, id uint64) (*models.TrafficRouter, error) {
	src, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if !acl.HaveAccessView(ctx, src) {
		return nil, acl.ErrNoPermissions.WithMessage("view")
	}
	return src, nil
}

// FetchList retrieves a list of TrafficRouter objects.
func (uc *UsecaseImpl) FetchList(ctx context.Context, qops ...Option) ([]*models.TrafficRouter, error) {
	if !acl.HaveAccessList(ctx, &models.TrafficRouter{}) {
		accountID := session.Account(ctx).ID
		if acl.HaveAccessList(ctx, &models.TrafficRouter{AccountID: accountID}) {
			qops = Options(qops).With(&Filter{AccountID: accountID})
		} else {
			return nil, acl.ErrNoPermissions.WithMessage("fetch list")
		}
	}
	return uc.repo.FetchList(ctx, qops...)
}

// Count returns the total number of TrafficRouter objects.
func (uc *UsecaseImpl) Count(ctx context.Context, qops ...Option) (int64, error) {
	if !acl.HaveAccessCount(ctx, &models.TrafficRouter{}) {
		accountID := session.Account(ctx).ID
		if acl.HaveAccessCount(ctx, &models.TrafficRouter{AccountID: accountID}) {
			qops = Options(qops).With(&Filter{AccountID: accountID})
		} else {
			return 0, acl.ErrNoPermissions.WithMessage("count")
		}
	}
	return uc.repo.Count(ctx, qops...)
}

// Create creates a new TrafficRouter object.
func (uc *UsecaseImpl) Create(ctx context.Context, router *models.TrafficRouter) (uint64, error) {
	if router.AccountID == 0 {
		router.AccountID = session.Account(ctx).ID
	}
	if !acl.HaveAccessCreate(ctx, router) {
		return 0, acl.ErrNoPermissions.WithMessage("create")
	}
	router.Status = models.ApproveStatus(
		sysops.Get(`logic.crud.default.approval`, models.StatusPending).
			Int())
	return uc.repo.Create(ctx, router)
}

// Update updates an existing TrafficRouter object.
func (uc *UsecaseImpl) Update(ctx context.Context, id uint64, router *models.TrafficRouter) error {
	routerObj := *router
	routerObj.ID = id
	if !acl.HaveAccessUpdate(ctx, routerObj) {
		return acl.ErrNoPermissions.WithMessage("update")
	}
	return uc.repo.Update(ctx, id, router)
}

// Delete removes the object by id
func (uc *UsecaseImpl) Delete(ctx context.Context, id uint64) error {
	router, err := uc.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessDelete(ctx, router) {
		return acl.ErrNoPermissions.WithMessage("delete")
	}
	return uc.repo.Delete(ctx, id)
}

// Run starts the router by id
func (uc *UsecaseImpl) Run(ctx context.Context, id uint64, message string) error {
	router, err := uc.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessRun(ctx, router) {
		return acl.ErrNoPermissions.WithMessage("run")
	}
	return uc.repo.Run(ctx, id, message)
}

// Pause stops the router by id
func (uc *UsecaseImpl) Pause(ctx context.Context, id uint64, message string) error {
	router, err := uc.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessPause(ctx, router) {
		return acl.ErrNoPermissions.WithMessage("pause")
	}
	return uc.repo.Pause(ctx, id, message)
}

// Approve approves the router by id
func (uc *UsecaseImpl) Approve(ctx context.Context, id uint64, message string) error {
	router, err := uc.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessApprove(ctx, router) {
		return acl.ErrNoPermissions.WithMessage("approve")
	}
	return uc.repo.Approve(ctx, id, message)
}

// Reject rejects the router by id
func (uc *UsecaseImpl) Reject(ctx context.Context, id uint64, message string) error {
	router, err := uc.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessReject(ctx, router) {
		return acl.ErrNoPermissions.WithMessage("reject")
	}
	return uc.repo.Reject(ctx, id, message)
}
