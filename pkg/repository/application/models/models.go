package models

import (
	"time"

	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/gosql/v2"
	"gorm.io/gorm"
)

// Application model describes site or mobile/desktop application
type Application struct {
	ID uint64 `json:"id" gorm:"primaryKey"`

	Title       string `json:"title"`
	Description string `json:"description"`

	AccountID uint64 `json:"account_id"`
	CreatorID uint64 `json:"creator_id"`

	// Unique application identifier like:
	//   - site domain -> domain.com
	//   - mobile/desktop application bundle -> com.application.game
	URI      string                `json:"uri"`
	Type     types.ApplicationType `gorm:"type:ApplicationType" json:"type"`
	Platform types.PlatformType    `gorm:"type:PlatformType" json:"platform"`
	Premium  bool                  `json:"premium"`

	// Status of the application
	Status types.ApproveStatus `gorm:"type:ApproveStatus" json:"status"`

	// Is Active application
	Active types.ActiveStatus `gorm:"type:ActiveStatus" json:"active"`

	// Is private campaign type
	Private types.PrivateStatus `gorm:"type:PrivateStatus;not null" json:"private"`

	Categories gosql.NullableNumberArray[uint] `gorm:"type:BIGINT[]" json:"categories,omitempty"`

	// RevenueShare is the percent of the raw income shared with the publisher
	RevenueShare float64 `json:"revenue_share,omitempty"`

	// Advertisement sources
	AllowedSources    gosql.NullableOrderedNumberArray[int64] `gorm:"type:BIGINT[]" json:"allowed_sources,omitempty"`
	DisallowedSources gosql.NullableOrderedNumberArray[int64] `gorm:"type:BIGINT[]" json:"disallowed_sources,omitempty"`

	// Time marks
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// TableName in database
func (app *Application) TableName() string {
	return "adv_application"
}

// RBACResourceName returns the name of the resource for the RBAC
func (app *Application) RBACResourceName() string {
	return "adv_application"
}
