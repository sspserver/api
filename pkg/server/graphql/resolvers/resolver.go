package resolvers

import (
	"context"
	"errors"

	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/pkg/auth/jwt"
	"github.com/geniusrabbit/blaze-api/pkg/context/session"
	pkgmodels "github.com/geniusrabbit/blaze-api/pkg/models"
	"github.com/geniusrabbit/blaze-api/pkg/requestid"
	accountrepoapi "github.com/geniusrabbit/blaze-api/repository/account"
	account_graphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"
	accountlogin "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql/account_login"
	accountrepo "github.com/geniusrabbit/blaze-api/repository/account/repository"
	accountusecase "github.com/geniusrabbit/blaze-api/repository/account/usecase"
	authclient_graphql "github.com/geniusrabbit/blaze-api/repository/authclient/delivery/graphql"
	authclientrepo "github.com/geniusrabbit/blaze-api/repository/authclient/repository"
	authclientusecase "github.com/geniusrabbit/blaze-api/repository/authclient/usecase"
	directaccesstoken_graphql "github.com/geniusrabbit/blaze-api/repository/directaccesstoken/delivery/graphql"
	datokenrepo "github.com/geniusrabbit/blaze-api/repository/directaccesstoken/repository"
	datokenusecase "github.com/geniusrabbit/blaze-api/repository/directaccesstoken/usecase"
	historylog_graphql "github.com/geniusrabbit/blaze-api/repository/historylog/delivery/graphql"
	historylogrepo "github.com/geniusrabbit/blaze-api/repository/historylog/repository"
	historylogusecase "github.com/geniusrabbit/blaze-api/repository/historylog/usecase"
	"github.com/geniusrabbit/blaze-api/repository/option"
	option_graphql "github.com/geniusrabbit/blaze-api/repository/option/delivery/graphql"
	rbac_graphql "github.com/geniusrabbit/blaze-api/repository/rbac/delivery/graphql"
	rbacrepo "github.com/geniusrabbit/blaze-api/repository/rbac/repository"
	rbacusecase "github.com/geniusrabbit/blaze-api/repository/rbac/usecase"
	socialaccount_graphql "github.com/geniusrabbit/blaze-api/repository/socialaccount/delivery/graphql"
	socaccrepo "github.com/geniusrabbit/blaze-api/repository/socialaccount/repository"
	socaccusecase "github.com/geniusrabbit/blaze-api/repository/socialaccount/usecase"
	userrepoapi "github.com/geniusrabbit/blaze-api/repository/user"
	userbase "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_base"
	useremail "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_email"
	userpassword "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_password"
	userpassreset "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql/user_password_reset"
	userrepo "github.com/geniusrabbit/blaze-api/repository/user/repository"
	userusecase "github.com/geniusrabbit/blaze-api/repository/user/usecase"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	basemodels "github.com/geniusrabbit/blaze-api/server/graphql/models"

	domainmodels "github.com/sspserver/api/pkg/models"
	adformat_graphql "github.com/sspserver/api/pkg/repository/adformat/delivery/graphql"
	"github.com/sspserver/api/pkg/repository/agreement"
	agreement_graphql "github.com/sspserver/api/pkg/repository/agreement/delivery/graphql"
	application_graphql "github.com/sspserver/api/pkg/repository/application/delivery/graphql"
	browser_graphql "github.com/sspserver/api/pkg/repository/browser/delivery/graphql"
	category_graphql "github.com/sspserver/api/pkg/repository/category/delivery/graphql"
	devicemaker_graphql "github.com/sspserver/api/pkg/repository/devicemaker/delivery/graphql"
	devicemodel_graphql "github.com/sspserver/api/pkg/repository/devicemodel/delivery/graphql"
	devicetype_graphql "github.com/sspserver/api/pkg/repository/devicetype/delivery/graphql"
	geo_graphql "github.com/sspserver/api/pkg/repository/geo/delivery/graphql"
	languages_graphql "github.com/sspserver/api/pkg/repository/languages/delivery/graphql"
	os_graphql "github.com/sspserver/api/pkg/repository/os/delivery/graphql"
	"github.com/sspserver/api/pkg/repository/rtbsource"
	rtbsource_graphql "github.com/sspserver/api/pkg/repository/rtbsource/delivery/graphql"
	"github.com/sspserver/api/pkg/repository/statistic"
	statistic_graphql "github.com/sspserver/api/pkg/repository/statistic/delivery/graphql"
	trafficrouter_graphql "github.com/sspserver/api/pkg/repository/trafficrouter/delivery/graphql"
	zone_graphql "github.com/sspserver/api/pkg/repository/zone/delivery/graphql"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
	"github.com/sspserver/api/private/agreements"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	general *generalResolver
	// Basic resolvers
	users             userQueryHandler
	accAuth           accountAuthHandler
	accounts          accountQueryHandler
	members           memberQueryHandler
	socAccounts       *socialaccount_graphql.QueryResolver
	roles             *rbac_graphql.QueryResolver
	authclients       *authclient_graphql.QueryResolver
	historylogs       *historylog_graphql.QueryResolver
	options           *option_graphql.QueryResolver
	directaccesstoken *directaccesstoken_graphql.QueryResolver
	// Current API extensions
	rtbsource     *rtbsource_graphql.QueryResolver
	adformat      *adformat_graphql.QueryResolver
	geo           *geo_graphql.QueryResolver
	langs         *languages_graphql.QueryResolver
	categories    *category_graphql.QueryResolver
	os            *os_graphql.QueryResolver
	browsers      *browser_graphql.QueryResolver
	device_types  *devicetype_graphql.QueryResolver
	device_models *devicemodel_graphql.QueryResolver
	device_makers *devicemaker_graphql.QueryResolver
	app           *application_graphql.QueryResolver
	zone          *zone_graphql.QueryResolver
	statistic     *statistic_graphql.QueryResolver
	trafficrouter *trafficrouter_graphql.Resolver
	agreements    *agreement_graphql.QueryResolver
}

type Usecases struct {
	Stats     statistic.Usecase
	RTBSource rtbsource.Usecase
	Options   option.Usecase
}

type userQueryHandler interface {
	CreateUser(ctx context.Context, input *gqlmodels.UserInput) (*gqlmodels.UserPayload, error)
	UpdateUser(ctx context.Context, id uint64, input *gqlmodels.UserInput) (*gqlmodels.UserPayload, error)
	ApproveUser(ctx context.Context, id uint64, msg *string) (*gqlmodels.UserPayload, error)
	RejectUser(ctx context.Context, id uint64, msg *string) (*gqlmodels.UserPayload, error)
	ResetUserPassword(ctx context.Context, email string) (*basemodels.StatusResponse, error)
	UpdateResetedUserPassword(ctx context.Context, token, email, password string) (*basemodels.StatusResponse, error)
	CurrentUser(ctx context.Context) (*gqlmodels.UserPayload, error)
	User(ctx context.Context, id uint64, email string) (*gqlmodels.UserPayload, error)
	ListUsers(ctx context.Context, filter *gqlmodels.UserListFilter, order []*gqlmodels.UserListOrder, page *basemodels.Page) (*connectors.CollectionConnection[*gqlmodels.User], error)
}

type accountAuthHandler interface {
	Login(ctx context.Context, login, password string, accountID ...uint64) (*basemodels.SessionToken, error)
	Logout(ctx context.Context) (bool, error)
	SwitchAccount(ctx context.Context, id uint64) (*basemodels.SessionToken, error)
	CurrentSession(ctx context.Context) (*basemodels.SessionToken, error)
	ListRolesAndPermissions(ctx context.Context, accountID uint64, order []*basemodels.RBACRoleListOrder) (*rbac_graphql.RBACRoleConnection, error)
}

type accountQueryHandler interface {
	CurrentAccount(ctx context.Context) (*gqlmodels.AccountPayload, error)
	Account(ctx context.Context, id uint64) (*gqlmodels.AccountPayload, error)
	RegisterAccount(ctx context.Context, input *gqlmodels.AccountCreateInput) (*gqlmodels.AccountCreatePayload, error)
	UpdateAccount(ctx context.Context, id uint64, input *gqlmodels.AccountInput) (*gqlmodels.AccountPayload, error)
	ApproveAccount(ctx context.Context, id uint64, msg string) (*gqlmodels.AccountPayload, error)
	RejectAccount(ctx context.Context, id uint64, msg string) (*gqlmodels.AccountPayload, error)
	ListAccounts(ctx context.Context, filter *gqlmodels.AccountListFilter, order []*gqlmodels.AccountListOrder, page *basemodels.Page) (*connectors.CollectionConnection[*gqlmodels.Account], error)
}

type memberQueryHandler interface {
	Invite(ctx context.Context, accountID uint64, member basemodels.InviteMemberInput) (*basemodels.MemberPayload, error)
	Update(ctx context.Context, memberID uint64, member basemodels.MemberInput) (*basemodels.MemberPayload, error)
	Remove(ctx context.Context, memberID uint64) (*basemodels.MemberPayload, error)
	Approve(ctx context.Context, memberID uint64, msg string) (*basemodels.MemberPayload, error)
	Reject(ctx context.Context, memberID uint64, msg string) (*basemodels.MemberPayload, error)
	List(ctx context.Context, filter *basemodels.MemberListFilter, order []*basemodels.MemberListOrder, page *basemodels.Page) (*connectors.CollectionConnection[*basemodels.Member], error)
}

type userRepoWithEmail struct {
	userrepoapi.Repository[*domainmodels.User]
	userrepoapi.EmailRepository[*domainmodels.User]
}

type userMapper struct{}

func (userMapper) New() *domainmodels.User {
	return &domainmodels.User{}
}

func (userMapper) ToGQL(u *domainmodels.User) *gqlmodels.User {
	if u == nil {
		return nil
	}
	return &gqlmodels.User{
		ID:        u.GetID(),
		Username:  u.GetEmail(),
		Status:    basemodels.ApproveStatusFrom(u.GetApprove()),
		CreatedAt: u.GetCreatedAt(),
		UpdatedAt: u.GetUpdatedAt(),
	}
}

func (userMapper) FromCreateInput(inp *gqlmodels.UserInput) *domainmodels.User {
	usr := &domainmodels.User{}
	if inp == nil {
		return usr
	}
	if inp.Username != nil {
		usr.SetEmail(*inp.Username)
	}
	if inp.Status != nil {
		usr.SetApprove(inp.Status.ModelStatus())
	}
	return usr
}

func (userMapper) FromUpdateInput(inp *gqlmodels.UserInput, dest *domainmodels.User) *domainmodels.User {
	if dest == nil {
		dest = &domainmodels.User{}
	}
	if inp == nil {
		return dest
	}
	if inp.Username != nil {
		dest.SetEmail(*inp.Username)
	}
	if inp.Status != nil {
		dest.SetApprove(inp.Status.ModelStatus())
	}
	return dest
}

func (m userMapper) NewPayload(clientMutationID string, userID uint64, u *gqlmodels.User) *gqlmodels.UserPayload {
	return &gqlmodels.UserPayload{
		ClientMutationID: clientMutationID,
		UserID:           userID,
		User:             u,
	}
}

func (userMapper) FromFilter(f *gqlmodels.UserListFilter) userrepoapi.QOption {
	if f == nil {
		return nil
	}
	return &userrepoapi.ListFilter{
		FilterBase: userrepoapi.FilterBase{ID: f.ID},
		FilterEmail: userrepoapi.FilterEmail{
			Emails: f.Emails,
		},
	}
}

func (userMapper) FromOrder(o *gqlmodels.UserListOrder) userrepoapi.QOption {
	if o == nil {
		return nil
	}
	return &userrepoapi.ListOrder{
		OrderBase: userrepoapi.OrderBase{
			ID:        asOrder(o.ID),
			Status:    asOrder(o.Status),
			CreatedAt: asOrder(o.CreatedAt),
			UpdatedAt: asOrder(o.UpdatedAt),
		},
		OrderEmail: userrepoapi.OrderEmail{Email: asOrder(o.Email)},
	}
}

type userQueryResolver struct {
	userbase.QueryResolverBase[
		*domainmodels.User,
		*gqlmodels.User,
		*gqlmodels.UserInput,
		*gqlmodels.UserInput,
		*gqlmodels.UserPayload,
		*gqlmodels.UserListFilter,
		*gqlmodels.UserListOrder,
	]
	useremail.QueryResolverEmail[
		*domainmodels.User,
		*gqlmodels.User,
		*gqlmodels.UserPayload,
	]
	userpassword.QueryResolverPassword[
		*domainmodels.User,
		*gqlmodels.User,
		*gqlmodels.UserInput,
		*gqlmodels.UserPayload,
		*gqlmodels.UserListFilter,
		*gqlmodels.UserListOrder,
	]
	userpassreset.PasswordResetQueryResolver[*domainmodels.User]
}

type accountAuthResolver struct {
	*account_graphql.AuthResolver[*domainmodels.User, *domainmodels.Account]
	*accountlogin.Resolver[*domainmodels.User, *domainmodels.Account]
}

type accountQueryResolver struct {
	accounts         accountrepoapi.Usecase[*domainmodels.User, *domainmodels.Account]
	users            userrepoapi.Usecase[*domainmodels.User]
	userPasswordRepo userrepoapi.PasswordRepository[*domainmodels.User]
}

func (r *accountQueryResolver) CurrentAccount(ctx context.Context) (*gqlmodels.AccountPayload, error) {
	acc, _ := session.Account(ctx).(*domainmodels.Account)
	if acc == nil {
		acc = &domainmodels.Account{}
	}
	return &gqlmodels.AccountPayload{
		ClientMutationID: requestid.Get(ctx),
		AccountID:        acc.GetID(),
		Account:          toGQLAccount(acc),
	}, nil
}

func (r *accountQueryResolver) Account(ctx context.Context, id uint64) (*gqlmodels.AccountPayload, error) {
	acc, err := r.accounts.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &gqlmodels.AccountPayload{
		ClientMutationID: requestid.Get(ctx),
		AccountID:        id,
		Account:          toGQLAccount(acc),
	}, nil
}

func (r *accountQueryResolver) RegisterAccount(ctx context.Context, input *gqlmodels.AccountCreateInput) (*gqlmodels.AccountCreatePayload, error) {
	if input == nil || input.Account == nil {
		return nil, errors.New("account data is required")
	}

	var (
		err      error
		ownerObj *domainmodels.User
	)

	mapper := userMapper{}
	switch {
	case input.OwnerID != nil && *input.OwnerID > 0:
		ownerObj, err = r.users.Get(ctx, *input.OwnerID)
		if err != nil {
			return nil, err
		}
	case input.Owner != nil:
		ownerObj = mapper.FromCreateInput(input.Owner)
		var uid uint64
		if input.Password != "" {
			uid, err = r.userPasswordRepo.CreateWithPassword(ctx, ownerObj, input.Password)
		} else {
			uid, err = r.users.Create(ctx, ownerObj)
		}
		if err != nil {
			return nil, err
		}
		ownerObj, err = r.users.Get(ctx, uid)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("owner is required")
	}

	accObj := &domainmodels.Account{}
	applyAccountInput(accObj, input.Account)
	if _, err = r.accounts.Register(ctx, ownerObj, accObj); err != nil {
		return nil, err
	}

	return &gqlmodels.AccountCreatePayload{
		ClientMutationID: requestid.Get(ctx),
		Account:          toGQLAccount(accObj),
		Owner:            mapper.ToGQL(ownerObj),
	}, nil
}

func (r *accountQueryResolver) UpdateAccount(ctx context.Context, id uint64, input *gqlmodels.AccountInput) (*gqlmodels.AccountPayload, error) {
	accObj, err := r.accounts.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	applyAccountInput(accObj, input)
	if _, err = r.accounts.Update(ctx, accObj); err != nil {
		return nil, err
	}
	accObj, err = r.accounts.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &gqlmodels.AccountPayload{
		ClientMutationID: requestid.Get(ctx),
		AccountID:        id,
		Account:          toGQLAccount(accObj),
	}, nil
}

func (r *accountQueryResolver) ApproveAccount(ctx context.Context, id uint64, msg string) (*gqlmodels.AccountPayload, error) {
	_ = msg
	accObj, err := r.accounts.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	accObj.SetApprove(pkgmodels.ApprovedApproveStatus)
	if _, err = r.accounts.Update(ctx, accObj); err != nil {
		return nil, err
	}
	return &gqlmodels.AccountPayload{ClientMutationID: requestid.Get(ctx), AccountID: id, Account: toGQLAccount(accObj)}, nil
}

func (r *accountQueryResolver) RejectAccount(ctx context.Context, id uint64, msg string) (*gqlmodels.AccountPayload, error) {
	_ = msg
	accObj, err := r.accounts.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	accObj.SetApprove(pkgmodels.DisapprovedApproveStatus)
	if _, err = r.accounts.Update(ctx, accObj); err != nil {
		return nil, err
	}
	return &gqlmodels.AccountPayload{ClientMutationID: requestid.Get(ctx), AccountID: id, Account: toGQLAccount(accObj)}, nil
}

func (r *accountQueryResolver) ListAccounts(ctx context.Context, filter *gqlmodels.AccountListFilter, order []*gqlmodels.AccountListOrder, page *basemodels.Page) (*connectors.CollectionConnection[*gqlmodels.Account], error) {
	if page == nil {
		page = &basemodels.Page{}
	}
	filterOpt := toAccountFilterOption(filter)
	orderOpts := toAccountOrderOptions(order)

	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Account]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Account, error) {
			opts := append([]accountrepoapi.QOption{}, orderOpts...)
			if filterOpt != nil {
				opts = append(opts, filterOpt)
			}
			opts = append(opts, page.Pagination())
			list, err := r.accounts.FetchList(ctx, opts...)
			if err != nil {
				return nil, err
			}
			return xtypes.SliceApply(list, toGQLAccount), nil
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			if filterOpt == nil {
				return r.accounts.Count(ctx)
			}
			return r.accounts.Count(ctx, filterOpt)
		},
	}, page), nil
}

type memberQueryResolver struct {
	core *account_graphql.MemberQueryResolver[*domainmodels.User, *domainmodels.Account, *gqlmodels.Account]
}

func (r *memberQueryResolver) Invite(ctx context.Context, accountID uint64, member basemodels.InviteMemberInput) (*basemodels.MemberPayload, error) {
	return r.core.Invite(ctx, accountID, &member)
}

func (r *memberQueryResolver) Update(ctx context.Context, memberID uint64, member basemodels.MemberInput) (*basemodels.MemberPayload, error) {
	return r.core.Update(ctx, memberID, &member)
}

func (r *memberQueryResolver) Remove(ctx context.Context, memberID uint64) (*basemodels.MemberPayload, error) {
	return r.core.Remove(ctx, memberID)
}

func (r *memberQueryResolver) Approve(ctx context.Context, memberID uint64, msg string) (*basemodels.MemberPayload, error) {
	return r.core.Approve(ctx, memberID, msg)
}

func (r *memberQueryResolver) Reject(ctx context.Context, memberID uint64, msg string) (*basemodels.MemberPayload, error) {
	return r.core.Reject(ctx, memberID, msg)
}

func (r *memberQueryResolver) List(ctx context.Context, filter *basemodels.MemberListFilter, order []*basemodels.MemberListOrder, page *basemodels.Page) (*connectors.CollectionConnection[*basemodels.Member], error) {
	return r.core.List(ctx, filter, order, page)
}

func asOrder(o *basemodels.Ordering) pkgmodels.Order {
	if o == nil {
		return pkgmodels.OrderUndefined
	}
	return pkgmodels.Order(o.AsOrder())
}

func toAccountFilterOption(filter *gqlmodels.AccountListFilter) accountrepoapi.QOption {
	if filter == nil {
		return nil
	}
	return &accountrepoapi.Filter{
		ID:     filter.ID,
		UserID: filter.UserID,
		Title:  filter.Title,
		Status: xtypes.SliceApply(filter.Status, func(v basemodels.ApproveStatus) pkgmodels.ApproveStatus { return v.ModelStatus() }),
	}
}

func toAccountOrderOptions(order []*gqlmodels.AccountListOrder) []accountrepoapi.QOption {
	return xtypes.SliceApply(order, func(o *gqlmodels.AccountListOrder) accountrepoapi.QOption {
		if o == nil {
			return nil
		}
		return &accountrepoapi.ListOrder{
			ID:     asOrder(o.ID),
			Title:  asOrder(o.Title),
			Status: asOrder(o.Status),
		}
	})
}

func toGQLAccount(acc *domainmodels.Account) *gqlmodels.Account {
	if acc == nil {
		return nil
	}
	return &gqlmodels.Account{
		ID:                acc.GetID(),
		Status:            basemodels.ApproveStatusFrom(acc.GetApprove()),
		Title:             acc.Name,
		Description:       acc.Description,
		LogoURI:           "",
		PolicyURI:         "",
		TermsOfServiceURI: "",
		ClientURI:         "",
		Contacts:          nil,
		CreatedAt:         acc.GetCreatedAt(),
		UpdatedAt:         acc.GetUpdatedAt(),
	}
}

func applyAccountInput(dest *domainmodels.Account, input *gqlmodels.AccountInput) {
	if dest == nil || input == nil {
		return
	}
	if input.Status != nil {
		dest.SetApprove(input.Status.ModelStatus())
	}
	if input.Title != nil {
		dest.Name = *input.Title
	}
	if input.Description != nil {
		dest.Description = *input.Description
	}
}

func NewResolver(usecases *Usecases, provider *jwt.Provider) *Resolver {
	newUser := func() *domainmodels.User { return &domainmodels.User{} }
	newAccount := func() *domainmodels.Account { return &domainmodels.Account{} }
	newMember := func() *accountrepoapi.Member[*domainmodels.User, *domainmodels.Account] {
		return &accountrepoapi.Member[*domainmodels.User, *domainmodels.Account]{}
	}

	userRepoInst := userrepo.NewRepository(newUser)
	userEmailRepo := userrepo.NewEmailRepository(userRepoInst, newUser)
	userPasswordRepo := userrepo.NewPasswordRepository(userRepoInst, newUser)
	userRepoWithEmail := &userRepoWithEmail{
		Repository:      userRepoInst,
		EmailRepository: userEmailRepo,
	}

	accountRepoInst := accountrepo.NewSessionRepository(newUser, newAccount, newMember)
	memberRepoInst := accountrepo.NewMemberRepositoryFor(newMember)
	rbacRepoInst := rbacrepo.New()

	userCoreUsecase := userusecase.NewUsecase(userRepoInst)
	userEmailUsecase := userusecase.NewEmailUsecase(userEmailRepo, newUser)
	userPasswordUsecase := userusecase.NewPasswordUsecase(userCoreUsecase, userPasswordRepo)

	accountUsecaseInst := accountusecase.NewAccountUsecase(userRepoWithEmail, accountRepoInst, memberRepoInst)
	memberUsecaseInst := accountusecase.NewMemberUsecase(userRepoWithEmail, accountRepoInst, memberRepoInst)

	userMappers := userMapper{}
	usersResolver := &userQueryResolver{
		QueryResolverBase: *userbase.NewQueryResolverBase(userbase.QueryResolverBaseConfig[
			*domainmodels.User,
			*gqlmodels.User,
			*gqlmodels.UserInput,
			*gqlmodels.UserInput,
			*gqlmodels.UserPayload,
			*gqlmodels.UserListFilter,
			*gqlmodels.UserListOrder,
		]{
			Core:   userCoreUsecase,
			Mapper: userMappers,
		}),
		QueryResolverEmail: *useremail.NewQueryResolverEmail(useremail.QueryResolverEmailConfig[
			*domainmodels.User,
			*gqlmodels.User,
			*gqlmodels.UserPayload,
		]{
			Core:       userCoreUsecase,
			Email:      userEmailUsecase,
			ToGraphQL:  userMappers.ToGQL,
			NewPayload: userMappers.NewPayload,
		}),
		QueryResolverPassword: *userpassword.NewQueryResolverPassword(userpassword.QueryResolverPasswordConfig[
			*domainmodels.User,
			*gqlmodels.User,
			*gqlmodels.UserInput,
			*gqlmodels.UserPayload,
			*gqlmodels.UserListFilter,
			*gqlmodels.UserListOrder,
		]{
			Core:     userCoreUsecase,
			Password: userPasswordUsecase,
			UserFromInput: func(input *gqlmodels.UserInput, _ ...pkgmodels.ApproveStatus) *domainmodels.User {
				return userMappers.FromCreateInput(input)
			},
			NewPayload: userMappers.NewPayload,
			ToGraphQL:  userMappers.ToGQL,
		}),
		PasswordResetQueryResolver: *userpassreset.NewPasswordResetQueryResolver(userpassreset.PasswordResetQueryResolverConfig[*domainmodels.User]{
			Email:    userEmailUsecase,
			Password: userPasswordUsecase,
		}),
	}

	accAuthCore := account_graphql.NewAuthResolver(provider, accountRepoInst, accountUsecaseInst, rbacRepoInst)
	loginResolver := accountlogin.New(
		provider,
		accountlogin.NewEmailPasswordLogin(userEmailRepo, userPasswordRepo),
		accountRepoInst,
	)
	membersResolver := &memberQueryResolver{
		core: account_graphql.NewMemberQueryResolver(account_graphql.MemberQueryResolverConfig[
			*domainmodels.User,
			*domainmodels.Account,
			*gqlmodels.Account,
		]{
			Accounts: accountUsecaseInst,
			Members:  memberUsecaseInst,
			UserRepo: userRepoInst,
		}),
	}

	res := &Resolver{
		users:             usersResolver,
		accAuth:           &accountAuthResolver{AuthResolver: accAuthCore, Resolver: loginResolver},
		accounts:          &accountQueryResolver{accounts: accountUsecaseInst, users: userCoreUsecase, userPasswordRepo: userPasswordRepo},
		members:           membersResolver,
		socAccounts:       socialaccount_graphql.NewQueryResolver(socaccusecase.NewSocaccUsecase(socaccrepo.NewSocaccRepository())),
		roles:             rbac_graphql.NewQueryResolver(rbacusecase.New(rbacRepoInst)),
		authclients:       authclient_graphql.NewQueryResolver(authclientusecase.NewAuthclientUsecase(authclientrepo.NewAuthclientRepository())),
		historylogs:       historylog_graphql.NewQueryResolver(historylogusecase.NewUsecase(historylogrepo.New())),
		options:           option_graphql.NewQueryResolver(usecases.Options),
		directaccesstoken: directaccesstoken_graphql.NewQueryResolver(datokenusecase.New(datokenrepo.NewDirectAccessTokenRepository())),
		// Current API extensions
		rtbsource:     rtbsource_graphql.NewQueryResolver(usecases.RTBSource),
		adformat:      adformat_graphql.NewQueryResolver(),
		geo:           geo_graphql.NewQueryResolver(),
		langs:         languages_graphql.NewQueryResolver(),
		categories:    category_graphql.NewQueryResolver(),
		os:            os_graphql.NewQueryResolver(),
		browsers:      browser_graphql.NewQueryResolver(),
		device_types:  devicetype_graphql.NewQueryResolver(),
		device_models: devicemodel_graphql.NewQueryResolver(),
		device_makers: devicemaker_graphql.NewQueryResolver(),
		app:           application_graphql.NewQueryResolver(),
		zone:          zone_graphql.NewQueryResolver(),
		statistic:     statistic_graphql.NewQueryResolver(usecases.Stats),
		trafficrouter: trafficrouter_graphql.NewResolver(),
		agreements: agreement_graphql.NewQueryResolver(
			agreement.NewRepositoryOptions(usecases.Options, ListAgreements()),
		),
	}
	res.general = &generalResolver{res}
	return res
}

func ListAgreements() []*domainmodels.Agreement {
	return xtypes.SliceApply(agreements.Agreements(),
		func(a *agreements.Agreement) *domainmodels.Agreement {
			return &domainmodels.Agreement{
				Codename:        a.Meta.Codename,
				Title:           a.Meta.Title,
				Description:     a.Meta.Description,
				Version:         a.Meta.Version,
				Type:            a.Meta.Type,
				IssuedBy:        a.Meta.IssuedBy,
				BodyMarkdown:    a.BodyMarkdown,
				BodyHTML:        a.BodyHTML,
				AcceptAccountID: 0,   // Will be filled later
				AcceptByUserID:  0,   // Will be filled later
				Signature:       "",  // Optional, will be filled later
				AcceptedAt:      nil, // Will be filled later
				CreatedAt:       a.Meta.CreatedAt,
			}
		},
	)
}
