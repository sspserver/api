package models

import (
	"time"

	bzgqlmodel "github.com/geniusrabbit/blaze-api/server/graphql/models"
	"gorm.io/gorm"
)

type (
	Page           = bzgqlmodel.Page
	SocialAccount  = bzgqlmodel.SocialAccount
	RBACRole       = bzgqlmodel.RBACRole
	AuthClient     = bzgqlmodel.AuthClient
	Option         = bzgqlmodel.Option
	SessionToken   = bzgqlmodel.SessionToken
	StatusResponse = bzgqlmodel.StatusResponse
	ActiveStatus   = bzgqlmodel.ActiveStatus

	HistoryAction           = bzgqlmodel.HistoryAction
	HistoryActionListFilter = bzgqlmodel.HistoryActionListFilter
	HistoryActionListOrder  = bzgqlmodel.HistoryActionListOrder
	HistoryActionPayload    = bzgqlmodel.HistoryActionPayload
)

const (
	ActiveStatusActive = bzgqlmodel.ActiveStatusActive
	ActiveStatusPaused = bzgqlmodel.ActiveStatusPaused
)

const (
	ResponseStatusSuccess = bzgqlmodel.ResponseStatusSuccess
	ResponseStatusError   = bzgqlmodel.ResponseStatusError
)

func DeletedAt(t gorm.DeletedAt) *time.Time {
	return bzgqlmodel.DeletedAt(t)
}
