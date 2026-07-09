package graphql

import (
	"strings"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/pkg/repository/application"
	"github.com/sspserver/api/pkg/repository/application/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func ApplicationFilterFromGraphQL(fl *gqlmodels.ApplicationListFilter) *application.Filter {
	if fl == nil {
		return nil
	}
	return &application.Filter{
		ID: fl.ID,
		AccountID: gocast.IfThenExec(fl.AccountID != nil,
			func() []uint64 { return []uint64{*fl.AccountID} },
			func() []uint64 { return nil }),
		Title:    gocast.PtrAsValue(fl.Title, ""),
		URI:      gocast.PtrAsValue(fl.URI, ""),
		Type:     xtypes.SliceApply(fl.Type, func(v gqlmodels.ApplicationType) types.ApplicationType { return v.ModelType() }),
		Platform: xtypes.SliceApply(fl.Platform, func(v gqlmodels.PlatformType) types.PlatformType { return v.ModelType() }),
		Permium:  fl.Premium,
		Status:   gqlmodels.ApproveStatusPtr(fl.Status),
		Active:   gqlmodels.ActiveStatusPtr(fl.Active),
	}
}

func ApplicationOrderFromGraphQL(ol *gqlmodels.ApplicationListOrder) *application.ListOrder {
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

func FillApplicationCreateInputModel(inp gqlmodels.ApplicationCreateInput, trg *models.Application) error {
	if trg == nil {
		return nil
	}
	trg.AccountID = gocast.PtrAsValue(inp.AccountID, trg.AccountID)
	trg.Title = strings.TrimSpace(inp.Title)
	trg.Description = strings.TrimSpace(gocast.PtrAsValue(inp.Description, trg.Description))

	trg.URI = strings.TrimSpace(inp.URI)
	trg.Type = inp.Type.ModelType()
	trg.Platform = inp.Platform.ModelType()

	trg.Categories = gocast.Slice[uint](inp.Categories)
	trg.RevenueShare = gocast.PtrAsValue(inp.RevenueShare, trg.RevenueShare)

	if trg.Title == "" {
		return gqlmodels.ErrorRequiredField("title")
	}
	if trg.URI == "" {
		return gqlmodels.ErrorRequiredField("uri")
	}
	return nil
}

func FillApplicationUpdateInputModel(inp gqlmodels.ApplicationUpdateInput, trg *models.Application) error {
	if trg == nil {
		return nil
	}
	trg.AccountID = gocast.PtrAsValue(inp.AccountID, trg.AccountID)
	trg.Title = strings.TrimSpace(gocast.PtrAsValue(inp.Title, trg.Title))
	trg.Description = strings.TrimSpace(gocast.PtrAsValue(inp.Description, trg.Description))

	trg.URI = strings.TrimSpace(gocast.PtrAsValue(inp.URI, trg.URI))
	trg.Type = gocast.IfThenExec(inp.Type != nil,
		func() types.ApplicationType { return inp.Type.ModelType() },
		func() types.ApplicationType { return trg.Type })
	trg.Platform = gocast.IfThenExec(inp.Platform != nil,
		func() types.PlatformType { return inp.Platform.ModelType() },
		func() types.PlatformType { return trg.Platform })

	trg.Categories = gocast.IfThen(inp.Categories != nil, gocast.Slice[uint](inp.Categories), trg.Categories)
	trg.RevenueShare = gocast.PtrAsValue(inp.RevenueShare, trg.RevenueShare)

	if trg.Title == "" {
		return gqlmodels.ErrorInvalidField("title", "can`t be empty")
	}
	if trg.URI == "" {
		return gqlmodels.ErrorInvalidField("uri", "can`t be empty")
	}
	return nil
}
