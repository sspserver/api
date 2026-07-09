package graphql

import (
	"strings"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/browser"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func BrowserFilterFromGraphQL(fl *gqlmodels.BrowserListFilter) *browser.Filter {
	if fl == nil {
		return nil
	}
	return &browser.Filter{
		ID:       fl.ID,
		ParentID: fl.ParentID,
		Name:     fl.Name,
		Active: gocast.IfThenExec(len(fl.Active) > 0, func() *types.ActiveStatus {
			st := gqlmodels.ActiveStatusFrom(fl.Active[0])
			return &st
		}, func() *types.ActiveStatus { return nil }),
	}
}

func FillBrowserListOrderFromGraphQL(src *gqlmodels.BrowserListOrder, dst *browser.ListOrder) {
	if src == nil || dst == nil {
		return
	}
	if src.ID != nil {
		dst.ID = src.ID.AsOrder()
	}
	if src.Name != nil {
		dst.Name = src.Name.AsOrder()
	}
	if src.Active != nil {
		dst.Active = src.Active.AsOrder()
	}
	if src.CreatedAt != nil {
		dst.CreatedAt = src.CreatedAt.AsOrder()
	}
	if src.UpdatedAt != nil {
		dst.UpdatedAt = src.UpdatedAt.AsOrder()
	}
	if src.YearRelease != nil {
		dst.YearRelease = src.YearRelease.AsOrder()
	}
}

func BrowserListOrderFromGraphQL(order []*gqlmodels.BrowserListOrder) browser.ListOrder {
	var out browser.ListOrder
	for _, item := range order {
		FillBrowserListOrderFromGraphQL(item, &out)
	}
	return out
}

func FillBrowserCreateInputModel(input gqlmodels.BrowserCreateInput, trg *models.Browser) error {
	if trg == nil {
		return nil
	}
	trg.Name = strings.TrimSpace(input.Name)
	trg.Version = types.IgnoreParseVersion(gocast.PtrAsValue(input.Version, trg.Version.String()))
	trg.Description = gocast.PtrAsValue(input.Description, trg.Description)
	trg.Active = gocast.PtrAsValue(gqlmodels.ActiveStatusPtr(input.Active), trg.Active)
	trg.MatchNameExp = gocast.PtrAsValue(input.MatchNameExp, trg.MatchNameExp)
	trg.MatchUserAgentExp = gocast.PtrAsValue(input.MatchUserAgentExp, trg.MatchUserAgentExp)
	trg.MatchVersionMinExp = gocast.PtrAsValue(input.MatchVersionMinExp, trg.MatchVersionMinExp)
	trg.MatchVersionMaxExp = gocast.PtrAsValue(input.MatchVersionMaxExp, trg.MatchVersionMaxExp)
	trg.YearRelease = gocast.PtrAsValue(input.YearRelease, trg.YearRelease)
	trg.YearEndSupport = gocast.PtrAsValue(input.YearEndSupport, trg.YearEndSupport)

	if trg.Name == "" {
		return gqlmodels.ErrorRequiredField("name")
	}
	return nil
}

func FillBrowserUpdateInputModel(input gqlmodels.BrowserUpdateInput, trg *models.Browser) error {
	if trg == nil {
		return nil
	}
	trg.Name = gocast.Or(strings.TrimSpace(gocast.PtrAsValue(input.Name, trg.Name)), trg.Name)
	trg.Version = types.IgnoreParseVersion(gocast.PtrAsValue(input.Version, trg.Version.String()))
	trg.Description = gocast.PtrAsValue(input.Description, trg.Description)
	trg.Active = gocast.PtrAsValue(gqlmodels.ActiveStatusPtr(input.Active), trg.Active)
	trg.MatchNameExp = gocast.PtrAsValue(input.MatchNameExp, trg.MatchNameExp)
	trg.MatchUserAgentExp = gocast.PtrAsValue(input.MatchUserAgentExp, trg.MatchUserAgentExp)
	trg.MatchVersionMinExp = gocast.PtrAsValue(input.MatchVersionMinExp, trg.MatchVersionMinExp)
	trg.MatchVersionMaxExp = gocast.PtrAsValue(input.MatchVersionMaxExp, trg.MatchVersionMaxExp)
	trg.YearRelease = gocast.PtrAsValue(input.YearRelease, trg.YearRelease)
	trg.YearEndSupport = gocast.PtrAsValue(input.YearEndSupport, trg.YearEndSupport)

	if trg.Name == "" {
		return gqlmodels.ErrorInvalidField("name", "can`t be empty")
	}
	return nil
}
