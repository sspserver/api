package models

import (
	"fmt"

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
	return &DeviceMaker{
		ID:          m.ID,
		Codename:    m.Codename,
		Name:        m.Name,
		Description: m.Description,
		MatchExp:    m.MatchExp,
		Models:      FromDeviceModelModelList(m.Models),
		Active:      FromActiveStatus(m.Active),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   DeletedAt(m.DeletedAt),
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
		Codename:  ol.Codename.AsOrder(),
		Name:      ol.Name.AsOrder(),
		Active:    ol.Active.AsOrder(),
		CreatedAt: ol.CreatedAt.AsOrder(),
		UpdatedAt: ol.UpdatedAt.AsOrder(),
	}
}

func (ol *DeviceMakerListOrder) Fill(vl *devicemaker.ListOrder) {
	if ol == nil || vl == nil {
		return
	}
	vl.ID = ol.ID.AsOrder()
	vl.Codename = ol.Codename.AsOrder()
	vl.Name = ol.Name.AsOrder()
	vl.Active = ol.Active.AsOrder()
	vl.CreatedAt = ol.CreatedAt.AsOrder()
	vl.UpdatedAt = ol.UpdatedAt.AsOrder()
}

func (inp *DeviceMakerCreateInput) FillModel(m *models.DeviceMaker) error {
	if inp == nil || m == nil {
		return nil
	}
	if inp.Name == "" {
		return fmt.Errorf("name is required")
	}
	if inp.Codename == "" {
		return fmt.Errorf("codename is required")
	}
	m.Name = inp.Name
	m.Codename = inp.Codename
	m.Description = gocast.PtrAsValue(inp.Description, "")
	m.MatchExp = gocast.PtrAsValue(inp.MatchExp, "")
	m.Active = ActiveStatusFrom(inp.Active)
	return nil
}

func (inp *DeviceMakerUpdateInput) FillModel(m *models.DeviceMaker) {
	if inp == nil || m == nil {
		return
	}
	m.Name = gocast.PtrAsValue(inp.Name, m.Name)
	m.Codename = gocast.PtrAsValue(inp.Codename, m.Codename)
	m.Description = gocast.PtrAsValue(inp.Description, m.Description)
	m.MatchExp = gocast.PtrAsValue(inp.MatchExp, m.MatchExp)
	m.Active = gocast.PtrAsValue(ActiveStatusPtr(inp.Active), m.Active)
}
