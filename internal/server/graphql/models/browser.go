package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/gosql/v2"

	"github.com/sspserver/api/internal/repository/browser"
	"github.com/sspserver/api/models"
)

func FromBrowserVersion(ver models.BrowserVersion) *BrowserVersion {
	return &BrowserVersion{
		Min:  ver.Min.String(),
		Max:  ver.Max.String(),
		Name: ver.Name,
	}
}

func FromBrowserModel(browser *models.Browser) *Browser {
	if browser == nil {
		return nil
	}
	return &Browser{
		ID:          browser.ID,
		Name:        browser.Name,
		Description: browser.Description,
		MatchExp:    browser.MatchExp,
		Active:      FromActiveStatus(browser.Active),
		Versions:    xtypes.SliceApply(browser.Versions, FromBrowserVersion),
		CreatedAt:   browser.CreatedAt,
		UpdatedAt:   browser.UpdatedAt,
		DeletedAt:   DeletedAt(browser.DeletedAt),
	}
}

func FromBrowserModelList(browser []*models.Browser) []*Browser {
	return xtypes.SliceApply(browser, FromBrowserModel)
}

func (fl *BrowserListFilter) Filter() *browser.Filter {
	if fl == nil {
		return nil
	}
	return &browser.Filter{
		ID:   fl.ID,
		Name: fl.Name,
		Active: gocast.IfThenExec(len(fl.Active) > 0, func() *types.ActiveStatus {
			st := ActiveStatusFrom(fl.Active[0])
			return &st
		}, func() *types.ActiveStatus { return nil }),
	}
}

func (ol *BrowserListOrder) Order() *browser.ListOrder {
	if ol == nil {
		return nil
	}
	return &browser.ListOrder{
		ID:        ol.ID.AsOrder(),
		Name:      ol.Name.AsOrder(),
		Active:    ol.Active.AsOrder(),
		CreatedAt: ol.CreatedAt.AsOrder(),
		UpdatedAt: ol.UpdatedAt.AsOrder(),
	}
}

func (inp *BrowserInput) FillModel(trg *models.Browser) {
	if inp == nil {
		return
	}
	trg.Name = gocast.PtrAsValue(inp.Name, trg.Name)
	trg.Description = gocast.PtrAsValue(inp.Description, trg.Description)
	trg.MatchExp = gocast.PtrAsValue(inp.MatchExp, trg.MatchExp)
	trg.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), trg.Active)
	trg.Versions = gosql.NullableJSONArray[models.BrowserVersion](
		xtypes.SliceApply(inp.Versions, func(v *BrowserVersionInput) models.BrowserVersion {
			return models.BrowserVersion{
				Min:  types.IgnoreParseVersion(gocast.PtrAsValue(v.Min, "")),
				Max:  types.IgnoreParseVersion(gocast.PtrAsValue(v.Max, "")),
				Name: gocast.PtrAsValue(v.Name, ""),
			}
		}),
	)
}
