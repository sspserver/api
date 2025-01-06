package repository

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/sspserver/api/internal/repository"
	"github.com/sspserver/api/internal/repository/devicemodel"
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

func (r *Repository) Get(ctx context.Context, id uint64) (*models.DeviceModel, error) {
	object := new(models.DeviceModel)
	if err := r.Slave(ctx).Preload(`Maker`).Find(object, id).Error; err != nil {
		return nil, err
	}
	if object.TypeID > 0 {
		tobj, err := r.typesRepo.Get(ctx, object.TypeID)
		if err != nil {
			return nil, err
		}
		object.Type = tobj
	}
	return object, nil
}

func (r *Repository) FetchList(ctx context.Context, qops ...devicemodel.Option) (list []*models.DeviceModel, err error) {
	query := r.Slave(ctx).Preload(`Maker`).Model((*models.DeviceModel)(nil))
	query = devicemodel.Options(qops).PrepareQuery(query)
	err = query.Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	for _, obj := range list {
		if obj.TypeID == 0 {
			continue
		}
		tobj, err := r.typesRepo.Get(ctx, obj.TypeID)
		if err != nil {
			return list, err
		}
		obj.Type = tobj
	}
	return list, err
}

func (r *Repository) Count(ctx context.Context, qops ...devicemodel.Option) (count int64, err error) {
	query := r.Slave(ctx).Model((*models.DeviceModel)(nil))
	err = devicemodel.Options(qops).PrepareQuery(query).
		Count(&count).Error
	return count, err
}

func (r *Repository) Create(ctx context.Context, object *models.DeviceModel) (uint64, error) {
	object.CreatedAt = time.Now()
	object.UpdatedAt = object.CreatedAt
	err := r.Master(ctx).Create(object).Error
	return object.ID, err
}

func (r *Repository) Update(ctx context.Context, id uint64, object *models.DeviceModel) error {
	obj := *object
	obj.ID = id
	return r.Master(ctx).Updates(&obj).Error
}

func (r *Repository) Delete(ctx context.Context, id uint64) error {
	return r.Master(ctx).Model((*models.DeviceModel)(nil)).Delete(`id=?`, id).Error
}
