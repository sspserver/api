package graphql

import (
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromBrowserModel(browser *models.Browser) *gqlmodels.Browser {
	if browser == nil {
		return nil
	}
	return &gqlmodels.Browser{
		ID:                 browser.ID,
		Name:               browser.Name,
		Version:            browser.Version.String(),
		Description:        browser.Description,
		Active:             gqlmodels.FromActiveStatus(browser.Active),
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
		DeletedAt:          gqlmodels.DeletedAt(browser.DeletedAt),
	}
}

func FromBrowserModelList(browser []*models.Browser) []*gqlmodels.Browser {
	return xtypes.SliceApply(browser, FromBrowserModel)
}
