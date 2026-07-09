package connectors

import (
	accountgraphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"
	authclientgraphql "github.com/geniusrabbit/blaze-api/repository/authclient/delivery/graphql"
	directaccesstokengraphql "github.com/geniusrabbit/blaze-api/repository/directaccesstoken/delivery/graphql"
	historygraphql "github.com/geniusrabbit/blaze-api/repository/historylog/delivery/graphql"
	optiongraphql "github.com/geniusrabbit/blaze-api/repository/option/delivery/graphql"
	rbacgraphql "github.com/geniusrabbit/blaze-api/repository/rbac/delivery/graphql"
	usergraphql "github.com/geniusrabbit/blaze-api/repository/user/delivery/graphql"
	blazegqlmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"

	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type Page = blazegqlmodels.Page

// AccountConnection implements collection accessor interface with pagination.
type AccountConnection = accountgraphql.AccountConnection[*gqlmodels.Account]

// RBACRoleConnection implements collection accessor interface with pagination.
type RBACRoleConnection = rbacgraphql.RBACRoleConnection

// AuthClientConnection implements collection accessor interface with pagination.
type AuthClientConnection = authclientgraphql.AuthClientConnection

// UserConnection implements collection accessor interface with pagination.
type UserConnection = usergraphql.UserConnection[*gqlmodels.User]

// MemberConnection implements collection accessor interface with pagination.
type MemberConnection = accountgraphql.MemberConnection

// HistoryActionConnection implements collection accessor interface with pagination.
type HistoryActionConnection = historygraphql.HistoryActionConnection

// OptionConnection implements collection accessor interface with pagination.
type OptionConnection = optiongraphql.OptionConnection

// DirectAccessTokenConnection implements collection accessor interface with pagination.
type DirectAccessTokenConnection = directaccesstokengraphql.DirectAccessTokenConnection
