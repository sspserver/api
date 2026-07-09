package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"
	gqtypes "github.com/geniusrabbit/blaze-api/server/graphql/types"
	"github.com/geniusrabbit/gosql/v2"

	"github.com/sspserver/api/pkg/repository/adformat"
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

func FilterFrom(fl *gqlmodels.AdFormatListFilter) *adformat.Filter {
	if fl == nil {
		return nil
	}
	return &adformat.Filter{
		ID:           fl.ID,
		Codename:     fl.Codename,
		CodenameLike: "",
		Type:         fl.Type,
		Active: gocast.IfThenExec(len(fl.Active) > 0,
			func() *types.ActiveStatus { return &[]types.ActiveStatus{gqlmodels.ActiveStatusFrom(fl.Active[0])}[0] },
			func() *types.ActiveStatus { return nil },
		),
	}
}

func OrderFrom(ord *gqlmodels.AdFormatListOrder) *adformat.ListOrder {
	if ord == nil {
		return nil
	}
	return &adformat.ListOrder{
		Title:     ord.Title.AsOrder(),
		Codename:  ord.Codename.AsOrder(),
		Type:      ord.Type.AsOrder(),
		Active:    ord.Active.AsOrder(),
		CreatedAt: ord.CreatedAt.AsOrder(),
		UpdatedAt: ord.UpdatedAt.AsOrder(),
	}
}

func FillModel(inp *gqlmodels.AdFormatInput, m *models.Format) {
	m.Codename = gocast.PtrAsValue(inp.Codename, m.Codename)
	m.Type = gocast.PtrAsValue(inp.Type, m.Type)

	m.Title = gocast.PtrAsValue(inp.Title, m.Title)
	m.Description = gocast.PtrAsValue(inp.Description, m.Description)

	m.Active = gocast.PtrAsValue(gqlmodels.ActiveStatusPtr(inp.Active), m.Active)

	m.Width = gocast.PtrAsValue(inp.Width, m.Width)
	m.Height = gocast.PtrAsValue(inp.Height, m.Height)
	m.MinWidth = gocast.PtrAsValue(inp.MinWidth, m.MinWidth)
	m.MinHeight = gocast.PtrAsValue(inp.MinHeight, m.MinHeight)

	m.Config = gocast.IfThenExec(inp.Config != nil,
		func() gosql.NullableJSON[types.FormatConfig] {
			return *gosql.MustNullableJSON[types.FormatConfig](inp.Config.Data)
		},
		func() gosql.NullableJSON[types.FormatConfig] { return m.Config })
}
