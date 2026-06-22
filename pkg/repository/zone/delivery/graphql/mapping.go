package graphql

import (
	"github.com/demdxx/xtypes"
	gqtypes "github.com/geniusrabbit/blaze-api/server/graphql/types"

	"github.com/sspserver/api/pkg/repository/zone/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromZoneModel(obj *models.Zone) *gqlmodels.Zone {
	if obj == nil {
		return nil
	}
	return &gqlmodels.Zone{
		ID:        obj.ID,
		Codename:  obj.Codename,
		AccountID: obj.AccountID,

		Title:       obj.Title,
		Description: obj.Description,

		Status: gqlmodels.FromApproveStatus(obj.Status),
		Active: gqlmodels.FromActiveStatus(obj.Active),

		DefaultCode: *gqtypes.MustNullableJSONFrom(obj.DefaultCode.DataOr(nil)),
		Context:     *gqtypes.MustNullableJSONFrom(obj.Context.DataOr(nil)),

		MinEcpm:            obj.MinECPM,
		FixedPurchasePrice: obj.FixedPurchasePrice,

		AllowedFormats:    obj.AllowedFormats,
		AllowedTypes:      obj.AllowedTypes,
		AllowedSources:    obj.AllowedSources,
		DisallowedSources: obj.DisallowedSources,
		Campaigns:         obj.Campaigns,

		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
	}
}

func FromZoneModelList(obj []*models.Zone) []*gqlmodels.Zone {
	return xtypes.SliceApply(obj, FromZoneModel)
}
