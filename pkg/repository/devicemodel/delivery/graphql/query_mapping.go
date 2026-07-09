package graphql

import (
	"errors"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/pkg/repository/devicemodel"
	"github.com/sspserver/api/pkg/repository/devicemodel/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func DeviceModelFilterFromGraphQL(fl *gqlmodels.DeviceModelListFilter) *devicemodel.Filter {
	if fl == nil {
		return nil
	}
	return &devicemodel.Filter{
		ID:       fl.ID,
		Codename: fl.Codename,
		Name:     fl.Name,
		ParentID: fl.ParentID,
		Active: gocast.IfThenExec(len(fl.Active) > 0, func() *types.ActiveStatus {
			st := activeStatusFromGraphQL(fl.Active[0])
			return &st
		}, func() *types.ActiveStatus { return nil }),
		MakerCodename: fl.MakerCodename,
		TypeCodename:  fl.TypeCodename,
	}
}

func FillDeviceModelListOrderFromGraphQL(src *gqlmodels.DeviceModelListOrder, dst *devicemodel.ListOrder) {
	if src == nil || dst == nil {
		return
	}
	dst.ID = src.ID.AsOrder()
	dst.Codename = src.Codename.AsOrder()
	dst.Name = src.Name.AsOrder()
	dst.Active = src.Active.AsOrder()
	dst.TypeCodename = src.TypeCodename.AsOrder()
	dst.MakerCodename = src.MakerCodename.AsOrder()
	dst.CreatedAt = src.CreatedAt.AsOrder()
	dst.UpdatedAt = src.UpdatedAt.AsOrder()
	dst.YearRelease = src.YearRelease.AsOrder()
}

func DeviceModelListOrderFromGraphQL(order []*gqlmodels.DeviceModelListOrder) devicemodel.ListOrder {
	var out devicemodel.ListOrder
	for _, item := range order {
		FillDeviceModelListOrderFromGraphQL(item, &out)
	}
	return out
}

func FillDeviceModelCreateInputModel(input gqlmodels.DeviceModelCreateInput, trg *models.DeviceModel) error {
	if trg == nil {
		return nil
	}
	if input.Codename == "" {
		return errors.New("codename is required")
	}
	if input.Name == "" {
		return errors.New("name is required")
	}
	trg.Codename = input.Codename
	trg.ParentID = gocast.PtrAsValue(input.ParentID, trg.ParentID)
	trg.Name = input.Name
	trg.Description = gocast.PtrAsValue(input.Description, trg.Description)
	trg.MatchExp = gocast.PtrAsValue(input.MatchExp, trg.MatchExp)
	trg.TypeCodename = input.TypeCodename
	trg.MakerCodename = input.MakerCodename
	trg.Active = activeStatusFromGraphQL(input.Active)
	return nil
}

func FillDeviceModelUpdateInputModel(input gqlmodels.DeviceModelUpdateInput, trg *models.DeviceModel) {
	if trg == nil {
		return
	}
	trg.Codename = gocast.PtrAsValue(input.Codename, trg.Codename)
	trg.ParentID = gocast.PtrAsValue(input.ParentID, trg.ParentID)
	trg.Name = gocast.PtrAsValue(input.Name, trg.Name)
	trg.Description = gocast.PtrAsValue(input.Description, trg.Description)
	trg.MatchExp = gocast.PtrAsValue(input.MatchExp, trg.MatchExp)
	trg.TypeCodename = gocast.PtrAsValue(input.TypeCodename, trg.TypeCodename)
	trg.MakerCodename = gocast.PtrAsValue(input.MakerCodename, trg.MakerCodename)
	if input.Active != nil {
		trg.Active = activeStatusFromGraphQL(*input.Active)
	}
}

func activeStatusFromGraphQL(status gqlmodels.ActiveStatus) types.ActiveStatus {
	if status == gqlmodels.ActiveStatusActive {
		return types.StatusActive
	}
	return types.StatusPause
}
