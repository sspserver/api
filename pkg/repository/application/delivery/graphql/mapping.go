package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/repository/application/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromApplicationModel(obj *models.Application) *gqlmodels.Application {
	return &gqlmodels.Application{
		ID: obj.ID,

		AccountID: obj.AccountID,
		CreatorID: obj.CreatorID,

		Title:       obj.Title,
		Description: obj.Description,

		URI:      obj.URI,
		Type:     gqlmodels.FromApplicationType(obj.Type),
		Platform: gqlmodels.FromPlatformType(obj.Platform),
		Premium:  obj.Premium,

		Status:  gqlmodels.FromApproveStatus(obj.Status),
		Active:  gqlmodels.FromActiveStatus(obj.Active),
		Private: gqlmodels.FromPrivateStatus(obj.Private),

		Categories:   xtypes.SliceApply(obj.Categories, func(v uint) int { return int(v) }),
		RevenueShare: gocast.IfThen(obj.RevenueShare > 0, &[]float64{obj.RevenueShare}[0], nil),

		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
	}
}

func FromApplicationModelList(list []*models.Application) []*gqlmodels.Application {
	return xtypes.SliceApply(list, FromApplicationModel)
}
