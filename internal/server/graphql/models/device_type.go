package models

import (
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/models"
)

func FromDeviceTypeModel(m *models.DeviceType) *DeviceType {
	if m == nil {
		return nil
	}
	return &DeviceType{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Active:      FromActiveStatus(m.Active),
	}
}

func FromDeviceTypeModelList(m []*models.DeviceType) []*DeviceType {
	return xtypes.SliceApply(m, FromDeviceTypeModel)
}
