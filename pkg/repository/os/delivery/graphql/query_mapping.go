package graphql

import (
	"strings"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	osrepo "github.com/sspserver/api/pkg/repository/os"
	"github.com/sspserver/api/pkg/repository/os/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func OSFilterFromGraphQL(fl *gqlmodels.OSListFilter) *osrepo.Filter {
	if fl == nil {
		return nil
	}
	return &osrepo.Filter{
		ID:       fl.ID,
		ParentID: fl.ParentID,
		Name:     fl.Name,
		Active: gocast.IfThenExec(fl.Active != nil,
			func() *types.ActiveStatus {
				st := activeStatusFromGraphQL(*fl.Active)
				return &st
			},
			func() *types.ActiveStatus { return nil }),
	}
}

func FillOSListOrderFromGraphQL(src *gqlmodels.OSListOrder, dst *osrepo.ListOrder) {
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

func OSListOrderFromGraphQL(order []*gqlmodels.OSListOrder) osrepo.ListOrder {
	var out osrepo.ListOrder
	for _, item := range order {
		FillOSListOrderFromGraphQL(item, &out)
	}
	return out
}

func FillOSCreateInputModel(input gqlmodels.OSCreateInput, trg *models.OS) error {
	if trg == nil {
		return nil
	}
	trg.Name = strings.TrimSpace(input.Name)
	trg.Version = types.IgnoreParseVersion(gocast.PtrAsValue(input.Version, trg.Version.String()))
	trg.Description = gocast.PtrAsValue(input.Description, trg.Description)
	if input.Active != nil {
		trg.Active = activeStatusFromGraphQL(*input.Active)
	}
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

func FillOSUpdateInputModel(input gqlmodels.OSUpdateInput, trg *models.OS) error {
	if trg == nil {
		return nil
	}
	trg.Name = gocast.Or(strings.TrimSpace(gocast.PtrAsValue(input.Name, trg.Name)), trg.Name)
	trg.Version = types.IgnoreParseVersion(gocast.PtrAsValue(input.Version, trg.Version.String()))
	trg.Description = gocast.PtrAsValue(input.Description, trg.Description)
	if input.Active != nil {
		trg.Active = activeStatusFromGraphQL(*input.Active)
	}
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

func activeStatusFromGraphQL(status gqlmodels.ActiveStatus) types.ActiveStatus {
	if status == gqlmodels.ActiveStatusActive {
		return types.StatusActive
	}
	return types.StatusPause
}
