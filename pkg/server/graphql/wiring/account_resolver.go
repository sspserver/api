package wiring

import (
	"context"

	accountgraphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"

	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// EmailPasswordLoginHandler handles the login(email, password, accountID) mutation.
type EmailPasswordLoginHandler interface {
	Login(ctx context.Context, email, password string, accountID ...uint64) (*gqlmodels.SessionToken, error)
}

// AccountQueryResolver is account QueryResolver wired to extended GraphQL models.
type AccountQueryResolver = accountgraphql.QueryResolver[
	*models.User,
	*models.Account,
	*gqlmodels.Account,
	*gqlmodels.AccountPayload,
	*gqlmodels.AccountCreateInput,
	*gqlmodels.AccountUpdateInput,
	*gqlmodels.AccountListFilter,
	*gqlmodels.AccountListOrder,
	*gqlmodels.User,
	*gqlmodels.UserCreateInput,
	*gqlmodels.UserUpdateInput,
]

// AccountQueryResolverConfig wires extended account GraphQL models.
type AccountQueryResolverConfig = accountgraphql.QueryResolverConfig[
	*models.User,
	*models.Account,
	*gqlmodels.Account,
	*gqlmodels.AccountPayload,
	*gqlmodels.AccountCreateInput,
	*gqlmodels.AccountUpdateInput,
	*gqlmodels.AccountListFilter,
	*gqlmodels.AccountListOrder,
	*gqlmodels.User,
	*gqlmodels.UserCreateInput,
	*gqlmodels.UserUpdateInput,
]

// NewAccountQueryResolver wires account resolvers with extended GraphQL models.
func NewAccountQueryResolver(cfg AccountQueryResolverConfig) *AccountQueryResolver {
	if cfg.AccountsMapper == nil {
		cfg.AccountsMapper = AccountGraphQLMappersImpl{}
	}
	if cfg.UsersMapper == nil {
		cfg.UsersMapper = UserGraphQLMappersImpl{}
	}
	return accountgraphql.NewQueryResolver(cfg)
}
