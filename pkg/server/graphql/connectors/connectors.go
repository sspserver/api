package connectors

import (
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	blazegqlmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

type Page = blazegqlmodels.Page

// AccountConnection implements collection accessor interface with pagination.
type AccountConnection = connectors.CollectionConnection[*gqlmodels.Account]

// RBACRoleConnection implements collection accessor interface with pagination.
type RBACRoleConnection = connectors.CollectionConnection[*blazegqlmodels.RBACRole]

// AuthClientConnection implements collection accessor interface with pagination.
type AuthClientConnection = connectors.CollectionConnection[*blazegqlmodels.AuthClient]

// UserConnection implements collection accessor interface with pagination.
type UserConnection = connectors.CollectionConnection[*gqlmodels.User]

// HistoryActionConnection implements collection accessor interface with pagination.
type HistoryActionConnection = connectors.CollectionConnection[*blazegqlmodels.HistoryAction]

// OptionConnection implements collection accessor interface with pagination.
type OptionConnection = connectors.CollectionConnection[*blazegqlmodels.Option]
