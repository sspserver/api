package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	gqtypes "github.com/geniusrabbit/blaze-api/server/graphql/types"

	"github.com/sspserver/api/pkg/repository/adformat/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromAdFormatModel(format *models.Format) *gqlmodels.AdFormat {
	if format == nil {
		return nil
	}
	return &gqlmodels.AdFormat{
		ID:       format.ID,
		Codename: format.Codename,
		Type:     format.Type,

		Title:       format.Title,
		Description: format.Description,

		Active: gqlmodels.FromActiveStatus(format.Active),

		Width:     format.Width,
		Height:    format.Height,
		MinWidth:  format.MinWidth,
		MinHeight: format.MinHeight,

		Config: *gqtypes.MustNullableJSONFrom(format.Config.Data),

		CreatedAt: format.CreatedAt,
		UpdatedAt: format.UpdatedAt,
		DeletedAt: gocast.IfThen(format.DeletedAt.Time.IsZero(), nil, &format.DeletedAt.Time),
	}
}

func FromAdFormatModelList(format []*models.Format) []*gqlmodels.AdFormat {
	return xtypes.SliceApply(format, FromAdFormatModel)
}
