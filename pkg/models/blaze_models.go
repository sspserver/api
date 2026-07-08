package models

import (
	coremodels "github.com/geniusrabbit/blaze-api/pkg/models"
	accountmodels "github.com/geniusrabbit/blaze-api/repository/account/models"
	authclientmodels "github.com/geniusrabbit/blaze-api/repository/authclient/models"
	datmodels "github.com/geniusrabbit/blaze-api/repository/directaccesstoken/models"
	historymodels "github.com/geniusrabbit/blaze-api/repository/historylog/models"
	optionmodels "github.com/geniusrabbit/blaze-api/repository/option/models"
	rbacmodels "github.com/geniusrabbit/blaze-api/repository/rbac/models"
	socialmodels "github.com/geniusrabbit/blaze-api/repository/socialaccount/models"
	usermodels "github.com/geniusrabbit/blaze-api/repository/user/models"
)

// API basic types
type (
	M2MAccountMemberRole = accountmodels.M2MAccountMemberRole
	AccountSocial        = socialmodels.AccountSocial
	AccountSocialSession = socialmodels.AccountSocialSession
	Role                 = rbacmodels.Role
	M2MRole              = rbacmodels.M2MRole
	AuthClient           = authclientmodels.AuthClient
	AuthSession          = authclientmodels.AuthSession
	HistoryAction        = historymodels.HistoryAction
	Option               = optionmodels.Option
	DirectAccessToken    = datmodels.DirectAccessToken
	UserPasswordReset    = usermodels.UserPasswordReset
)

// OptionType type casting
type OptionType = optionmodels.OptionType

const (
	UndefinedOptionType = optionmodels.UndefinedOptionType
	UserOptionType      = optionmodels.UserOptionType
	AccountOptionType   = optionmodels.AccountOptionType
	SystemOptionType    = optionmodels.SystemOptionType
)

// Order type casting
type Order = coremodels.Order

const (
	OrderUndefined = coremodels.OrderUndefined
	OrderAsc       = coremodels.OrderAsc
	OrderDesc      = coremodels.OrderDesc
)

// PrepareQuery returns the query with applied order
func OrderFromStr(s string) Order { return coremodels.OrderFromStr(s) }

// AvailableStatus type
type AvailableStatus = coremodels.AvailableStatus

// AvailableStatus option constants...
const (
	UndefinedAvailableStatus   = coremodels.UndefinedAvailableStatus
	AvailableAvailableStatus   = coremodels.AvailableAvailableStatus
	UnavailableAvailableStatus = coremodels.UnavailableAvailableStatus
)

// BlazeApproveStatus of the model
type BlazeApproveStatus = coremodels.ApproveStatus

// ApproveStatus option constants...
const (
	UndefinedApproveStatus   = coremodels.UndefinedApproveStatus
	ApprovedApproveStatus    = coremodels.ApprovedApproveStatus
	DisapprovedApproveStatus = coremodels.DisapprovedApproveStatus
	BannedApproveStatus      = coremodels.BannedApproveStatus
)
