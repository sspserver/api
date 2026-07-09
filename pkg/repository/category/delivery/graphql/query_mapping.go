package graphql

import (
	"database/sql"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/pkg/models"
	"github.com/sspserver/api/pkg/repository/category"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FillCategoryInputModel(input gqlmodels.CategoryInput, trg *models.Category) {
	if trg == nil {
		return
	}
	trg.Name = gocast.PtrAsValue(input.Name, trg.Name)
	trg.Description = gocast.PtrAsValue(input.Description, trg.Description)
	trg.ParentID = sql.Null[uint64]{V: gocast.PtrAsValue(input.ParentID, 0), Valid: input.ParentID != nil && *input.ParentID > 0}
	trg.IABCode = gocast.PtrAsValue(input.IABCode, trg.IABCode)
	if input.Active != nil {
		trg.Active = activeStatusFromGraphQL(*input.Active)
	}
	trg.Position = uint64(gocast.PtrAsValue(input.Position, int(trg.Position)))
}

func CategoryFilterFromGraphQL(fl *gqlmodels.CategoryListFilter) *category.Filter {
	if fl == nil {
		return nil
	}
	return &category.Filter{
		ID:       fl.ID,
		Name:     fl.Name,
		ParentID: fl.ParentID,
		IABCode:  fl.IABCode,
		Active: gocast.IfThenExec(len(fl.Active) > 0,
			func() *types.ActiveStatus { return &[]types.ActiveStatus{activeStatusFromGraphQL(fl.Active[0])}[0] },
			func() *types.ActiveStatus { return nil },
		),
	}
}

func activeStatusFromGraphQL(status gqlmodels.ActiveStatus) types.ActiveStatus {
	if status == gqlmodels.ActiveStatusActive {
		return models.StatusActive
	}
	return models.StatusPause
}

func CategoryOrderFromGraphQL(src *gqlmodels.CategoryListOrder) *category.ListOrder {
	if src == nil {
		return nil
	}
	return &category.ListOrder{
		ID:        src.ID.AsOrder(),
		Name:      src.Name.AsOrder(),
		IABCode:   src.IABCode.AsOrder(),
		ParentID:  src.ParentID.AsOrder(),
		Position:  src.Position.AsOrder(),
		Active:    src.Active.AsOrder(),
		CreatedAt: src.CreatedAt.AsOrder(),
		UpdatedAt: src.UpdatedAt.AsOrder(),
	}
}
