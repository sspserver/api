package statistic

import (
	"context"

	"github.com/geniusrabbit/archivarius/client"
	"github.com/geniusrabbit/blaze-api/pkg/acl"
	"github.com/geniusrabbit/blaze-api/pkg/context/session"

	gqmodels "github.com/sspserver/api/internal/server/graphql/models"
)

// RBACStatisic is a resource for statistic
var RBACStatisic = acl.RBACType{ResourceName: "statistic"}

// Usecase is the interface for the statistic usecase.
type UsecaseImpl struct {
	repo Repository
}

// NewUsecase is the constructor for the statistic usecase.
func NewUsecase(repo Repository) Usecase {
	return &UsecaseImpl{repo: repo}
}

// Statistic returns the statistic of the ads
func (u *UsecaseImpl) Statistic(
	ctx context.Context,
	group []gqmodels.StatisticKey,
	order []*gqmodels.StatisticAdKeyOrder,
	filter *gqmodels.StatisticAdListFilter,
	page *gqmodels.Page,
) (*client.StatisticResponse, error) {
	if !acl.HaveAccessList(ctx, &RBACStatisic) {
		if acl.HaveAccessList(ctx, RBACStatisic.WithAccountID(session.Account(ctx).ID)) {
			filter.Conditions = append(filter.Conditions, &gqmodels.StatisticAdKeyCondition{
				Key:   gqmodels.StatisticKeyAccountID, // Pub or Adv account ID
				Op:    gqmodels.StatisticConditionEq,
				Value: []any{session.Account(ctx).ID},
			})
		} else {
			return nil, acl.ErrNoPermissions.WithMessage("Statistic")
		}
	}
	return u.repo.Statistic(ctx, group, order, filter, page)
}
