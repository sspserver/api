package models

import (
	"strings"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/internal/repository/os"
	"github.com/sspserver/api/models"
)

func FromOSModel(os *models.OS) *Os {
	if os == nil {
		return nil
	}
	return &Os{
		ID:                 os.ID,
		Name:               os.Name,
		Version:            os.Version.String(),
		Description:        os.Description,
		Active:             FromActiveStatus(os.Active),
		MatchNameExp:       os.MatchNameExp,
		MatchUserAgentExp:  os.MatchUserAgentExp,
		MatchVersionMinExp: os.MatchVersionMinExp,
		MatchVersionMaxExp: os.MatchVersionMaxExp,
		YearRelease:        os.YearRelease,
		YearEndSupport:     os.YearEndSupport,
		ParentID:           os.ParentID.V,
		Parent:             FromOSModel(os.Parent),
		Versions:           xtypes.SliceApply(os.Versions, FromOSModel),
		CreatedAt:          os.CreatedAt,
		UpdatedAt:          os.UpdatedAt,
		DeletedAt:          DeletedAt(os.DeletedAt),
	}
}

func FromOSModelList(os []*models.OS) []*Os {
	return xtypes.SliceApply(os, FromOSModel)
}

func (fl *OSListFilter) Filter() *os.Filter {
	if fl == nil {
		return nil
	}
	return &os.Filter{
		ID:       fl.ID,
		ParentID: fl.ParentID,
		Name:     fl.Name,
		Active: gocast.IfThenExec(fl.Active != nil,
			func() *types.ActiveStatus {
				st := ActiveStatusFrom(*fl.Active)
				return &st
			},
			func() *types.ActiveStatus { return nil }),
	}
}

func (ol *OSListOrder) Order() *os.ListOrder {
	if ol == nil {
		return nil
	}
	return &os.ListOrder{
		ID:        ol.ID.AsOrder(),
		Name:      ol.Name.AsOrder(),
		Active:    ol.Active.AsOrder(),
		CreatedAt: ol.CreatedAt.AsOrder(),
		UpdatedAt: ol.UpdatedAt.AsOrder(),
	}
}

func (ol *OSListOrder) Fill(order *os.ListOrder) {
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

func (inp *OSCreateInput) FillModel(trg *models.OS) {
	if trg == nil {
		return
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
}

func (inp *OSUpdateInput) FillModel(trg *models.OS) {
	if trg == nil {
		return
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
}
