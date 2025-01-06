package repository

import (
	"context"
	"errors"
	"time"

	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository/historylog"
	"gorm.io/gorm"

	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/internal/repository/rtbsource"
	"github.com/sspserver/api/models"
)

type Repository struct {
	repository.Repository
}

// New create new repository
func New() *Repository {
	return &Repository{}
}

func (r *Repository) Get(ctx context.Context, id uint64) (*models.RTBSource, error) {
	object := new(models.RTBSource)
	if err := r.Slave(ctx).Find(object, id).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func (r *Repository) FetchList(ctx context.Context, qops ...rtbsource.Option) (list []*models.RTBSource, err error) {
	query := r.Slave(ctx).Model((*models.RTBSource)(nil))
	query = rtbsource.Options(qops).PrepareQuery(query)
	err = query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return list, err
}

func (r *Repository) Count(ctx context.Context, qops ...rtbsource.Option) (count int64, err error) {
	query := r.Slave(ctx).Model((*models.RTBSource)(nil))
	err = rtbsource.Options(qops).PrepareQuery(query).
		Count(&count).Error
	return count, err
}

func (r *Repository) Create(ctx context.Context, source *models.RTBSource) (uint64, error) {
	source.CreatedAt = time.Now()
	source.UpdatedAt = source.CreatedAt
	source.Status = types.StatusPending
	err := r.Master(ctx).Create(source).Error
	return source.ID, err
}

func (r *Repository) Update(ctx context.Context, id uint64, source *models.RTBSource) error {
	obj := *source
	obj.ID = id
	return r.Master(ctx).Omit(`status`, `active`, `account_id`).
		Save(&obj).Error
}

func (r *Repository) Run(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.RTBSource)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusActive).Error
}

func (r *Repository) Pause(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.RTBSource)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusPause).Error
}

func (r *Repository) Approve(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.RTBSource)(nil)).Where(`id=?`, id).
		Update(`status`, types.StatusApproved).Error
}

func (r *Repository) Reject(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.RTBSource)(nil)).Where(`id=?`, id).
		Update(`status`, types.StatusRejected).Error
}

func (r *Repository) Delete(ctx context.Context, id uint64) error {
	return r.Master(ctx).Model((*models.RTBSource)(nil)).Delete(`id=?`, id).Error
}
