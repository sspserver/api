package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/gosql/v2"

	"github.com/sspserver/api/internal/repository/devicemodel"
	"github.com/sspserver/api/models"
)

func FromDeviceModelVersion(ver models.DeviceModelVersion) *DeviceModelVersion {
	return &DeviceModelVersion{
		Min:  ver.Min.String(),
		Max:  ver.Max.String(),
		Name: ver.Name,
	}
}

func FromDeviceModelModel(m *models.DeviceModel) *DeviceModel {
	if m == nil {
		return nil
	}
	return &DeviceModel{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		MatchExp:    m.MatchExp,
		Active:      FromActiveStatus(m.Active),
		MakerID:     m.MakerID,
		Maker:       FromDeviceMakerModel(m.Maker),
		TypeID:      m.TypeID,
		Type:        FromDeviceTypeModel(m.Type),
		Versions:    xtypes.SliceApply(m.Versions, FromDeviceModelVersion),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   DeletedAt(m.DeletedAt),
	}
}

func FromDeviceModelModelList(list []*models.DeviceModel) []*DeviceModel {
	return xtypes.SliceApply(list, FromDeviceModelModel)
}

func (fl *DeviceModelListFilter) Filter() *devicemodel.Filter {
	if fl == nil {
		return nil
	}
	return &devicemodel.Filter{
		ID:   fl.ID,
		Name: fl.Name,
		Active: gocast.IfThenExec(len(fl.Active) > 0, func() *types.ActiveStatus {
			st := ActiveStatusFrom(fl.Active[0])
			return &st
		}, func() *types.ActiveStatus { return nil }),
	}
}

func (ol *DeviceModelListOrder) Order() *devicemodel.ListOrder {
	if ol == nil {
		return nil
	}
	return &devicemodel.ListOrder{
		ID:        ol.ID.AsOrder(),
		Name:      ol.Name.AsOrder(),
		Active:    ol.Active.AsOrder(),
		CreatedAt: ol.CreatedAt.AsOrder(),
		UpdatedAt: ol.UpdatedAt.AsOrder(),
	}
}

func (v *DeviceModelVersionInput) Model() models.DeviceModelVersion {
	if v == nil {
		return models.DeviceModelVersion{}
	}
	return models.DeviceModelVersion{
		Min:  types.IgnoreParseVersion(gocast.PtrAsValue(v.Min, "")),
		Max:  types.IgnoreParseVersion(gocast.PtrAsValue(v.Max, "")),
		Name: gocast.PtrAsValue(v.Name, ""),
	}
}

func (inp *DeviceModelInput) FillModel(m *models.DeviceModel) {
	if inp == nil || m == nil {
		return
	}
	m.Name = gocast.PtrAsValue(inp.Name, m.Name)
	m.Description = gocast.PtrAsValue(inp.Description, m.Description)
	m.MatchExp = gocast.PtrAsValue(inp.MatchExp, m.MatchExp)
	m.TypeID = gocast.PtrAsValue(inp.TypeID, m.TypeID)
	m.MakerID = gocast.PtrAsValue(inp.MakerID, m.MakerID)
	m.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), m.Active)
	m.Versions = gosql.NullableJSONArray[models.DeviceModelVersion](
		gocast.IfThen(inp.Versions != nil,
			xtypes.SliceApply(inp.Versions, func(v *DeviceModelVersionInput) models.DeviceModelVersion { return v.Model() }),
			xtypes.Slice[models.DeviceModelVersion](m.Versions)),
	)
}
