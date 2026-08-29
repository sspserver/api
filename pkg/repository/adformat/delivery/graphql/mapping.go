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

		Config: fromFormatConfig(format.Config.Data),

		CreatedAt: format.CreatedAt,
		UpdatedAt: format.UpdatedAt,
		DeletedAt: gocast.IfThen(format.DeletedAt.Time.IsZero(), nil, &format.DeletedAt.Time),
	}
}

func fromFormatConfig(cfg *types.FormatConfig) *gqlmodels.AdFormatConfig {
	if cfg == nil {
		return &gqlmodels.AdFormatConfig{}
	}
	return &gqlmodels.AdFormatConfig{
		Assets: xtypes.SliceApply(cfg.Assets, fromFormatAsset),
		Fields: xtypes.SliceApply(cfg.Fields, fromFormatField),
	}
}

func fromFormatAsset(asset types.FormatFileRequirement) *gqlmodels.AdFormatAsset {
	return &gqlmodels.AdFormatAsset{
		ID:           gocast.IfThen(asset.ID != 0, gocast.Ptr(asset.ID), nil),
		Required:     gocast.IfThen(asset.Required, gocast.Ptr(asset.Required), nil),
		Name:         gocast.IfThen(asset.Name != "", gocast.Ptr(asset.Name), nil),
		AdjustSize:   gocast.IfThen(asset.AdjustSize, gocast.Ptr(asset.AdjustSize), nil),
		Width:        gocast.IfThen(asset.Width != 0, gocast.Ptr(asset.Width), nil),
		Height:       gocast.IfThen(asset.Height != 0, gocast.Ptr(asset.Height), nil),
		MinWidth:     gocast.IfThen(asset.MinWidth != 0, gocast.Ptr(asset.MinWidth), nil),
		MinHeight:    gocast.IfThen(asset.MinHeight != 0, gocast.Ptr(asset.MinHeight), nil),
		Animated:     gocast.IfThen(asset.Animated, gocast.Ptr(asset.Animated), nil),
		Sound:        gocast.IfThen(asset.Sound, gocast.Ptr(asset.Sound), nil),
		Thumbs:       asset.Thumbs,
		AllowedTypes: asset.AllowedTypes,
	}
}

func fromFormatField(field types.FormatField) *gqlmodels.AdFormatField {
	typ := string(field.Type)
	return &gqlmodels.AdFormatField{
		ID:          gocast.IfThen(field.ID != 0, gocast.Ptr(field.ID), nil),
		Required:    gocast.IfThen(field.Required, gocast.Ptr(field.Required), nil),
		Title:       gocast.IfThen(field.Title != "", gocast.Ptr(field.Title), nil),
		Description: gocast.IfThen(field.Description != "", gocast.Ptr(field.Description), nil),
		Name:        field.Name,
		Type:        gocast.IfThen(typ != "", gocast.Ptr(typ), nil),
		Exclude:     field.Exclude,
		Select: xtypes.SliceApply(field.Select, func(v any) *gqtypes.JSON {
			return gqtypes.MustJSONFrom(v)
		}),
		Min:       gocast.IfThen(field.Min != 0, gocast.Ptr(field.Min), nil),
		Max:       gocast.IfThen(field.Max != 0, gocast.Ptr(field.Max), nil),
		Mask:      gocast.IfThen(field.Mask != "", gocast.Ptr(field.Mask), nil),
		Regexp:    gocast.IfThen(field.RegExp != "", gocast.Ptr(field.RegExp), nil),
		Multiline: gocast.IfThen(field.Multiline > 0, gocast.Ptr(field.Multiline), nil),
		Editable:  field.IsEditable(),
		Multilang: field.IsMultilang(),
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
