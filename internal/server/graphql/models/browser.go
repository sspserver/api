package models

import (
	"strings"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/internal/repository/browser"
	"github.com/sspserver/api/models"
)

func FromBrowserModel(browser *models.Browser) *Browser {
	if browser == nil {
		return nil
	}
	return &Browser{
		ID:                 browser.ID,
		Name:               browser.Name,
		Version:            browser.Version.String(),
		Description:        browser.Description,
		Active:             FromActiveStatus(browser.Active),
		MatchNameExp:       browser.MatchNameExp,
		MatchUserAgentExp:  browser.MatchUserAgentExp,
		MatchVersionMinExp: browser.MatchVersionMinExp,
		MatchVersionMaxExp: browser.MatchVersionMaxExp,
		YearRelease:        browser.YearRelease,
		YearEndSupport:     browser.YearEndSupport,
		ParentID:           browser.ParentID.V,
		Parent:             FromBrowserModel(browser.Parent),
		Versions:           xtypes.SliceApply(browser.Versions, FromBrowserModel),
		CreatedAt:          browser.CreatedAt,
		UpdatedAt:          browser.UpdatedAt,
		DeletedAt:          DeletedAt(browser.DeletedAt),
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
		ID:       fl.ID,
		ParentID: fl.ParentID,
		Name:     fl.Name,
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

func (ol *BrowserListOrder) Fill(order *browser.ListOrder) {
	if ol == nil || order == nil {
		return
	}
	if ol.ID != nil {
		order.ID = ol.ID.AsOrder()
	}
	if ol.Name != nil {
		order.Name = ol.Name.AsOrder()
	}
	if ol.Active != nil {
		order.Active = ol.Active.AsOrder()
	}
	if ol.CreatedAt != nil {
		order.CreatedAt = ol.CreatedAt.AsOrder()
	}
	if ol.UpdatedAt != nil {
		order.UpdatedAt = ol.UpdatedAt.AsOrder()
	}
	if ol.YearRelease != nil {
		order.YearRelease = ol.YearRelease.AsOrder()
	}
}

func (inp *BrowserCreateInput) FillModel(trg *models.Browser) error {
	if trg == nil {
		return nil
	}
	trg.Name = strings.TrimSpace(inp.Name)
	trg.Version = types.IgnoreParseVersion(gocast.PtrAsValue(inp.Version, trg.Version.String()))
	trg.Description = gocast.PtrAsValue(inp.Description, trg.Description)
	trg.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), trg.Active)
	trg.MatchNameExp = gocast.PtrAsValue(inp.MatchNameExp, trg.MatchNameExp)
	trg.MatchUserAgentExp = gocast.PtrAsValue(inp.MatchUserAgentExp, trg.MatchUserAgentExp)
	trg.MatchVersionMinExp = gocast.PtrAsValue(inp.MatchVersionMinExp, trg.MatchVersionMinExp)
	trg.MatchVersionMaxExp = gocast.PtrAsValue(inp.MatchVersionMaxExp, trg.MatchVersionMaxExp)
	trg.YearRelease = gocast.PtrAsValue(inp.YearRelease, trg.YearRelease)
	trg.YearEndSupport = gocast.PtrAsValue(inp.YearEndSupport, trg.YearEndSupport)

	if trg.Name == "" {
		return ErrorRequiredField("name")
	}
	return nil
}

func (inp *BrowserUpdateInput) FillModel(trg *models.Browser) error {
	if trg == nil {
		return nil
	}
	trg.Name = gocast.Or(strings.TrimSpace(gocast.PtrAsValue(inp.Name, trg.Name)), trg.Name)
	trg.Version = types.IgnoreParseVersion(gocast.PtrAsValue(inp.Version, trg.Version.String()))
	trg.Description = gocast.PtrAsValue(inp.Description, trg.Description)
	trg.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), trg.Active)
	trg.MatchNameExp = gocast.PtrAsValue(inp.MatchNameExp, trg.MatchNameExp)
	trg.MatchUserAgentExp = gocast.PtrAsValue(inp.MatchUserAgentExp, trg.MatchUserAgentExp)
	trg.MatchVersionMinExp = gocast.PtrAsValue(inp.MatchVersionMinExp, trg.MatchVersionMinExp)
	trg.MatchVersionMaxExp = gocast.PtrAsValue(inp.MatchVersionMaxExp, trg.MatchVersionMaxExp)
	trg.YearRelease = gocast.PtrAsValue(inp.YearRelease, trg.YearRelease)
	trg.YearEndSupport = gocast.PtrAsValue(inp.YearEndSupport, trg.YearEndSupport)

	if trg.Name == "" {
		return ErrorInvalidField("name", "can`t be empty")
	}
	return nil
}
