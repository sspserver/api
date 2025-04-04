package graphql

import (
	"context"
	"strings"

	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/adcorelib/i18n/languages"
	gqlmodels "github.com/sspserver/api/internal/server/graphql/models"
	"golang.org/x/exp/slices"
)

type QueryResolver struct {
}

func NewQueryResolver() *QueryResolver {
	return &QueryResolver{}
}

func (r *QueryResolver) Languages(ctx context.Context, filter *gqlmodels.LangListFilter) ([]*gqlmodels.Lang, error) {
	list := languages.Languages

	if filter != nil {
		list = xtypes.Slice[languages.Language](languages.Languages).
			Filter(func(lang languages.Language) bool {
				if len(filter.ID) > 0 && slices.Contains(filter.ID, uint64(lang.ID)) {
					return true
				}
				if len(filter.Iso2) > 0 && slices.Contains(filter.Iso2, string(lang.Code[:])) {
					return true
				}
				if len(filter.Name) > 0 {
					nlw := strings.ToLower(lang.Name)
					if slices.ContainsFunc(filter.Name, func(v string) bool { return nlw == strings.ToLower(v) }) {
						return true
					}
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

	return xtypes.SliceApply(list, func(lang languages.Language) *gqlmodels.Lang {
		return &gqlmodels.Lang{
			ID:         uint64(lang.ID),
			Iso2:       string(lang.Code[:]),
			Name:       lang.Name,
			NativeName: lang.NativeName,
		}
	}), nil
}
