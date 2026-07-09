package usecase

import (
	"context"

	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/pkg/context/session"

	"github.com/sspserver/api/pkg/acl"
	"github.com/sspserver/api/pkg/repository/trafficrouter"
	"github.com/sspserver/api/pkg/repository/trafficrouter/models"
	"github.com/sspserver/api/pkg/sysops"
)

type Usecase struct {
	repo trafficrouter.Repository
}

// New creates a new instance of Usecase.
func New(repo trafficrouter.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

// Get returns the object by id
func (uc *Usecase) Get(ctx context.Context, id uint64) (*models.TrafficRouter, error) {
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
func (uc *Usecase) FetchList(ctx context.Context, qops ...trafficrouter.Option) ([]*models.TrafficRouter, error) {
	if !acl.HaveAccessList(ctx, &models.TrafficRouter{}) {
		accountID := session.AccountID(ctx)
		if acl.HaveAccessList(ctx, &models.TrafficRouter{AccountID: accountID}) {
			qops = trafficrouter.Options(qops).With(&trafficrouter.Filter{AccountID: accountID})
		} else {
			return nil, acl.ErrNoPermissions.WithMessage("fetch list")
		}
	}
	return uc.repo.FetchList(ctx, qops...)
}

// Count returns the total number of TrafficRouter objects.
func (uc *Usecase) Count(ctx context.Context, qops ...trafficrouter.Option) (int64, error) {
	if !acl.HaveAccessCount(ctx, &models.TrafficRouter{}) {
		accountID := session.AccountID(ctx)
		if acl.HaveAccessCount(ctx, &models.TrafficRouter{AccountID: accountID}) {
			qops = trafficrouter.Options(qops).With(&trafficrouter.Filter{AccountID: accountID})
		} else {
			return 0, acl.ErrNoPermissions.WithMessage("count")
		}
	}
	return uc.repo.Count(ctx, qops...)
}

// Create creates a new TrafficRouter object.
func (uc *Usecase) Create(ctx context.Context, router *models.TrafficRouter) (uint64, error) {
	if router.AccountID == 0 {
		router.AccountID = session.AccountID(ctx)
	}
	if !acl.HaveAccessCreate(ctx, router) {
		return 0, acl.ErrNoPermissions.WithMessage("create")
	}
	router.Status = types.ApproveStatus(
		sysops.Get(`logic.crud.default.approval`, int(types.StatusPending)).
			Int())
	return uc.repo.Create(ctx, router)
}

// Update updates an existing TrafficRouter object.
func (uc *Usecase) Update(ctx context.Context, id uint64, router *models.TrafficRouter) error {
	routerObj := *router
	routerObj.ID = id
	if !acl.HaveAccessUpdate(ctx, routerObj) {
		return acl.ErrNoPermissions.WithMessage("update")
	}
	return uc.repo.Update(ctx, id, router)
}

// Delete removes the object by id
func (uc *Usecase) Delete(ctx context.Context, id uint64) error {
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
func (uc *Usecase) Run(ctx context.Context, id uint64, message string) error {
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
func (uc *Usecase) Pause(ctx context.Context, id uint64, message string) error {
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
func (uc *Usecase) Approve(ctx context.Context, id uint64, message string) error {
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
func (uc *Usecase) Reject(ctx context.Context, id uint64, message string) error {
	router, err := uc.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	if !acl.HaveAccessReject(ctx, router) {
		return acl.ErrNoPermissions.WithMessage("reject")
	}
	return uc.repo.Reject(ctx, id, message)
}
