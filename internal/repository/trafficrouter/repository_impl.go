package trafficrouter

import (
	"context"
	"errors"

	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository/historylog"
	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/models"
	"gorm.io/gorm"
)

var (
	ErrInvalidTrafficRouter = errors.New(`invalid traffic router`)
)

type RepositoryImpl struct {
	repository.Repository
}

// NewRepositoryImpl creates a new instance of RepositoryImpl.
func NewRepositoryImpl() *RepositoryImpl {
	return &RepositoryImpl{}
}

// Get returns the object by id
func (rep *RepositoryImpl) Get(ctx context.Context, id uint64) (*models.TrafficRouter, error) {
	obj := &models.TrafficRouter{}
	if err := rep.Slave(ctx).First(obj, id).Error; err != nil {
		return nil, err
	}
	return obj, nil
}

// FetchList retrieves a list of TrafficRouter objects.
func (rep *RepositoryImpl) FetchList(ctx context.Context, qops ...Option) ([]*models.TrafficRouter, error) {
	var list []*models.TrafficRouter
	query := rep.Slave(ctx).Model((*models.TrafficRouter)(nil))
	query = Options(qops).PrepareQuery(query)
	err := query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return list, err
}

// Count returns the total number of TrafficRouter objects.
func (rep *RepositoryImpl) Count(ctx context.Context, qops ...Option) (int64, error) {
	var count int64
	query := rep.Slave(ctx).Model((*models.TrafficRouter)(nil))
	query = Options(qops).PrepareQuery(query)
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Create creates a new TrafficRouter object.
func (rep *RepositoryImpl) Create(ctx context.Context, router *models.TrafficRouter) (uint64, error) {
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
func (rep *RepositoryImpl) Update(ctx context.Context, id uint64, router *models.TrafficRouter) error {
	if router == nil {
		return ErrInvalidTrafficRouter
	}
	obj := *router
	obj.ID = id
	return rep.Master(ctx).Omit(`active`, `account_id`).
		Save(&obj).Error
}

// Delete deletes a TrafficRouter object by ID.
func (rep *RepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return rep.Master(ctx).Model((*models.TrafficRouter)(nil)).Delete(`id=?`, id).Error
}

// Run sets the TrafficRouter object to active status.
func (rep *RepositoryImpl) Run(ctx context.Context, id uint64, message string) error {
	return rep.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.TrafficRouter)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusActive).Error
}

// Pause sets the TrafficRouter object to pause status.
func (rep *RepositoryImpl) Pause(ctx context.Context, id uint64, message string) error {
	return rep.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.TrafficRouter)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusPause).Error
}
