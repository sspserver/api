package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/internal/repository/category"
	"github.com/sspserver/api/models"
)

type Repository struct {
	repository.Repository
}

// New create new repository
func New() *Repository {
	return &Repository{}
}

func (r *Repository) Get(ctx context.Context, id uint64) (*models.Category, error) {
	object := new(models.Category)
	if err := r.Slave(ctx).Preload(`Parent`).Find(object, id).Error; err != nil {
		return nil, err
	}
	return object, nil
}

func (r *Repository) FetchList(ctx context.Context, qops ...category.Option) (list []*models.Category, err error) {
	query := r.Slave(ctx).Model((*models.Category)(nil))
	query = category.Options(qops).PrepareQuery(query)
	err = query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return list, err
}

func (r *Repository) Count(ctx context.Context, qops ...category.Option) (count int64, err error) {
	query := r.Slave(ctx).Model((*models.Category)(nil))
	err = category.Options(qops).PrepareQuery(query).
		Count(&count).Error
	return count, err
}

func (r *Repository) Create(ctx context.Context, object *models.Category) (uint64, error) {
	object.CreatedAt = time.Now()
	object.UpdatedAt = object.CreatedAt
	err := r.Master(ctx).Create(object).Error
	return object.ID, err
}

func (r *Repository) Update(ctx context.Context, id uint64, object *models.Category) error {
	obj := *object
	obj.ID = id
	return r.Master(ctx).Updates(&obj).Error
}

func (r *Repository) Delete(ctx context.Context, id uint64) error {
	return r.Master(ctx).Model((*models.Category)(nil)).Delete(`id=?`, id).Error
}
