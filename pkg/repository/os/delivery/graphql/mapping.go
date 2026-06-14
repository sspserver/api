package graphql

import (
	"github.com/demdxx/xtypes"

	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

func FromOSModel(os *models.OS) *gqlmodels.Os {
	if os == nil {
		return nil
	}
	return &gqlmodels.Os{
		ID:                 os.ID,
		Name:               os.Name,
		Version:            os.Version.String(),
		Description:        os.Description,
		Active:             gqlmodels.FromActiveStatus(os.Active),
		MatchNameExp:       os.MatchNameExp,
		MatchUserAgentExp:  os.MatchUserAgentExp,
		MatchVersionMinExp: os.MatchVersionMinExp,
		MatchVersionMaxExp: os.MatchVersionMaxExp,
		YearRelease:        os.YearRelease,
		YearEndSupport:     os.YearEndSupport,
		ParentID:           os.ParentID.V,
		Parent:             FromOSModel(os.Parent),
		Versions:           xtypes.SliceApply(os.Versions, FromOSModel),
		CreatedAt:          os.CreatedAt,
		UpdatedAt:          os.UpdatedAt,
		DeletedAt:          gqlmodels.DeletedAt(os.DeletedAt),
	}
}

func FromOSModelList(os []*models.OS) []*gqlmodels.Os {
	return xtypes.SliceApply(os, FromOSModel)
}
