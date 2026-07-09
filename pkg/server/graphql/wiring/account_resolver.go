package wiring

import (
	"context"

	accountgraphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"

	"github.com/sspserver/api/pkg/models"
	gqlaccounts "github.com/sspserver/api/pkg/server/graphql/accounts"
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
	*gqlaccounts.Account,
	*gqlaccounts.AccountPayload,
	*gqlaccounts.AccountCreateInput,
	*gqlaccounts.AccountUpdateInput,
	*gqlaccounts.AccountListFilter,
	*gqlaccounts.AccountListOrder,
	*gqlaccounts.User,
	*gqlaccounts.UserCreateInput,
	*gqlaccounts.UserUpdateInput,
]

// AccountQueryResolverConfig wires extended account GraphQL models.
type AccountQueryResolverConfig = accountgraphql.QueryResolverConfig[
	*models.User,
	*models.Account,
	*gqlaccounts.Account,
	*gqlaccounts.AccountPayload,
	*gqlaccounts.AccountCreateInput,
	*gqlaccounts.AccountUpdateInput,
	*gqlaccounts.AccountListFilter,
	*gqlaccounts.AccountListOrder,
	*gqlaccounts.User,
	*gqlaccounts.UserCreateInput,
	*gqlaccounts.UserUpdateInput,
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
