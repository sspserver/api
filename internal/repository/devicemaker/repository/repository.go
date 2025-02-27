package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/internal/repository/devicemaker"
	"github.com/sspserver/api/internal/repository/devicetype"
	devicetyperepo "github.com/sspserver/api/internal/repository/devicetype/repository"
	"github.com/sspserver/api/models"
)

type Repository struct {
	repository.Repository
	typesRepo devicetype.Repository
}

// New create new repository
func New() *Repository {
	return &Repository{
		typesRepo: devicetyperepo.NewStaticRepository(),
	}
}

func (r *Repository) Get(ctx context.Context, id uint64, preloads ...string) (*models.DeviceMaker, error) {
	object := new(models.DeviceMaker)
	query := r.Slave(ctx)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	if err := query.Find(object, id).Error; err != nil {
		return nil, err
	}
	for _, model := range object.Models {
		if model.TypeCodename == "" {
			continue
		}
		tobj, err := r.typesRepo.GetByCodename(ctx, model.TypeCodename)
		if err != nil {
			return nil, err
		}
		model.Type = tobj
	}
	return object, nil
}

func (r *Repository) GetByCodename(ctx context.Context, codename string, preloads ...string) (*models.DeviceMaker, error) {
	object := new(models.DeviceMaker)
	query := r.Slave(ctx)
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	if err := query.Where(`codename=?`, codename).Find(object).Error; err != nil {
		return nil, err
	}
	for _, model := range object.Models {
		if model.TypeCodename == "" {
			continue
		}
		tobj, err := r.typesRepo.GetByCodename(ctx, model.TypeCodename)
		if err != nil {
			return nil, err
		}
		model.Type = tobj
	}
	return object, nil
}

func (r *Repository) FetchList(ctx context.Context, qops ...devicemaker.Option) (list []*models.DeviceMaker, err error) {
	query := r.Slave(ctx).Model((*models.DeviceMaker)(nil))
	query = devicemaker.Options(qops).PrepareQuery(query)
	err = query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	for _, obj := range list {
		for _, model := range obj.Models {
			if model.TypeCodename == "" {
				continue
			}
			tobj, err := r.typesRepo.GetByCodename(ctx, model.TypeCodename)
			if err != nil {
				return list, err
			}
			model.Type = tobj
		}
	}
	return list, err
}

func (r *Repository) Count(ctx context.Context, qops ...devicemaker.Option) (count int64, err error) {
	query := r.Slave(ctx).Model((*models.DeviceMaker)(nil))
	err = devicemaker.Options(qops).PrepareQuery(query).
		Count(&count).Error
	return count, err
}

func (r *Repository) Create(ctx context.Context, object *models.DeviceMaker) (uint64, error) {
	object.CreatedAt = time.Now()
	object.UpdatedAt = object.CreatedAt
	err := r.Master(ctx).Create(object).Error
	return object.ID, err
}

func (r *Repository) Update(ctx context.Context, id uint64, object *models.DeviceMaker) error {
	obj := *object
	obj.ID = id
	return r.Master(ctx).Updates(&obj).Error
}

func (r *Repository) Delete(ctx context.Context, id uint64) error {
	return r.Master(ctx).Model((*models.DeviceMaker)(nil)).Delete(`id=?`, id).Error
}
