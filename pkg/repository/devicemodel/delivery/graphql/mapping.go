package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"

	gqldevicemaker "github.com/sspserver/api/pkg/repository/devicemaker/delivery/graphql"
	"github.com/sspserver/api/pkg/repository/devicemodel/models"
	gqldevicetype "github.com/sspserver/api/pkg/repository/devicetype/delivery/graphql"
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
		Maker:         gqldevicemaker.FromDeviceMakerModel(m.Maker),
		TypeCodename:  m.TypeCodename,
		Type:          gqldevicetype.FromDeviceTypeModel(m.Type),
		Versions:      xtypes.SliceApply(m.Versions, FromDeviceModelModel),
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		DeletedAt:     gqlmodels.DeletedAt(m.DeletedAt),
	}
}

func FromDeviceModelModelList(list []*models.DeviceModel) []*gqlmodels.DeviceModel {
	return xtypes.SliceApply(list, FromDeviceModelModel)
}
