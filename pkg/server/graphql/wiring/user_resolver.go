package wiring

import (
	"github.com/geniusrabbit/blaze-api/repository/user"
	usergraphql "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql"
	userbase "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_base"
	useremail "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_email"
	userpassword "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_password"
	userpassreset "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_password_reset"

	"github.com/sspserver/api/pkg/models"
	gqlaccounts "github.com/sspserver/api/pkg/server/graphql/accounts"
)

type UserQueryResolver interface {
	usergraphql.UserBaseQueryResolver[
		*models.User,
		*gqlaccounts.User,
		*gqlaccounts.UserCreateInput,
		*gqlaccounts.UserUpdateInput,
		*gqlaccounts.UserPayload,
		*gqlaccounts.UserListFilter,
		*gqlaccounts.UserListOrder,
	]
	usergraphql.UserEmailQueryResolver[
		*models.User,
		*gqlaccounts.User,
		*gqlaccounts.UserPayload,
	]
	// usergraphql.UserUsernameQueryResolver[
	// 	*models.User,
	// 	*gqlaccounts.User,
	// 	*gqlaccounts.UserCreateInput,
	// 	*gqlaccounts.UserPayload,
	// ]
	usergraphql.UserPasswordQueryResolver[
		*models.User,
		*gqlaccounts.User,
		*gqlaccounts.UserCreateInput,
		*gqlaccounts.UserPayload,
	]
	usergraphql.UserPasswordResetQueryResolver[
		*models.User,
		*gqlaccounts.User,
		*gqlaccounts.UserPayload,
	]
}

type userQueryResolver struct {
	userbase.QueryResolverBase[
		*models.User,
		*gqlaccounts.User,
		*gqlaccounts.UserCreateInput,
		*gqlaccounts.UserUpdateInput,
		*gqlaccounts.UserPayload,
		*gqlaccounts.UserListFilter,
		*gqlaccounts.UserListOrder,
	]
	useremail.QueryResolverEmail[
		*models.User,
		*gqlaccounts.User,
		*gqlaccounts.UserPayload,
	]
	// userusername.QueryResolverUsername[
	// 	*models.User,
	// 	*gqlaccounts.User,
	// 	*gqlaccounts.UserCreateInput,
	// 	*gqlaccounts.UserPayload,
	// 	*gqlaccounts.UserListFilter,
	// 	*gqlaccounts.UserListOrder,
	// ]
	userpassword.QueryResolverPassword[
		*models.User,
		*gqlaccounts.User,
		*gqlaccounts.UserCreateInput,
		*gqlaccounts.UserPayload,
		*gqlaccounts.UserListFilter,
		*gqlaccounts.UserListOrder,
	]
	userpassreset.PasswordResetQueryResolver[*models.User]
}

func NewUserQueryResolver(
	core user.Usecase[*models.User],
	emailUsecase user.EmailUsecase[*models.User],
	passwordUsecase user.PasswordUsecase[*models.User],
) UserQueryResolver {
	mapper := &UserGraphQLMappersImpl{}
	return &userQueryResolver{
		QueryResolverBase: *userbase.NewQueryResolverBase(userbase.QueryResolverBaseConfig[
			*models.User,
			*gqlaccounts.User,
			*gqlaccounts.UserCreateInput,
			*gqlaccounts.UserUpdateInput,
			*gqlaccounts.UserPayload,
			*gqlaccounts.UserListFilter,
			*gqlaccounts.UserListOrder,
		]{
			Core:   core,
			Mapper: mapper,
		}),
		QueryResolverEmail: *useremail.NewQueryResolverEmail(useremail.QueryResolverEmailConfig[
			*models.User,
			*gqlaccounts.User,
			*gqlaccounts.UserPayload,
		]{
			Core:       core,
			Email:      emailUsecase,
			ToGraphQL:  mapper.ToGQL,
			NewPayload: mapper.NewPayload,
		}),
		// QueryResolverUsername: *userusername.NewQueryResolverUsername(userusername.QueryResolverUsernameConfig[
		// 	*models.User,
		// 	*gqlaccounts.User,
		// 	*gqlaccounts.UserCreateInput,
		// 	*gqlaccounts.UserPayload,
		// 	*gqlaccounts.UserListFilter,
		// 	*gqlaccounts.UserListOrder,
		// ]{
		// 	Core:      core,
		// 	ToGraphQL: mapper.ToGQL,
		// }),
		QueryResolverPassword: *userpassword.NewQueryResolverPassword(userpassword.QueryResolverPasswordConfig[
			*models.User,
			*gqlaccounts.User,
			*gqlaccounts.UserCreateInput,
			*gqlaccounts.UserPayload,
			*gqlaccounts.UserListFilter,
			*gqlaccounts.UserListOrder,
		]{
			Core:      core,
			ToGraphQL: mapper.ToGQL,
		}),
		PasswordResetQueryResolver: *userpassreset.NewPasswordResetQueryResolver(userpassreset.PasswordResetQueryResolverConfig[*models.User]{
			Email:    emailUsecase,
			Password: passwordUsecase,
		}),
	}
}
