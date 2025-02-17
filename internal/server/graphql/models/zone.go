package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"
	gqtypes "github.com/geniusrabbit/blaze-api/server/graphql/types"

	"github.com/sspserver/api/internal/repository/zone"
	"github.com/sspserver/api/models"
)

func FromZoneModel(obj *models.Zone) *Zone {
	if obj == nil {
		return nil
	}
	return &Zone{
		ID:        obj.ID,
		Codename:  obj.Codename,
		AccountID: obj.AccountID,

		Title:       obj.Title,
		Description: obj.Description,

		Status: FromApproveStatus(obj.Status),
		Active: FromActiveStatus(obj.Active),

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

func FromZoneModelList(obj []*models.Zone) []*Zone {
	return xtypes.SliceApply(obj, FromZoneModel)
}

func (fl *ZoneListFilter) Filter() *zone.Filter {
	if fl == nil {
		return nil
	}
	return &zone.Filter{
		ID:        fl.ID,
		Codename:  fl.Codename,
		AccountID: fl.AccountID,
		Status: gocast.IfThenExec(fl.Status != nil,
			func() *types.ApproveStatus {
				st := types.ApproveStatus(fl.Status.ModelStatus())
				return &st
			},
			func() *types.ApproveStatus { return nil }),
		Active:  ActiveStatusPtr(fl.Active),
		MinECPM: fl.MinEcpm,
		MaxECPM: fl.MaxEcpm,
	}
}

func (ord *ZoneListOrder) Order() *zone.ListOrder {
	if ord == nil {
		return nil
	}
	return &zone.ListOrder{
		ID:       ord.ID.AsOrder(),
		Codename: ord.Codename.AsOrder(),

		Title:     ord.Title.AsOrder(),
		AccountID: ord.AccountID.AsOrder(),

		Type:   ord.Type.AsOrder(),
		Status: ord.Status.AsOrder(),
		Active: ord.Active.AsOrder(),

		MinECPM: ord.MinEcpm.AsOrder(),

		CreatedAt: ord.CreatedAt.AsOrder(),
		UpdatedAt: ord.UpdatedAt.AsOrder(),
	}
}

func (inp *ZoneInput) FillModel(obj *models.Zone) {
	if obj == nil {
		return
	}
	obj.Codename = gocast.PtrAsValue(inp.Codename, obj.Codename)
	obj.AccountID = gocast.PtrAsValue(inp.AccountID, obj.AccountID)

	obj.Title = gocast.PtrAsValue(inp.Title, obj.Title)
	obj.Description = gocast.PtrAsValue(inp.Description, obj.Description)

	if inp.DefaultCode != nil && inp.DefaultCode.Data != nil {
		_ = obj.DefaultCode.SetValue(inp.DefaultCode.Data)
	}
	if inp.Context != nil && inp.Context.Data != nil {
		_ = obj.Context.SetValue(inp.Context.Data)
	}

	obj.MinECPM = gocast.PtrAsValue(inp.MinEcpm, obj.MinECPM)
	obj.FixedPurchasePrice = gocast.PtrAsValue(inp.FixedPurchasePrice, obj.FixedPurchasePrice)

	obj.AllowedFormats = gocast.IfThen(inp.AllowedFormats != nil, inp.AllowedFormats, obj.AllowedFormats)
	obj.AllowedTypes = gocast.IfThen(inp.AllowedTypes != nil, inp.AllowedTypes, obj.AllowedTypes)
	obj.AllowedSources = gocast.IfThen(inp.AllowedSources != nil, inp.AllowedSources, obj.AllowedSources)
	obj.DisallowedSources = gocast.IfThen(inp.DisallowedSources != nil, inp.DisallowedSources, obj.DisallowedSources)
	obj.Campaigns = gocast.IfThen(inp.Campaigns != nil, inp.Campaigns, obj.Campaigns)
}
