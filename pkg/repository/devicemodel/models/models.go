package models

import (
	"time"

	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/udetect"
	"gorm.io/gorm"
)

// Device set of types
const (
	DeviceTypeUnknown   = udetect.DeviceTypeUnknown
	DeviceTypeMobile    = udetect.DeviceTypeMobile    // Mobile/Tablet
	DeviceTypePC        = udetect.DeviceTypePC        // Desktop
	DeviceTypeTV        = udetect.DeviceTypeTV        // TV
	DeviceTypePhone     = udetect.DeviceTypePhone     // SmartPhone, SmallScreen
	DeviceTypeTablet    = udetect.DeviceTypeTablet    // Tablet
	DeviceTypeConnected = udetect.DeviceTypeConnected // Console, EReader, Watch
	DeviceTypeSetTopBox = udetect.DeviceTypeSetTopBox // MediaHub
	DeviceTypeWatch     = udetect.DeviceTypeWatch     // SmartWatch
	DeviceTypeGlasses   = udetect.DeviceTypeGlasses   // Glasses
	DeviceTypeOOH       = udetect.DeviceTypeOOH       // Out of Home
)

// DeviceType model
type DeviceType struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Codename    string `json:"codename" gorm:"unique"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Active types.ActiveStatus `gorm:"type:ActiveStatus" json:"active,omitempty"`

	// Time marks
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// TableName in database
func (m *DeviceType) TableName() string {
	return `type_device_type`
}

// RBACResourceName returns the name of the resource for the RBAC
func (m *DeviceType) RBACResourceName() string {
	return "device_type"
}

// DeviceTypeList is a list of DeviceType
var DeviceTypeList = []*DeviceType{
	{ID: uint64(DeviceTypeUnknown), Codename: "unknown", Name: "Unknown", Description: "Unknown device type", Active: types.StatusActive},
	{ID: uint64(DeviceTypeMobile), Codename: "mobile", Name: "Mobile", Description: "Mobile device", Active: types.StatusActive},
	{ID: uint64(DeviceTypePC), Codename: "pc", Name: "PC", Description: "Personal Computer", Active: types.StatusActive},
	{ID: uint64(DeviceTypeTV), Codename: "tv", Name: "TV", Description: "TV device", Active: types.StatusActive},
	{ID: uint64(DeviceTypePhone), Codename: "phone", Name: "Phone", Description: "Phone device", Active: types.StatusActive},
	{ID: uint64(DeviceTypeTablet), Codename: "tablet", Name: "Tablet", Description: "Tablet device", Active: types.StatusActive},
	{ID: uint64(DeviceTypeConnected), Codename: "connected", Name: "Connected", Description: "Connected device (Console, EReader)", Active: types.StatusActive},
	{ID: uint64(DeviceTypeSetTopBox), Codename: "settopbox", Name: "SetTopBox", Description: "SetTopBox device", Active: types.StatusActive},
	{ID: uint64(DeviceTypeWatch), Codename: "watch", Name: "Watch", Description: "Watch device", Active: types.StatusActive},
	{ID: uint64(DeviceTypeGlasses), Codename: "glasses", Name: "Glasses", Description: "Glasses device", Active: types.StatusActive},
	{ID: uint64(DeviceTypeOOH), Codename: "ooh", Name: "OOH", Description: "Out of Home device", Active: types.StatusActive},
}

// DeviceMaker model
type DeviceMaker struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Codename    string `json:"codename" gorm:"unique"`
	Name        string `json:"name"`
	Description string `json:"description"`

	// "github.com/IGLOU-EU/go-wildcard/v2" package syntax
	MatchExp string `json:"match_exp,omitempty"`

	Active types.ActiveStatus `gorm:"type:ActiveStatus" json:"active,omitempty"`

	Models []*DeviceModel `json:"models,omitempty" gorm:"foreignKey:MakerCodename;references:Codename"`

	// Time marks
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// TableName in database
func (m *DeviceMaker) TableName() string {
	return `type_device_maker`
}

// RBACResourceName returns the name of the resource for the RBAC
func (m *DeviceMaker) RBACResourceName() string {
	return "device_maker"
}

// DeviceModel model
type DeviceModel struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Codename    string `json:"codename" gorm:"unique"`
	Name        string `json:"name"`
	Description string `json:"description"`

	ParentID uint64       `json:"parent_id,omitempty"`
	Parent   *DeviceModel `json:"parent,omitempty" gorm:"foreignKey:ParentID;references:ID"`

	YearRelease int `json:"year_release,omitempty"`

	Active types.ActiveStatus `gorm:"type:ActiveStatus" json:"active,omitempty"`

	MatchExp string `json:"match_exp,omitempty"`

	// Link to device maker
	MakerCodename string       `json:"maker_codename,omitempty"`
	Maker         *DeviceMaker `json:"maker,omitempty" gorm:"foreignKey:MakerCodename;references:Codename"`

	// Device type
	TypeCodename string      `json:"type_codename,omitempty"`
	Type         *DeviceType `json:"type,omitempty" gorm:"foreignKey:TypeCodename;references:Codename"`

	// Versions of the model
	Versions []*DeviceModel `json:"versions,omitempty" gorm:"foreignKey:ParentID;references:ID"`

	// Time marks
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// TableName in database
func (m *DeviceModel) TableName() string {
	return `type_device_model`
}

// RBACResourceName returns the name of the resource for the RBAC
func (m *DeviceModel) RBACResourceName() string {
	return "device_model"
}
