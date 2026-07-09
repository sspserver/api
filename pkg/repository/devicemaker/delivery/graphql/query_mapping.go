package graphql

import (
	"fmt"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/devicemaker"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func DeviceMakerFilterFromGraphQL(fl *gqlmodels.DeviceMakerListFilter) *devicemaker.Filter {
	if fl == nil {
		return nil
	}
	return &devicemaker.Filter{
		ID:   fl.ID,
		Name: fl.Name,
		Active: gocast.IfThenExec(len(fl.Active) > 0, func() *types.ActiveStatus {
			st := activeStatusFromGraphQL(fl.Active[0])
			return &st
		}, func() *types.ActiveStatus { return nil }),
	}
}

func FillDeviceMakerListOrderFromGraphQL(src *gqlmodels.DeviceMakerListOrder, dst *devicemaker.ListOrder) {
	if src == nil || dst == nil {
		return
	}
	dst.ID = src.ID.AsOrder()
	dst.Codename = src.Codename.AsOrder()
	dst.Name = src.Name.AsOrder()
	dst.Active = src.Active.AsOrder()
	dst.CreatedAt = src.CreatedAt.AsOrder()
	dst.UpdatedAt = src.UpdatedAt.AsOrder()
}

func DeviceMakerListOrderFromGraphQL(order []*gqlmodels.DeviceMakerListOrder) devicemaker.ListOrder {
	var out devicemaker.ListOrder
	for _, item := range order {
		FillDeviceMakerListOrderFromGraphQL(item, &out)
	}
	return out
}

func FillDeviceMakerCreateInputModel(input gqlmodels.DeviceMakerCreateInput, trg *models.DeviceMaker) error {
	if trg == nil {
		return nil
	}
	if input.Name == "" {
		return fmt.Errorf("name is required")
	}
	if input.Codename == "" {
		return fmt.Errorf("codename is required")
	}
	trg.Name = input.Name
	trg.Codename = input.Codename
	trg.Description = gocast.PtrAsValue(input.Description, "")
	trg.MatchExp = gocast.PtrAsValue(input.MatchExp, "")
	trg.Active = activeStatusFromGraphQL(input.Active)
	return nil
}

func FillDeviceMakerUpdateInputModel(input gqlmodels.DeviceMakerUpdateInput, trg *models.DeviceMaker) {
	if trg == nil {
		return
	}
	trg.Name = gocast.PtrAsValue(input.Name, trg.Name)
	trg.Codename = gocast.PtrAsValue(input.Codename, trg.Codename)
	trg.Description = gocast.PtrAsValue(input.Description, trg.Description)
	trg.MatchExp = gocast.PtrAsValue(input.MatchExp, trg.MatchExp)
	if input.Active != nil {
		trg.Active = activeStatusFromGraphQL(*input.Active)
	}
}

func activeStatusFromGraphQL(status gqlmodels.ActiveStatus) types.ActiveStatus {
	if status == gqlmodels.ActiveStatusActive {
		return models.StatusActive
	}
	return models.StatusPause
}
