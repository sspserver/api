package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/gosql/v2"

	"github.com/sspserver/api/internal/repository/os"
	"github.com/sspserver/api/models"
)

func FromOSVersion(ver models.OSVersion) *OSVersion {
	return &OSVersion{
		Min:  ver.Min.String(),
		Max:  ver.Max.String(),
		Name: ver.Name,
	}
}

func FromOSModel(os *models.OS) *Os {
	if os == nil {
		return nil
	}
	return &Os{
		ID:          os.ID,
		Name:        os.Name,
		Description: os.Description,
		MatchExp:    os.MatchExp,
		Active:      FromActiveStatus(os.Active),
		Versions:    xtypes.SliceApply(os.Versions, FromOSVersion),
		CreatedAt:   os.CreatedAt,
		UpdatedAt:   os.UpdatedAt,
		DeletedAt:   DeletedAt(os.DeletedAt),
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
		ID:   fl.ID,
		Name: fl.Name,
		Active: gocast.IfThenExec(len(fl.Active) > 0, func() *types.ActiveStatus {
			st := ActiveStatusFrom(fl.Active[0])
			return &st
		}, func() *types.ActiveStatus { return nil }),
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

func (inp *OSInput) FillModel(trg *models.OS) {
	if trg == nil {
		return
	}
	trg.Name = gocast.PtrAsValue(inp.Name, trg.Name)
	trg.Description = gocast.PtrAsValue(inp.Description, trg.Description)
	trg.MatchExp = gocast.PtrAsValue(inp.MatchExp, trg.MatchExp)
	trg.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), trg.Active)
	trg.Versions = gosql.NullableJSONArray[models.OSVersion](
		xtypes.SliceApply(inp.Versions, func(v *OSVersionInput) models.OSVersion {
			return models.OSVersion{
				Min:  types.IgnoreParseVersion(gocast.PtrAsValue(v.Min, "")),
				Max:  types.IgnoreParseVersion(gocast.PtrAsValue(v.Max, "")),
				Name: gocast.PtrAsValue(v.Name, ""),
			}
		}),
	)
}
