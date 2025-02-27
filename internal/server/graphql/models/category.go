package models

import (
	"database/sql"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/admodels/types"

	"github.com/sspserver/api/internal/repository/category"
	"github.com/sspserver/api/models"
)

func FromCategoryModel(category *models.Category) *Category {
	if category == nil {
		return nil
	}
	return &Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID: gocast.IfThenExec(category.ParentID.V > 0,
			func() *uint64 { return &[]uint64{category.ParentID.V}[0] },
			func() *uint64 { return nil }),
		Parent:    FromCategoryModel(category.Parent),
		Childrens: xtypes.SliceApply(category.Childrens, FromCategoryModel),
		IABCode:   category.IABCode,
		Active:    FromActiveStatus(category.Active),
		Position:  int(category.Position),
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		DeletedAt: DeletedAt(category.DeletedAt),
	}
}

func FromCategoryModelList(categories []*models.Category) []*Category {
	return xtypes.SliceApply(categories, FromCategoryModel)
}

func (m *CategoryInput) FillModel(category *models.Category) {
	if category == nil {
		return
	}
	category.Name = gocast.PtrAsValue(m.Name, category.Name)
	category.Description = gocast.PtrAsValue(m.Description, category.Description)
	category.ParentID = sql.Null[uint64]{V: gocast.PtrAsValue(m.ParentID, 0), Valid: m.ParentID != nil && *m.ParentID > 0}
	category.IABCode = gocast.PtrAsValue(m.IABCode, category.IABCode)
	category.Active = gocast.PtrAsValue(ActiveStatusPtr(m.Active), category.Active)
	category.Position = uint64(gocast.PtrAsValue(m.Position, int(category.Position)))
}

func (fl *CategoryListFilter) Filter() *category.Filter {
	if fl == nil {
		return nil
	}
	return &category.Filter{
		ID:       fl.ID,
		Name:     fl.Name,
		ParentID: fl.ParentID,
		IABCode:  fl.IABCode,
		Active: gocast.IfThenExec(len(fl.Active) > 0,
			func() *types.ActiveStatus { return &[]models.ActiveStatus{ActiveStatusFrom(fl.Active[0])}[0] },
			func() *types.ActiveStatus { return nil },
		),
	}
}

func (ol *CategoryListOrder) Order() *category.ListOrder {
	if ol == nil {
		return nil
	}
	return &category.ListOrder{
		ID:        ol.ID.AsOrder(),
		Name:      ol.Name.AsOrder(),
		IABCode:   ol.IABCode.AsOrder(),
		ParentID:  ol.ParentID.AsOrder(),
		Position:  ol.Position.AsOrder(),
		Active:    ol.Active.AsOrder(),
		CreatedAt: ol.CreatedAt.AsOrder(),
		UpdatedAt: ol.UpdatedAt.AsOrder(),
	}
}
