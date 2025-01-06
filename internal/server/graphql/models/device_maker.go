package models

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/internal/repository/devicemaker"
	"github.com/sspserver/api/models"
)

func FromDeviceMakerModel(m *models.DeviceMaker) *DeviceMaker {
	if m == nil {
		return nil
	}
	processed := map[uint64]bool{}
	return &DeviceMaker{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		MatchExp:    m.MatchExp,
		Types: xtypes.SliceApply(m.Models, func(m *models.DeviceModel) *DeviceType {
			if m != nil && m.Type != nil && !processed[m.Type.ID] {
				processed[m.Type.ID] = true
				return FromDeviceTypeModel(m.Type)
			}
			return nil
		}).Filter(func(m *DeviceType) bool { return m != nil }),
		Models:    FromDeviceModelModelList(m.Models),
		Active:    FromActiveStatus(m.Active),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: DeletedAt(m.DeletedAt),
	}
}

func FromDeviceMakerModelList(list []*models.DeviceMaker) []*DeviceMaker {
	return xtypes.SliceApply(list, FromDeviceMakerModel)
}

func (fl *DeviceMakerListFilter) Filter() *devicemaker.Filter {
	if fl == nil {
		return nil
	}
	return &devicemaker.Filter{
		ID:   fl.ID,
		Name: fl.Name,
		Active: gocast.IfThenExec(len(fl.Active) > 0, func() *types.ActiveStatus {
			st := ActiveStatusFrom(fl.Active[0])
			return &st
		}, func() *types.ActiveStatus { return nil }),
	}
}

func (ol *DeviceMakerListOrder) Order() *devicemaker.ListOrder {
	if ol == nil {
		return nil
	}
	return &devicemaker.ListOrder{
		ID:        ol.ID.AsOrder(),
		Name:      ol.Name.AsOrder(),
		Active:    ol.Active.AsOrder(),
		CreatedAt: ol.CreatedAt.AsOrder(),
		UpdatedAt: ol.UpdatedAt.AsOrder(),
	}
}

func (inp *DeviceMakerInput) FillModel(m *models.DeviceMaker) {
	if inp == nil || m == nil {
		return
	}
	m.Name = gocast.PtrAsValue(inp.Name, m.Name)
	m.Description = gocast.PtrAsValue(inp.Description, m.Description)
	m.MatchExp = gocast.PtrAsValue(inp.MatchExp, m.MatchExp)
	m.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), m.Active)
}
