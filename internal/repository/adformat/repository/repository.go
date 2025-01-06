package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/internal/repository/adformat"
	"github.com/sspserver/api/models"
)

type Repository struct {
	repository.Repository
}

// New create new repository
func New() *Repository {
	return &Repository{}
}

func (r *Repository) Get(ctx context.Context, id uint64) (*models.Format, error) {
	object := new(models.Format)
	if err := r.Slave(ctx).Find(object, id).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func (r *Repository) GetByCodename(ctx context.Context, codename string) (*models.Format, error) {
	object := new(models.Format)
	if err := r.Slave(ctx).Where(`codename=?`, codename).First(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func (r *Repository) FetchList(ctx context.Context, qops ...adformat.Option) (list []*models.Format, err error) {
	query := r.Slave(ctx).Model((*models.Format)(nil))
	query = adformat.Options(qops).PrepareQuery(query)
	err = query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return list, err
}

func (r *Repository) Count(ctx context.Context, qops ...adformat.Option) (count int64, err error) {
	query := r.Slave(ctx).Model((*models.Format)(nil))
	err = adformat.Options(qops).PrepareQuery(query).
		Count(&count).Error
	return count, err
}

func (r *Repository) Create(ctx context.Context, format *models.Format) (uint64, error) {
	format.CreatedAt = time.Now()
	format.UpdatedAt = format.CreatedAt
	err := r.Master(ctx).Create(format).Error
	return format.ID, err
}

func (r *Repository) Update(ctx context.Context, id uint64, format *models.Format) error {
	obj := *format
	obj.ID = id
	return r.Master(ctx).Save(&obj).Error
}

func (r *Repository) Delete(ctx context.Context, id uint64) error {
	return r.Master(ctx).Model((*models.Format)(nil)).Delete(`id=?`, id).Error
}

func (r *Repository) DeleteByCodename(ctx context.Context, codename string) error {
	return r.Master(ctx).Model((*models.Format)(nil)).Delete(`codename=?`, codename).Error
}
