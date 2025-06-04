package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"

	"github.com/sspserver/api/pkg/repository/agreement"
)

type RequireAgreementsFunc = func(ctx context.Context, obj any, next graphql.Resolver) (any, error)

func RequireAgreements(uc *agreement.UsecaseImpl) RequireAgreementsFunc {
	return func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
		// Check if the user has accepted the required agreements
		if agg, err := uc.NextAwailable(ctx); agg != nil || err != nil {
			// If there is an agreement that has not been accepted, return an error
			if err != nil {
				return nil, fmt.Errorf("failed to check agreements: %w", err)
			}
			return nil, fmt.Errorf("you must accept the agreements before proceeding")
		}
		return next(ctx)
	}
}
