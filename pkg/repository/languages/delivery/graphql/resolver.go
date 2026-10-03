package graphql

import (
	"context"
	"slices"
	"strings"

	"github.com/demdxx/langlib"
	"github.com/demdxx/xtypes"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type QueryResolver struct {
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{}
}

func (r *QueryResolver) Languages(ctx context.Context, filter *gqlmodels.LangListFilter) ([]*gqlmodels.Lang, error) {
	list := langlib.Languages

	if filter != nil {
		list = xtypes.Slice[langlib.Language](langlib.Languages).
			Filter(func(lang langlib.Language) bool {
				if slices.Contains(filter.ID, uint64(lang.ID)) {
					return true
				}
				if slices.Contains(filter.Iso2, string(lang.Code[:])) {
					return true
				}
				lngLow := strings.ToLower(lang.Name)
				if slices.ContainsFunc(filter.Name, func(v string) bool { return lngLow == strings.ToLower(v) }) {
					return true
				}
				if len(filter.NativeName) > 0 {
					nlw := strings.ToLower(lang.NativeName)
					if slices.ContainsFunc(filter.NativeName, func(v string) bool { return nlw == strings.ToLower(v) }) {
						return true
					}
				}
				return true
			})
	}

	return xtypes.SliceApply(list, func(lang langlib.Language) *gqlmodels.Lang {
		return &gqlmodels.Lang{
			ID:         uint64(lang.ID),
			Iso2:       string(lang.Code[:]),
			Name:       lang.Name,
			NativeName: lang.NativeName,
		}
	}), nil
}
