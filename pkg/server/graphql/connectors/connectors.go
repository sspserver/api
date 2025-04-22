package connectors

import (
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"
	blazegqlmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"
)

type Page = blazegqlmodels.Page

// AccountConnection implements collection accessor interface with pagination.
type AccountConnection = connectors.AccountConnection

// RBACRoleConnection implements collection accessor interface with pagination.
type RBACRoleConnection = connectors.RBACRoleConnection

// AuthClientConnection implements collection accessor interface with pagination.
type AuthClientConnection = connectors.AuthClientConnection

// UserConnection implements collection accessor interface with pagination.
type UserConnection = connectors.UserConnection

// HistoryActionConnection implements collection accessor interface with pagination.
type HistoryActionConnection = connectors.HistoryActionConnection

// OptionConnection implements collection accessor interface with pagination.
type OptionConnection = connectors.OptionConnection
