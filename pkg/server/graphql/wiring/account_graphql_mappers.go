package wiring

import (
	pkgModels "github.com/geniusrabbit/blaze-api/pkg/models"
	accountrepo "github.com/geniusrabbit/blaze-api/repository/account"
	accountgraphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"
	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// AccountGraphQLMappers is the 2-param consumer alias for example/api.
// It pins the 4 concrete types (create input, update input, filter, order)
// so callers need only supply T (domain account) and TGQLAccount (GraphQL account type).
type AccountGraphQLMappers[T accountrepo.Model, TGQLAccount any] = accountgraphql.AccountGraphQLMappers[
	T,
	TGQLAccount,
	*gqlmodels.AccountPayload,     // TGQLPayload
	*gqlmodels.AccountCreateInput, // TGQLCreateInput
	*gqlmodels.AccountUpdateInput, // TGQLUpdateInput (status-only update in base schema)
	*gqlmodels.AccountListFilter,  // TFilter
	*gqlmodels.AccountListOrder,   // TOrder
	*gqlmodels.User,               // TGQLUser
]

// AccountGraphQLMappersImpl implements AccountGraphQLMappers for example/api.
// All methods delegate to the free functions in domain/graphql.go.
// Instantiate with AccountGraphQLMappersImpl{} — stateless.
type AccountGraphQLMappersImpl struct{}

// Compile-time assertion: AccountGraphQLMappersImpl satisfies the 2-param alias when T = *Account.
var _ AccountGraphQLMappers[*models.Account, *gqlmodels.Account] = AccountGraphQLMappersImpl{}

// New creates a new empty domain Account.
func (AccountGraphQLMappersImpl) New() *models.Account {
	return new(models.Account)
}

// ToGQL maps a domain Account to the extended GraphQL Account model.
func (AccountGraphQLMappersImpl) ToGQL(a *models.Account) *gqlmodels.Account {
	return AccountToGraphQL(a)
}

// NewPayload builds account payload from parts.
func (AccountGraphQLMappersImpl) NewPayload(clientMutationID string, accountID uint64, account *gqlmodels.Account) *gqlmodels.AccountPayload {
	acc := account
	return &gqlmodels.AccountPayload{
		ClientMutationID: clientMutationID,
		AccountID:        accountID,
		Account:          acc,
	}
}

// FromCreateInput builds a new domain Account from a create-account account input.
func (AccountGraphQLMappersImpl) FromCreateInput(inp *gqlmodels.AccountCreateInput) *models.Account {
	if inp == nil {
		return new(models.Account)
	}
	return FillAccountFromCreateInput(new(models.Account), inp)
}

// FromUpdateInput merges an update mutation input into an existing domain Account.
// AccountUpdateInput carries only the approval status; profile edits use a separate mutation.
func (AccountGraphQLMappersImpl) FromUpdateInput(inp *gqlmodels.AccountUpdateInput, dest *models.Account) *models.Account {
	if inp == nil || dest == nil {
		return dest
	}
	if inp.Status != nil {
		dest.SetApprove(inp.Status.ModelStatus())
	}
	return dest
}

// FromFilter converts the extended GraphQL account list filter to a domain QOption.
func (AccountGraphQLMappersImpl) FromFilter(f *gqlmodels.AccountListFilter) accountrepo.QOption {
	if f == nil {
		return nil
	}
	return &accountrepo.Filter{
		ID:     f.ID,
		UserID: f.UserID,
	}
}

// FromOrder converts the extended GraphQL account list order to a domain QOption.
func (AccountGraphQLMappersImpl) FromOrder(o *gqlmodels.AccountListOrder) accountrepo.QOption {
	if o == nil {
		return nil
	}
	return &accountrepo.ListOrder{
		ID:        o.ID.AsOrder(),
		Status:    o.Status.AsOrder(),
		CreatedAt: o.CreatedAt.AsOrder(),
		UpdatedAt: o.UpdatedAt.AsOrder(),
	}
}

// FillAccountFromInputWithStatus is a helper for wiring account repos that need an approve-status override.
// Used internally by account graphql wiring.
func FillAccountFromInputWithStatus(dest *models.Account, inp *gqlmodels.AccountUpdateInput, appStatus ...pkgModels.ApproveStatus) *models.Account {
	return FillAccountFromUpdateInput(dest, inp, appStatus...)
}
