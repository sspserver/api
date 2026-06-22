package graphql

import (
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/repository/devicemaker/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromDeviceMakerModel(m *models.DeviceMaker) *gqlmodels.DeviceMaker {
	if m == nil {
		return nil
	}
	return &gqlmodels.DeviceMaker{
		ID:          m.ID,
		Codename:    m.Codename,
		Name:        m.Name,
		Description: m.Description,
		MatchExp:    m.MatchExp,
		Models:      gqlmodels.FromDeviceModelModelList(m.Models),
		Active:      gqlmodels.FromActiveStatus(m.Active),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   gqlmodels.DeletedAt(m.DeletedAt),
	}
}

func FromDeviceMakerModelList(list []*models.DeviceMaker) []*gqlmodels.DeviceMaker {
	return xtypes.SliceApply(list, FromDeviceMakerModel)
}
