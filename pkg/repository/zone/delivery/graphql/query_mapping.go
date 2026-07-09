package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/adcorelib/admodels/types"
	bzgqlmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/pkg/repository/zone"
	"github.com/sspserver/api/pkg/repository/zone/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func ZoneFilterFromGraphQL(fl *gqlmodels.ZoneListFilter) *zone.Filter {
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
		Active:  activeStatusPtrFromBlazeGraphQL(fl.Active),
		MinECPM: fl.MinEcpm,
		MaxECPM: fl.MaxEcpm,
	}
}

func ZoneOrderFromGraphQL(ord *gqlmodels.ZoneListOrder) *zone.ListOrder {
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

func FillZoneCreateInputModel(inp gqlmodels.ZoneCreateInput, obj *models.Zone) {
	if obj == nil {
		return
	}
	obj.AccountID = gocast.PtrAsValue(inp.AccountID, obj.AccountID)

	obj.Title = inp.Title
	obj.Description = gocast.PtrAsValue(inp.Description, obj.Description)

	if inp.DefaultCode != nil && inp.DefaultCode.Data != nil {
		_ = obj.DefaultCode.SetValue(inp.DefaultCode.Data)
	}
	if inp.Context != nil && inp.Context.Data != nil {
		_ = obj.Context.SetValue(inp.Context.Data)
	}

	obj.MinECPM = gocast.PtrAsValue(inp.MinEcpm, obj.MinECPM)
	obj.FixedPurchasePrice = gocast.PtrAsValue(inp.FixedPurchasePrice, obj.FixedPurchasePrice)

	obj.AllowedFormats = inp.AllowedFormats
	obj.AllowedTypes = inp.AllowedTypes
	obj.AllowedSources = inp.AllowedSources
	obj.DisallowedSources = inp.DisallowedSources
	obj.Campaigns = inp.Campaigns
}

func FillZoneUpdateInputModel(inp gqlmodels.ZoneUpdateInput, obj *models.Zone) {
	if obj == nil {
		return
	}
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

func activeStatusPtrFromBlazeGraphQL(status *bzgqlmodels.ActiveStatus) *types.ActiveStatus {
	if status == nil {
		return nil
	}
	st := activeStatusFromBlazeGraphQL(*status)
	return &st
}

func activeStatusFromBlazeGraphQL(status bzgqlmodels.ActiveStatus) types.ActiveStatus {
	if status == bzgqlmodels.ActiveStatusActive {
		return types.StatusActive
	}
	return types.StatusPause
}
