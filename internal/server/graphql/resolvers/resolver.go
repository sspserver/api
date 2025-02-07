package resolvers

import (
	"github.com/geniusrabbit/blaze-api/pkg/auth/jwt"
	account_graphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"
	authclient_graphql "github.com/geniusrabbit/blaze-api/repository/authclient/delivery/graphql"
	directaccesstoken_graphql "github.com/geniusrabbit/blaze-api/repository/directaccesstoken/delivery/graphql"
	historylog_graphql "github.com/geniusrabbit/blaze-api/repository/historylog/delivery/graphql"
	option_graphql "github.com/geniusrabbit/blaze-api/repository/option/delivery/graphql"
	rbac_graphql "github.com/geniusrabbit/blaze-api/repository/rbac/delivery/graphql"
	rbac "github.com/geniusrabbit/blaze-api/repository/rbac/repository"
	socialaccount_graphql "github.com/geniusrabbit/blaze-api/repository/socialaccount/delivery/graphql"
	user_graphql "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql"

	adformat_graphql "github.com/sspserver/api/internal/repository/adformat/delivery/graphql"
	application_graphql "github.com/sspserver/api/internal/repository/application/delivery/graphql"
	browser_graphql "github.com/sspserver/api/internal/repository/browser/delivery/graphql"
	category_graphql "github.com/sspserver/api/internal/repository/category/delivery/graphql"
	devicemaker_graphql "github.com/sspserver/api/internal/repository/devicemaker/delivery/graphql"
	devicemodel_graphql "github.com/sspserver/api/internal/repository/devicemodel/delivery/graphql"
	devicetype_graphql "github.com/sspserver/api/internal/repository/devicetype/delivery/graphql"
	geo_graphql "github.com/sspserver/api/internal/repository/geo/delivery/graphql"
	os_graphql "github.com/sspserver/api/internal/repository/os/delivery/graphql"
	"github.com/sspserver/api/internal/repository/rtbsource"
	rtbsource_graphql "github.com/sspserver/api/internal/repository/rtbsource/delivery/graphql"
	"github.com/sspserver/api/internal/repository/statistic"
	statistic_graphql "github.com/sspserver/api/internal/repository/statistic/delivery/graphql"
	zone_graphql "github.com/sspserver/api/internal/repository/zone/delivery/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users             *user_graphql.QueryResolver
	accAuth           *account_graphql.AuthResolver
	accounts          *account_graphql.QueryResolver
	members           *account_graphql.MemberQueryResolver
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
	catergories   *category_graphql.QueryResolver
	os            *os_graphql.QueryResolver
	browsers      *browser_graphql.QueryResolver
	device_types  *devicetype_graphql.QueryResolver
	device_models *devicemodel_graphql.QueryResolver
	device_makers *devicemaker_graphql.QueryResolver
	app           *application_graphql.QueryResolver
	zone          *zone_graphql.QueryResolver
	statistic     *statistic_graphql.QueryResolver
}

type Usecases struct {
	Stats     statistic.Usecase
	RTBSource rtbsource.Usecase
}

func NewResolver(usecases *Usecases, provider *jwt.Provider) *Resolver {
	return &Resolver{
		users:             user_graphql.NewQueryResolver(),
		accAuth:           account_graphql.NewAuthResolver(provider, rbac.New()),
		accounts:          account_graphql.NewQueryResolver(),
		members:           account_graphql.NewMemberQueryResolver(),
		socAccounts:       socialaccount_graphql.NewQueryResolver(),
		roles:             rbac_graphql.NewQueryResolver(),
		authclients:       authclient_graphql.NewQueryResolver(),
		historylogs:       historylog_graphql.NewQueryResolver(),
		options:           option_graphql.NewQueryResolver(),
		directaccesstoken: directaccesstoken_graphql.NewQueryResolver(),
		// Current API extensions
		rtbsource:     rtbsource_graphql.NewQueryResolver(usecases.RTBSource),
		adformat:      adformat_graphql.NewQueryResolver(),
		geo:           geo_graphql.NewQueryResolver(),
		catergories:   category_graphql.NewQueryResolver(),
		os:            os_graphql.NewQueryResolver(),
		browsers:      browser_graphql.NewQueryResolver(),
		device_types:  devicetype_graphql.NewQueryResolver(),
		device_models: devicemodel_graphql.NewQueryResolver(),
		device_makers: devicemaker_graphql.NewQueryResolver(),
		app:           application_graphql.NewQueryResolver(),
		zone:          zone_graphql.NewQueryResolver(),
		statistic:     statistic_graphql.NewQueryResolver(usecases.Stats),
	}
}
