package models

import (
	"github.com/geniusrabbit/adcorelib/admodels/types"
	"github.com/geniusrabbit/adcorelib/adtype"
	"github.com/geniusrabbit/adcorelib/models"
)

// Models set of types
type (
	Format         = models.Format
	Application    = models.Application
	Zone           = models.Zone
	RTBSource      = models.RTBSource
	Category       = models.Category
	OS             = models.OS
	Browser        = models.Browser
	BrowserVersion = models.BrowserVersion
	DeviceType     = models.DeviceType
	DeviceMaker    = models.DeviceMaker
	DeviceModel    = models.DeviceModel
)

var DeviceTypeList = models.DeviceTypeList

type (
	ApproveStatus   = types.ApproveStatus
	ActiveStatus    = types.ActiveStatus
	PrivateStatus   = types.PrivateStatus
	ZoneType        = types.ZoneType
	TypeSex         = adtype.TypeSex
	RTBRequestType  = types.RTBRequestType
	AuctionType     = types.AuctionType
	ApplicationType = types.ApplicationType
	PlatformType    = types.PlatformType
	PricingModel    = types.PricingModel
)

type (
	RTBSourceFlags = models.RTBSourceFlags
)

// PricingModel set of types
const (
	PricingModelUndefined = types.PricingModelUndefined
	PricingModelCPM       = types.PricingModelCPM
	PricingModelCPC       = types.PricingModelCPC
	PricingModelCPA       = types.PricingModelCPA
)

// ActiveStatus set of types
const (
	StatusPending      = types.StatusPending
	StatusApproved     = types.StatusApproved
	StatusRejected     = types.StatusRejected
	StatusPendingName  = types.StatusPendingName
	StatusApprovedName = types.StatusApprovedName
	StatusRejectedName = types.StatusRejectedName
	StatusActive       = types.StatusActive
	StatusPause        = types.StatusPause
	StatusPublic       = types.StatusPublic
	StatusPrivate      = types.StatusPrivate
)

// PlatformType set of types
const (
	PlatformUndefined      = types.PlatformUndefined
	PlatformWeb            = types.PlatformWeb
	PlatformDesktop        = types.PlatformDesktop
	PlatformMobile         = types.PlatformMobile
	PlatformSmartPhone     = types.PlatformSmartPhone
	PlatformTablet         = types.PlatformTablet
	PlatformSmartTV        = types.PlatformSmartTV
	PlatformGameStation    = types.PlatformGameStation
	PlatformSmartWatch     = types.PlatformSmartWatch
	PlatformVR             = types.PlatformVR
	PlatformSmartGlasses   = types.PlatformSmartGlasses
	PlatformSmartBillboard = types.PlatformSmartBillboard
)

// PlatformName set of names
const (
	PlatformUndefinedName      = types.PlatformUndefinedName
	PlatformWebName            = types.PlatformWebName
	PlatformDesktopName        = types.PlatformDesktopName
	PlatformMobileName         = types.PlatformMobileName
	PlatformSmartPhoneName     = types.PlatformSmartPhoneName
	PlatformTabletName         = types.PlatformTabletName
	PlatformSmartTVName        = types.PlatformSmartTVName
	PlatformGameStationName    = types.PlatformGameStationName
	PlatformSmartWatchName     = types.PlatformSmartWatchName
	PlatformVRName             = types.PlatformVRName
	PlatformSmartGlassesName   = types.PlatformSmartGlassesName
	PlatformSmartBillboardName = types.PlatformSmartBillboardName
)

// UserSex set of types
const (
	UserSexUndefined = adtype.UserSexUndefined
	UserSexMale      = adtype.UserSexMale
	UserSexFemale    = adtype.UserSexFemale
)

// ZoneType set of types
const (
	ZoneTypeDefault   = types.ZoneTypeDefault
	ZoneTypeSmartlink = types.ZoneTypeSmartlink
)

// RTBRequestType set of types
const (
	RTBRequestTypeUndefined       = types.RTBRequestTypeUndefined
	RTBRequestTypeJSON            = types.RTBRequestTypeJSON
	RTBRequestTypeXML             = types.RTBRequestTypeXML
	RTBRequestTypeProtoBUFF       = types.RTBRequestTypeProtoBUFF
	RTBRequestTypePOSTFormEncoded = types.RTBRequestTypePOSTFormEncoded
	RTBRequestTypePLAINTEXT       = types.RTBRequestTypePLAINTEXT
)

// AuctionType set of types
const (
	AuctionTypeUndefined   = types.UndefinedAuctionType
	AuctionTypeFirstPrice  = types.FirstPriceAuctionType
	AuctionTypeSecondPrice = types.SecondPriceAuctionType
)

// ApplicationType set of types
const (
	ApplicationUndefined = types.ApplicationUndefined
	ApplicationSite      = types.ApplicationSite
	ApplicationApp       = types.ApplicationApp
	ApplicationGame      = types.ApplicationGame
)
