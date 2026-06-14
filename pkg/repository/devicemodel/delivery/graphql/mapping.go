package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromDeviceModelModel(m *models.DeviceModel) *gqlmodels.DeviceModel {
	if m == nil {
		return nil
	}
	return &gqlmodels.DeviceModel{
		ID:            m.ID,
		Name:          m.Name,
		Codename:      m.Codename,
		Description:   m.Description,
		Active:        gqlmodels.FromActiveStatus(m.Active),
		YearRelease:   m.YearRelease,
		MatchExp:      m.MatchExp,
		ParentID:      gocast.IfThen(m.ParentID > 0, &m.ParentID, nil),
		MakerCodename: m.MakerCodename,
		Maker:         gqlmodels.FromDeviceMakerModel(m.Maker),
		TypeCodename:  m.TypeCodename,
		Type:          gqlmodels.FromDeviceTypeModel(m.Type),
		Versions:      xtypes.SliceApply(m.Versions, FromDeviceModelModel),
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		DeletedAt:     gqlmodels.DeletedAt(m.DeletedAt),
	}
}

func FromDeviceModelModelList(list []*models.DeviceModel) []*gqlmodels.DeviceModel {
	return xtypes.SliceApply(list, FromDeviceModelModel)
}
