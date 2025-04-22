package models

import (
	blzmodel "github.com/geniusrabbit/blaze-api/model"
)

// API basic types
type (
	Account              = blzmodel.Account
	M2MAccountMemberRole = blzmodel.M2MAccountMemberRole
	AccountMember        = blzmodel.AccountMember
	User                 = blzmodel.User
	AccountSocial        = blzmodel.AccountSocial
	AccountSocialSession = blzmodel.AccountSocialSession
	Role                 = blzmodel.Role
	M2MRole              = blzmodel.M2MRole
	AuthClient           = blzmodel.AuthClient
	AuthSession          = blzmodel.AuthSession
	HistoryAction        = blzmodel.HistoryAction
	Option               = blzmodel.Option
	DirectAccessToken    = blzmodel.DirectAccessToken
	UserPasswordReset    = blzmodel.UserPasswordReset
)

// OptionType type casting
type OptionType = blzmodel.OptionType

const (
	UndefinedOptionType = blzmodel.UndefinedOptionType
	UserOptionType      = blzmodel.UserOptionType
	AccountOptionType   = blzmodel.AccountOptionType
	SystemOptionType    = blzmodel.SystemOptionType
)

// Order type casting
type Order = blzmodel.Order

const (
	OrderUndefined = blzmodel.OrderUndefined
	OrderAsc       = blzmodel.OrderAsc
	OrderDesc      = blzmodel.OrderDesc
)

// PrepareQuery returns the query with applied order
func OrderFromStr(s string) Order { return blzmodel.OrderFromStr(s) }

// AvailableStatus type
type AvailableStatus = blzmodel.AvailableStatus

// AvailableStatus option constants...
const (
	UndefinedAvailableStatus   = blzmodel.UndefinedAvailableStatus
	AvailableAvailableStatus   = blzmodel.AvailableAvailableStatus
	UnavailableAvailableStatus = blzmodel.UnavailableAvailableStatus
)

// BlazeApproveStatus of the model
type BlazeApproveStatus = blzmodel.ApproveStatus

// ApproveStatus option constants...
const (
	UndefinedApproveStatus   = blzmodel.UndefinedApproveStatus
	ApprovedApproveStatus    = blzmodel.ApprovedApproveStatus
	DisapprovedApproveStatus = blzmodel.DisapprovedApproveStatus
	BannedApproveStatus      = blzmodel.BannedApproveStatus
)
