package repository

import (
	"context"
	"errors"

	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository/historylog"
	"gorm.io/gorm"

	"github.com/sspserver/api/pkg/models"
	pkgrepo "github.com/sspserver/api/pkg/repository"
	"github.com/sspserver/api/pkg/repository/trafficrouter"
)

var (
	ErrInvalidTrafficRouter = errors.New(`invalid traffic router`)
)

type Repository struct {
	pkgrepo.Repository
}

// New creates a new instance of Repository.
func New() *Repository {
	return &Repository{}
}

// Get returns the object by id
func (rep *Repository) Get(ctx context.Context, id uint64) (*models.TrafficRouter, error) {
	obj := &models.TrafficRouter{}
	if err := rep.Slave(ctx).First(obj, id).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

// FetchList retrieves a list of TrafficRouter objects.
func (rep *Repository) FetchList(ctx context.Context, qops ...trafficrouter.Option) ([]*models.TrafficRouter, error) {
	var list []*models.TrafficRouter
	query := rep.Slave(ctx).Model((*models.TrafficRouter)(nil))
	query = trafficrouter.Options(qops).PrepareQuery(query)
	err := query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return list, err
}

// Count returns the total number of TrafficRouter objects.
func (rep *Repository) Count(ctx context.Context, qops ...trafficrouter.Option) (int64, error) {
	var count int64
	query := rep.Slave(ctx).Model((*models.TrafficRouter)(nil))
	query = trafficrouter.Options(qops).PrepareQuery(query)
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Create creates a new TrafficRouter object.
func (rep *Repository) Create(ctx context.Context, router *models.TrafficRouter) (uint64, error) {
	if router == nil {
		return 0, ErrInvalidTrafficRouter
	}
	err := rep.Master(ctx).Create(router).Error
	if err != nil {
		return 0, err
	}
	return router.ID, nil
}

// Update updates an existing TrafficRouter object.
func (rep *Repository) Update(ctx context.Context, id uint64, router *models.TrafficRouter) error {
	if router == nil {
		return ErrInvalidTrafficRouter
	}
	obj := *router
	obj.ID = id
	return rep.Master(ctx).
		Omit(`active`, `status`, `account_id`).
		Save(&obj).Error
}

// Delete deletes a TrafficRouter object by ID.
func (rep *Repository) Delete(ctx context.Context, id uint64) error {
	return rep.Master(
		historylog.WithPK(ctx, id),
	).Model((*models.TrafficRouter)(nil)).Delete(`id=?`, id).Error
}

// Run sets the TrafficRouter object to active status.
func (rep *Repository) Run(ctx context.Context, id uint64, message string) error {
	return rep.Master(
		historylog.WithMessageAndPK(ctx, message, id),
	).Model((*models.TrafficRouter)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusActive).Error
}

// Pause sets the TrafficRouter object to pause status.
func (rep *Repository) Pause(ctx context.Context, id uint64, message string) error {
	return rep.Master(
		historylog.WithMessageAndPK(ctx, message, id),
	).Model((*models.TrafficRouter)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusPause).Error
}

// Approve sets the TrafficRouter object to approved status.
func (rep *Repository) Approve(ctx context.Context, id uint64, message string) error {
	return rep.Master(
		historylog.WithMessageAndPK(ctx, message, id),
	).Model((*models.TrafficRouter)(nil)).
		Where(`id=?`, id).Update(`status`, types.StatusApproved).Error
}

// Reject sets the TrafficRouter object to rejected status.
func (rep *Repository) Reject(ctx context.Context, id uint64, message string) error {
	return rep.Master(
		historylog.WithMessageAndPK(ctx, message, id),
	).Model((*models.TrafficRouter)(nil)).
		Where(`id=?`, id).Update(`status`, types.StatusRejected).Error
}
