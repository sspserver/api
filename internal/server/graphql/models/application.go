package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/internal/repository/application"
	"github.com/sspserver/api/models"
)

func FromApplicationModel(obj *models.Application) *Application {
	return &Application{
		ID: obj.ID,

		AccountID: obj.AccountID,
		CreatorID: obj.CreatorID,

		Title:       obj.Title,
		Description: obj.Description,

		URI:      obj.URI,
		Type:     FromApplicationType(obj.Type),
		Platform: FromPlatformType(obj.Platform),
		Premium:  obj.Premium,

		Status:  FromApproveStatus(obj.Status),
		Active:  FromActiveStatus(obj.Active),
		Private: FromPrivateStatus(obj.Private),

		Categories:   xtypes.SliceApply(obj.Categories, func(v uint) int { return int(v) }),
		RevenueShare: gocast.IfThen(obj.RevenueShare > 0, &[]float64{obj.RevenueShare}[0], nil),

		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
	}
}

func FromApplicationModelList(list []*models.Application) []*Application {
	return xtypes.SliceApply(list, FromApplicationModel)
}

func (inp *ApplicationInput) FillModel(trg *models.Application) {
	if inp == nil || trg == nil {
		return
	}
	trg.Title = gocast.PtrAsValue(inp.Title, trg.Title)
	trg.Description = gocast.PtrAsValue(inp.Description, trg.Description)

	trg.URI = gocast.PtrAsValue(inp.URI, trg.URI)
	trg.Type = gocast.IfThen(inp.Type != nil, inp.Type.ModelType(), trg.Type)
	trg.Platform = gocast.IfThen(inp.Platform != nil, inp.Platform.ModelType(), trg.Platform)
	trg.Premium = gocast.PtrAsValue(inp.Premium, trg.Premium)

	trg.Status = gocast.PtrAsValue(ApproveStatusPtr(inp.Status), trg.Status)
	trg.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), trg.Active)
	trg.Private = gocast.IfThen(inp.Private != nil, inp.Private.ModelStatus(), trg.Private)

	trg.Categories = gocast.IfThen(inp.Categories != nil, gocast.Slice[uint](inp.Categories), trg.Categories)
	trg.RevenueShare = gocast.PtrAsValue(inp.RevenueShare, trg.RevenueShare)
}

func (fl *ApplicationListFilter) Filter() *application.Filter {
	if fl == nil {
		return nil
	}
	return &application.Filter{
		ID:       fl.ID,
		Title:    gocast.PtrAsValue(fl.Title, ""),
		URI:      gocast.PtrAsValue(fl.URI, ""),
		Type:     xtypes.SliceApply(fl.Type, func(v ApplicationType) models.ApplicationType { return v.ModelType() }),
		Platform: xtypes.SliceApply(fl.Platform, func(v PlatformType) models.PlatformType { return v.ModelType() }),
		Permium:  fl.Premium,
		Status:   ApproveStatusPtr(fl.Status),
		Active:   ActiveStatusPtr(fl.Active),
	}
}

func (ol *ApplicationListOrder) Order() *application.ListOrder {
	if ol == nil {
		return nil
	}
	return &application.ListOrder{
		ID: ol.ID.AsOrder(),

		Title: ol.Title.AsOrder(),
		URI:   ol.URI.AsOrder(),

		Type:     ol.Type.AsOrder(),
		Platform: ol.Platform.AsOrder(),
		Premium:  ol.Premium.AsOrder(),
		Status:   ol.Status.AsOrder(),
		Active:   ol.Active.AsOrder(),

		CreatedAt: ol.CreatedAt.AsOrder(),
		UpdatedAt: ol.UpdatedAt.AsOrder(),
	}
}
