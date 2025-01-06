package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository/historylog"

	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/internal/repository/application"
	"github.com/sspserver/api/models"
)

type Repository struct {
	repository.Repository
}

// New create new repository
func New() *Repository {
	return &Repository{}
}

func (r *Repository) Get(ctx context.Context, id uint64) (*models.Application, error) {
	object := new(models.Application)
	if err := r.Slave(ctx).Find(object, id).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func (r *Repository) FetchList(ctx context.Context, qops ...application.Option) (list []*models.Application, err error) {
	query := r.Slave(ctx).Model((*models.Application)(nil))
	query = application.Options(qops).PrepareQuery(query)
	err = query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return list, err
}

func (r *Repository) Count(ctx context.Context, qops ...application.Option) (count int64, err error) {
	query := r.Slave(ctx).Model((*models.Application)(nil))
	err = application.Options(qops).PrepareQuery(query).
		Count(&count).Error
	return count, err
}

func (r *Repository) Create(ctx context.Context, object *models.Application) (uint64, error) {
	object.CreatedAt = time.Now()
	object.UpdatedAt = object.CreatedAt
	sqlStr := r.Master(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB { return tx.Create(object) })
	err := r.Master(ctx).Raw(sqlStr).Scan(&object.ID).Error
	return object.ID, err
}

func (r *Repository) Update(ctx context.Context, id uint64, object *models.Application) error {
	obj := *object
	sqlStr := r.Master(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Where(`id=?`, id).Save(&obj)
	})
	return r.Master(ctx).Exec(sqlStr).Error
}

func (r *Repository) Delete(ctx context.Context, id uint64, message string) error {
	return r.Master(historylog.WithMessage(ctx, message)).
		Model((*models.Application)(nil)).Delete(`id=?`, id).Error
}

func (r *Repository) Run(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.Application)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusActive).Error
}

func (r *Repository) Pause(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.Application)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusPause).Error
}

func (r *Repository) Approve(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.Application)(nil)).
		Where(`id=?`, id).Update(`status`, types.StatusApproved).Error
}

func (r *Repository) Reject(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.Application)(nil)).
		Where(`id=?`, id).Update(`status`, types.StatusRejected).Error
}
