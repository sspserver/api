package wiring

import (
	userrepo "github.com/geniusrabbit/blaze-api/repository/user"
	usergraphql "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql"

	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// UserGraphQLMappers is the 2-param consumer alias for example/api.
// It pins the 5 concrete types (create input, update input, payload, filter, order)
// so callers need only supply T (domain user) and TGQLUser (GraphQL user type).
//
//	m := domain.UserGraphQLMappersImpl{}
//	resolver := resolvers.NewUserResolver[*domain.User](uc, m)
type UserGraphQLMappers[T userrepo.Model, TGQLUser any] = usergraphql.UserGraphQLMappers[
	T,
	TGQLUser,
	*gqlmodels.UserCreateInput, // TGQLCreateInput
	*gqlmodels.UserUpdateInput, // TGQLUpdateInput (same as create in base schema)
	*gqlmodels.UserPayload,     // TGQLPayload
	*gqlmodels.UserListFilter,  // TFilter
	*gqlmodels.UserListOrder,   // TOrder
]

// UserGraphQLMappersImpl implements UserGraphQLMappers for example/api.
// All methods delegate to the free functions in domain/graphql.go.
// Instantiate with UserGraphQLMappersImpl{} — stateless.
type UserGraphQLMappersImpl struct{}

// Compile-time assertion: UserGraphQLMappersImpl satisfies the 2-param alias when T = *User.
var _ UserGraphQLMappers[*models.User, *gqlmodels.User] = UserGraphQLMappersImpl{}

// New creates a new empty domain User.
func (UserGraphQLMappersImpl) New() *models.User {
	return new(models.User)
}

// ToGQL maps a domain User to the base GraphQL User model.
func (UserGraphQLMappersImpl) ToGQL(u *models.User) *gqlmodels.User {
	return UserToGraphQL(u)
}

// FromCreateInput builds a new domain User from a create mutation input.
func (UserGraphQLMappersImpl) FromCreateInput(inp *gqlmodels.UserCreateInput) *models.User {
	return UserFromCreateInput(inp)
}

// FromUpdateInput merges an update mutation input into an existing domain User.
func (UserGraphQLMappersImpl) FromUpdateInput(inp *gqlmodels.UserUpdateInput, dest *models.User) *models.User {
	return UserFromUpdateInput(inp, dest)
}

// NewPayload wraps a GQL User in the mutation payload type.
func (UserGraphQLMappersImpl) NewPayload(clientMutationID string, userID uint64, u *gqlmodels.User) *gqlmodels.UserPayload {
	v := u
	return &gqlmodels.UserPayload{
		ClientMutationID: clientMutationID,
		UserID:           userID,
		User:             v,
	}
}

// ToAccountGQL converts a domain User to the account-facing GQL user type.
func (UserGraphQLMappersImpl) ToAccountGQL(u *models.User) *gqlmodels.User {
	return UserToGraphQLPtr(u)
}

// FromFilter converts the extended GraphQL user list filter to a domain QOption.
func (UserGraphQLMappersImpl) FromFilter(f *gqlmodels.UserListFilter) userrepo.QOption {
	return UserListFilterMapper(f)
}

// FromOrder converts the extended GraphQL user list order to a domain QOption.
func (UserGraphQLMappersImpl) FromOrder(o *gqlmodels.UserListOrder) userrepo.QOption {
	return UserListOrderMapper(o)
}
