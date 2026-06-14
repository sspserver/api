package graphql

import (
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromCategoryModel(category *models.Category) *gqlmodels.Category {
	if category == nil {
		return nil
	}
	return &gqlmodels.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ParentID: gocast.IfThenExec(category.ParentID.V > 0,
			func() *uint64 { return &[]uint64{category.ParentID.V}[0] },
			func() *uint64 { return nil }),
		Parent:    FromCategoryModel(category.Parent),
		Childrens: xtypes.SliceApply(category.Childrens, FromCategoryModel),
		IABCode:   category.IABCode,
		Active:    gqlmodels.FromActiveStatus(category.Active),
		Position:  int(category.Position),
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		DeletedAt: gqlmodels.DeletedAt(category.DeletedAt),
	}
}

func FromCategoryModelList(categories []*models.Category) []*gqlmodels.Category {
	return xtypes.SliceApply(categories, FromCategoryModel)
}
