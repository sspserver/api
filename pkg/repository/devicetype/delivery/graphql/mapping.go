package graphql

import (
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/repository/devicetype/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromDeviceTypeModel(m *models.DeviceType) *gqlmodels.DeviceType {
	if m == nil {
		return nil
	}
	return &gqlmodels.DeviceType{
		ID:          m.ID,
		Name:        m.Name,
		Codename:    m.Codename,
		Description: m.Description,
		Active:      gqlmodels.FromActiveStatus(m.Active),
	}
}

func FromDeviceTypeModelList(m []*models.DeviceType) []*gqlmodels.DeviceType {
	return xtypes.SliceApply(m, FromDeviceTypeModel)
}
