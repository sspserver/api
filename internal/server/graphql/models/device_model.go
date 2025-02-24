package models

import (
	"errors"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/internal/repository/devicemodel"
	"github.com/sspserver/api/models"
)

func FromDeviceModelModel(m *models.DeviceModel) *DeviceModel {
	if m == nil {
		return nil
	}
	return &DeviceModel{
		ID:            m.ID,
		Name:          m.Name,
		Codename:      m.Codename,
		Description:   m.Description,
		Active:        FromActiveStatus(m.Active),
		MatchExp:      m.MatchExp,
		ParentID:      gocast.IfThen(m.ParentID > 0, &m.ParentID, nil),
		MakerCodename: m.MakerCodename,
		Maker:         FromDeviceMakerModel(m.Maker),
		TypeCodename:  m.TypeCodename,
		Type:          FromDeviceTypeModel(m.Type),
		Versions:      xtypes.SliceApply(m.Versions, FromDeviceModelModel),
		CreatedAt:     m.CreatedAt,
		UpdatedAt:     m.UpdatedAt,
		DeletedAt:     DeletedAt(m.DeletedAt),
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

func (inp *DeviceModelCreateInput) FillModel(m *models.DeviceModel) error {
	if inp == nil || m == nil {
		return nil
	}
	if inp.Codename == "" {
		return errors.New("codename is required")
	}
	if inp.Name == "" {
		return errors.New("name is required")
	}
	m.Codename = inp.Codename
	m.ParentID = gocast.PtrAsValue(inp.ParentID, m.ParentID)
	m.Name = inp.Name
	m.Description = gocast.PtrAsValue(inp.Description, m.Description)
	m.MatchExp = gocast.PtrAsValue(inp.MatchExp, m.MatchExp)
	m.TypeCodename = inp.TypeCodename
	m.MakerCodename = inp.MakerCodename
	m.Active = ActiveStatusFrom(inp.Active)
	return nil
}

func (inp *DeviceModelUpdateInput) FillModel(m *models.DeviceModel) {
	if inp == nil || m == nil {
		return
	}
	m.Codename = gocast.PtrAsValue(inp.Codename, m.Codename)
	m.ParentID = gocast.PtrAsValue(inp.ParentID, m.ParentID)
	m.Name = gocast.PtrAsValue(inp.Name, m.Name)
	m.Description = gocast.PtrAsValue(inp.Description, m.Description)
	m.MatchExp = gocast.PtrAsValue(inp.MatchExp, m.MatchExp)
	m.TypeCodename = gocast.PtrAsValue(inp.TypeCodename, m.TypeCodename)
	m.MakerCodename = gocast.PtrAsValue(inp.MakerCodename, m.MakerCodename)
	m.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), m.Active)
}
