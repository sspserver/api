package wiring

import (
	"github.com/geniusrabbit/blaze-api/repository/user"
	usergraphql "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql"
	userbase "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_base"
	useremail "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_email"
	userpassword "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_password"
	userpassreset "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_password_reset"
	"github.com/sspserver/api/pkg/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type UserQueryResolver interface {
	usergraphql.UserBaseQueryResolver[
		*models.User,
		*gqlmodels.User,
		*gqlmodels.UserCreateInput,
		*gqlmodels.UserUpdateInput,
		*gqlmodels.UserPayload,
		*gqlmodels.UserListFilter,
		*gqlmodels.UserListOrder,
	]
	usergraphql.UserEmailQueryResolver[
		*models.User,
		*gqlmodels.User,
		*gqlmodels.UserPayload,
	]
	// usergraphql.UserUsernameQueryResolver[
	// 	*models.User,
	// 	*gqlmodels.User,
	// 	*gqlmodels.UserCreateInput,
	// 	*gqlmodels.UserPayload,
	// ]
	usergraphql.UserPasswordQueryResolver[
		*models.User,
		*gqlmodels.User,
		*gqlmodels.UserCreateInput,
		*gqlmodels.UserPayload,
	]
	usergraphql.UserPasswordResetQueryResolver[
		*models.User,
		*gqlmodels.User,
		*gqlmodels.UserPayload,
	]
}

type userQueryResolver struct {
	userbase.QueryResolverBase[
		*models.User,
		*gqlmodels.User,
		*gqlmodels.UserCreateInput,
		*gqlmodels.UserUpdateInput,
		*gqlmodels.UserPayload,
		*gqlmodels.UserListFilter,
		*gqlmodels.UserListOrder,
	]
	useremail.QueryResolverEmail[
		*models.User,
		*gqlmodels.User,
		*gqlmodels.UserPayload,
	]
	// userusername.QueryResolverUsername[
	// 	*models.User,
	// 	*gqlmodels.User,
	// 	*gqlmodels.UserCreateInput,
	// 	*gqlmodels.UserPayload,
	// 	*gqlmodels.UserListFilter,
	// 	*gqlmodels.UserListOrder,
	// ]
	userpassword.QueryResolverPassword[
		*models.User,
		*gqlmodels.User,
		*gqlmodels.UserCreateInput,
		*gqlmodels.UserPayload,
		*gqlmodels.UserListFilter,
		*gqlmodels.UserListOrder,
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
			*gqlmodels.User,
			*gqlmodels.UserCreateInput,
			*gqlmodels.UserUpdateInput,
			*gqlmodels.UserPayload,
			*gqlmodels.UserListFilter,
			*gqlmodels.UserListOrder,
		]{
			Core:   core,
			Mapper: mapper,
		}),
		QueryResolverEmail: *useremail.NewQueryResolverEmail(useremail.QueryResolverEmailConfig[
			*models.User,
			*gqlmodels.User,
			*gqlmodels.UserPayload,
		]{
			Core:       core,
			Email:      emailUsecase,
			ToGraphQL:  mapper.ToGQL,
			NewPayload: mapper.NewPayload,
		}),
		// QueryResolverUsername: *userusername.NewQueryResolverUsername(userusername.QueryResolverUsernameConfig[
		// 	*models.User,
		// 	*gqlmodels.User,
		// 	*gqlmodels.UserCreateInput,
		// 	*gqlmodels.UserPayload,
		// 	*gqlmodels.UserListFilter,
		// 	*gqlmodels.UserListOrder,
		// ]{
		// 	Core:      core,
		// 	ToGraphQL: mapper.ToGQL,
		// }),
		QueryResolverPassword: *userpassword.NewQueryResolverPassword(userpassword.QueryResolverPasswordConfig[
			*models.User,
			*gqlmodels.User,
			*gqlmodels.UserCreateInput,
			*gqlmodels.UserPayload,
			*gqlmodels.UserListFilter,
			*gqlmodels.UserListOrder,
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
