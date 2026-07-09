package resolvers

import (
	"github.com/demdxx/xtypes"

	"github.com/geniusrabbit/blaze-api/pkg/auth/jwt"
	accountrepoapi "github.com/geniusrabbit/blaze-api/repository/account"
	account_graphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"
	accountgraphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"
	accountlogin "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql/account_login"
	accountrepo "github.com/geniusrabbit/blaze-api/repository/account/repository"
	accountusecase "github.com/geniusrabbit/blaze-api/repository/account/usecase"
	authclient_graphql "github.com/geniusrabbit/blaze-api/repository/authclient/delivery/graphql"
	directaccesstoken_graphql "github.com/geniusrabbit/blaze-api/repository/directaccesstoken/delivery/graphql"
	historylog_graphql "github.com/geniusrabbit/blaze-api/repository/historylog/delivery/graphql"
	"github.com/geniusrabbit/blaze-api/repository/option"
	option_graphql "github.com/geniusrabbit/blaze-api/repository/option/delivery/graphql"
	rbac_graphql "github.com/geniusrabbit/blaze-api/repository/rbac/delivery/graphql"
	rbacrepo "github.com/geniusrabbit/blaze-api/repository/rbac/repository"
	socialaccount_graphql "github.com/geniusrabbit/blaze-api/repository/socialaccount/delivery/graphql"
	userrepo "github.com/geniusrabbit/blaze-api/repository/user/repository"
	userusecase "github.com/geniusrabbit/blaze-api/repository/user/usecase"

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
	"github.com/sspserver/api/pkg/server/graphql/wiring"
	"github.com/sspserver/api/private/agreements"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	general *generalResolver
	// Basic resolvers
	users             wiring.UserQueryResolver
	accAuth           accountgraphql.AuthQueryHandler
	accLogin          accountgraphql.AccountLoginHandler
	loginHandler      wiring.EmailPasswordLoginHandler
	accounts          wiring.AccountQueryHandler
	members           accountgraphql.MemberQueryHandler
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

func NewResolver(usecases *Usecases, provider *jwt.Provider) *Resolver {
	newUser := func() *domainmodels.User { return &domainmodels.User{} }
	newAccount := func() *domainmodels.Account { return &domainmodels.Account{} }
	newMember := func() *accountrepoapi.Member[*domainmodels.User, *domainmodels.Account] {
		return &accountrepoapi.Member[*domainmodels.User, *domainmodels.Account]{}
	}

	userRepoInst := userrepo.NewRepository(newUser)
	userCoreUsecase := userusecase.NewUsecase(userRepoInst)
	userEmailRepo := userrepo.NewEmailRepository(userRepoInst, newUser)
	userEmailUsecase := userusecase.NewEmailUsecase(userEmailRepo, newUser)
	userPasswordRepo := userrepo.NewPasswordRepository(userRepoInst, newUser)
	userPasswordUsecase := userusecase.NewPasswordUsecase(userCoreUsecase, userPasswordRepo)

	accountRepoInst := accountrepo.NewSessionRepository(newUser, newAccount, newMember)
	memberRepoInst := accountrepo.NewMemberRepositoryFor(newMember)
	accountUsecaseInst := accountusecase.NewAccountUsecase(userRepoInst, accountRepoInst, memberRepoInst)
	memberUsecaseInst := accountusecase.NewMemberUsecase(userRepoInst, accountRepoInst, memberRepoInst)

	rbacRepoInst := rbacrepo.New()
	accAuthCore := account_graphql.NewAuthResolver(provider,
		accountRepoInst, accountUsecaseInst, rbacRepoInst)
	loginResolver := accountlogin.New(provider,
		accountlogin.NewEmailPasswordLogin(userEmailRepo, userPasswordRepo),
		accountRepoInst,
	)

	res := &Resolver{
		users: wiring.NewUserQueryResolver(
			userCoreUsecase,
			userEmailUsecase,
			userPasswordUsecase,
		),
		accAuth:  accAuthCore,
		accLogin: loginResolver,
		accounts: wiring.NewAccountQueryResolver(wiring.AccountQueryResolverConfig{
			Accounts:       accountUsecaseInst,
			AccountsMapper: wiring.AccountGraphQLMappersImpl{},
			Users:          userCoreUsecase,
			UsersMapper:    wiring.UserGraphQLMappersImpl{},
			Members:        memberUsecaseInst,
		}),
		members: account_graphql.NewMemberQueryResolver(account_graphql.MemberQueryResolverConfig[
			*domainmodels.User,
			*domainmodels.Account,
			*gqlmodels.Account,
		]{
			Accounts: accountUsecaseInst,
			Members:  memberUsecaseInst,
			UserRepo: userRepoInst,
		}),

		socAccounts:       socialaccount_graphql.NewDefaultQueryResolver(),
		roles:             rbac_graphql.NewDefaultQueryResolver(),
		authclients:       authclient_graphql.NewDefaultQueryResolver(),
		historylogs:       historylog_graphql.NewDefaultQueryResolver(),
		options:           option_graphql.NewQueryResolver(usecases.Options),
		directaccesstoken: directaccesstoken_graphql.NewDefaultQueryResolver(),

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
