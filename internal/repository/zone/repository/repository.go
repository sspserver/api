package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/blaze-api/repository/historylog"

	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/internal/repository/zone"
	"github.com/sspserver/api/models"
)

type Repository struct {
	repository.Repository
}

// New create new repository
func New() *Repository {
	return &Repository{}
}

func (r *Repository) Get(ctx context.Context, id uint64) (*models.Zone, error) {
	object := new(models.Zone)
	if err := r.Slave(ctx).Find(object, id).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func (r *Repository) GetByCodename(ctx context.Context, codename string) (*models.Zone, error) {
	object := new(models.Zone)
	if err := r.Slave(ctx).Where(`codename=?`, codename).First(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func (r *Repository) FetchList(ctx context.Context, qops ...zone.Option) (list []*models.Zone, err error) {
	query := r.Slave(ctx).Model((*models.Zone)(nil))
	query = zone.Options(qops).PrepareQuery(query)
	err = query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return list, err
}

func (r *Repository) Count(ctx context.Context, qops ...zone.Option) (count int64, err error) {
	query := r.Slave(ctx).Model((*models.Zone)(nil))
	err = zone.Options(qops).PrepareQuery(query).
		Count(&count).Error
	return count, err
}

func (r *Repository) Create(ctx context.Context, object *models.Zone) (uint64, error) {
	object.CreatedAt = time.Now()
	object.UpdatedAt = object.CreatedAt
	err := r.Master(ctx).Create(object).Error
	return object.ID, err
}

func (r *Repository) Update(ctx context.Context, id uint64, object *models.Zone) error {
	obj := *object
	obj.ID = id
	return r.Master(ctx).Save(&obj).Error
}

func (r *Repository) Delete(ctx context.Context, id uint64, message string) error {
	return r.Master(historylog.WithMessage(ctx, message)).
		Model((*models.Zone)(nil)).Delete(`id=?`, id).Error
}

func (r *Repository) Run(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.Zone)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusActive).Error
}

func (r *Repository) Pause(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.Zone)(nil)).
		Where(`id=?`, id).Update(`active`, types.StatusPause).Error
}

func (r *Repository) Approve(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.Zone)(nil)).
		Where(`id=?`, id).Update(`status`, types.StatusApproved).Error
}

func (r *Repository) Reject(ctx context.Context, id uint64, message string) error {
	return r.Master(
		historylog.WithMessage(ctx, message),
	).Model((*models.Zone)(nil)).
		Where(`id=?`, id).Update(`status`, types.StatusRejected).Error
}
