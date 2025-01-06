package repository

import (
	"context"

	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/models"
)

type StaticRepository struct {
}

func NewStaticRepository() *StaticRepository {
	return &StaticRepository{}
}

// Get return object by id
func (r *StaticRepository) Get(ctx context.Context, id uint64) (*models.DeviceType, error) {
	for _, v := range models.DeviceTypeList {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, nil
}

// List return list of objects
func (r *StaticRepository) FetchList(ctx context.Context) ([]*models.DeviceType, error) {
	return models.DeviceTypeList, nil
}

func (r *StaticRepository) FetchListByIDs(ctx context.Context, ids []uint64) ([]*models.DeviceType, error) {
	var list []*models.DeviceType
	for _, id := range xtypes.SliceUnique(ids) {
		if ir, err := r.Get(ctx, id); err == nil {
			list = append(list, ir)
		} else {
			return nil, err
		}
	}
	return list, nil
}

// Count return count of objects
func (r *StaticRepository) Count(ctx context.Context) (int64, error) {
	return int64(len(models.DeviceTypeList)), nil
}
