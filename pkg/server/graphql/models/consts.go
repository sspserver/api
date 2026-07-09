package models

import (
	"github.com/geniusrabbit/adcorelib/admodels/types"
	qmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"
)

func FromApproveStatus(status types.ApproveStatus) qmodels.ApproveStatus {
	switch status {
	case types.StatusApproved:
		return qmodels.ApproveStatusApproved
	case types.StatusRejected:
		return qmodels.ApproveStatusRejected
	default:
		return qmodels.ApproveStatusPending
	}
}

func ApproveStatusFrom(status qmodels.ApproveStatus) types.ApproveStatus {
	switch status {
	case qmodels.ApproveStatusApproved:
		return types.StatusApproved
	case qmodels.ApproveStatusRejected:
		return types.StatusRejected
	}
	return types.StatusPending
}

func ApproveStatusPtr(status *qmodels.ApproveStatus) *types.ApproveStatus {
	if status == nil {
		return nil
	}
	s := ApproveStatusFrom(*status)
	return &s
}

func FromActiveStatus(status types.ActiveStatus) qmodels.ActiveStatus {
	switch status {
	case types.StatusActive:
		return qmodels.ActiveStatusActive
	case types.StatusPause:
		return qmodels.ActiveStatusPaused
	default:
		return qmodels.ActiveStatusPaused
	}
}

func ActiveStatusFrom(status ActiveStatus) types.ActiveStatus {
	switch status {
	case qmodels.ActiveStatusActive:
		return types.StatusActive
	case qmodels.ActiveStatusPaused:
		return types.StatusPause
	}
	return types.StatusPause
}

func ActiveStatusPtr(status *ActiveStatus) *types.ActiveStatus {
	if status == nil {
		return nil
	}
	s := ActiveStatusFrom(*status)
	return &s
}

func FromRTBRequestType(t types.RTBRequestType) RTBRequestFormatType {
	switch t {
	case types.RTBRequestTypeJSON:
		return RTBRequestFormatTypeJSON
	case types.RTBRequestTypeXML:
		return RTBRequestFormatTypeXML
	}
	return RTBRequestFormatTypeUndefined
}

func (e *RTBRequestFormatType) RequestType() types.RTBRequestType {
	if e == nil {
		return types.RTBRequestTypeUndefined
	}
	switch *e {
	case RTBRequestFormatTypeJSON:
		return types.RTBRequestTypeJSON
	case RTBRequestFormatTypeXML:
		return types.RTBRequestTypeXML
	}
	return types.RTBRequestTypeUndefined
}

func (e *RTBRequestFormatType) RequestTypePtr() *types.RTBRequestType {
	if e == nil {
		return nil
	}
	t := e.RequestType()
	return &t
}

func FromAuctionType(t types.AuctionType) AuctionType {
	switch t {
	case types.FirstPriceAuctionType:
		return AuctionTypeFirstPrice
	case types.SecondPriceAuctionType:
		return AuctionTypeSecondPrice
	}
	return AuctionTypeUndefined
}

func (e *AuctionType) AuctionType() types.AuctionType {
	if e == nil {
		return types.UndefinedAuctionType
	}
	switch *e {
	case AuctionTypeFirstPrice:
		return types.FirstPriceAuctionType
	case AuctionTypeSecondPrice:
		return types.SecondPriceAuctionType
	}
	return types.UndefinedAuctionType
}

func (e *AuctionType) AuctionTypePtr() *types.AuctionType {
	if e == nil {
		return nil
	}
	t := e.AuctionType()
	return &t
}

func FromAnyOnlyExclude(status int) AnyOnlyExclude {
	switch status {
	case 0:
		return AnyOnlyExcludeAny
	case 1:
		return AnyOnlyExcludeOnly
	case 2, -1:
		return AnyOnlyExcludeExclude
	}
	return AnyOnlyExcludeAny
}

func (e *AnyOnlyExclude) Int() int {
	if e == nil {
		return 0
	}
	switch *e {
	case AnyOnlyExcludeAny:
		return 0
	case AnyOnlyExcludeOnly:
		return 1
	case AnyOnlyExcludeExclude:
		return 2
	}
	return 0
}

func (e *AnyOnlyExclude) IntPtr() *int {
	if e == nil {
		return nil
	}
	i := e.Int()
	return &i
}

func FromAnyIPv4IPv6(status int) AnyIPv4IPv6 {
	switch status {
	case 0:
		return AnyIPv4IPv6Any
	case 1:
		return AnyIPv4IPv6IPv4
	case 2:
		return AnyIPv4IPv6IPv6
	}
	return AnyIPv4IPv6Any
}

func (e *AnyIPv4IPv6) Int() int {
	if e == nil {
		return 0
	}
	switch *e {
	case AnyIPv4IPv6Any:
		return 0
	case AnyIPv4IPv6IPv4:
		return 1
	case AnyIPv4IPv6IPv6:
		return 2
	}
	return 0
}

func (e *AnyIPv4IPv6) IntPtr() *int {
	if e == nil {
		return nil
	}
	i := e.Int()
	return &i
}

func FromPrivateStatus(status types.PrivateStatus) PrivateStatus {
	switch status {
	case types.StatusPrivate:
		return PrivateStatusPrivate
	case types.StatusPublic:
		return PrivateStatusPublic
	}
	return PrivateStatusPublic
}

func (e *PrivateStatus) ModelStatus() types.PrivateStatus {
	if e == nil {
		return types.StatusPublic
	}
	switch *e {
	case PrivateStatusPrivate:
		return types.StatusPrivate
	case PrivateStatusPublic:
		return types.StatusPublic
	}
	return types.StatusPublic
}

func (e *PrivateStatus) ModelStatusPtr() *types.PrivateStatus {
	if e == nil {
		return nil
	}
	t := e.ModelStatus()
	return &t
}

func FromApplicationType(tp types.ApplicationType) ApplicationType {
	switch tp {
	case types.ApplicationSite:
		return ApplicationTypeSite
	case types.ApplicationApp:
		return ApplicationTypeApp
	case types.ApplicationGame:
		return ApplicationTypeGame
	}
	return ApplicationTypeUndefined
}

func (e *ApplicationType) ModelType() types.ApplicationType {
	if e == nil {
		return types.ApplicationUndefined
	}
	switch *e {
	case ApplicationTypeSite:
		return types.ApplicationSite
	case ApplicationTypeApp:
		return types.ApplicationApp
	case ApplicationTypeGame:
		return types.ApplicationGame
	}
	return types.ApplicationUndefined
}

func FromPlatformType(tp types.PlatformType) PlatformType {
	switch tp {
	case types.PlatformWeb:
		return PlatformTypeWeb
	case types.PlatformDesktop:
		return PlatformTypeDesktop
	case types.PlatformMobile:
		return PlatformTypeMobile
	case types.PlatformSmartPhone:
		return PlatformTypeSmartPhone
	case types.PlatformTablet:
		return PlatformTypeTablet
	case types.PlatformSmartTV:
		return PlatformTypeSmartTv
	case types.PlatformGameStation:
		return PlatformTypeGameStation
	case types.PlatformSmartWatch:
		return PlatformTypeSmartWatch
	case types.PlatformVR:
		return PlatformTypeVr
	case types.PlatformSmartGlasses:
		return PlatformTypeSmartGlasses
	case types.PlatformSmartBillboard:
		return PlatformTypeSmartBillboard
	}
	return PlatformTypeUndefined
}

func (e *PlatformType) ModelType() types.PlatformType {
	if e == nil {
		return types.PlatformUndefined
	}
	switch *e {
	case PlatformTypeWeb:
		return types.PlatformWeb
	case PlatformTypeDesktop:
		return types.PlatformDesktop
	case PlatformTypeMobile:
		return types.PlatformMobile
	case PlatformTypeSmartPhone:
		return types.PlatformSmartPhone
	case PlatformTypeTablet:
		return types.PlatformTablet
	case PlatformTypeSmartTv:
		return types.PlatformSmartTV
	case PlatformTypeGameStation:
		return types.PlatformGameStation
	case PlatformTypeSmartWatch:
		return types.PlatformSmartWatch
	case PlatformTypeVr:
		return types.PlatformVR
	case PlatformTypeSmartGlasses:
		return types.PlatformSmartGlasses
	case PlatformTypeSmartBillboard:
		return types.PlatformSmartBillboard
	}
	return types.PlatformUndefined
}

func FromPricingModel(pm types.PricingModel) PricingModel {
	switch pm {
	case types.PricingModelCPM:
		return PricingModelCpm
	case types.PricingModelCPC:
		return PricingModelCpc
	case types.PricingModelCPA:
		return PricingModelCpa
	}
	return PricingModelUndefined
}

func PricingModelFrom(pm PricingModel) types.PricingModel {
	return pm.ModelType()
}

func (pm *PricingModel) ModelType() types.PricingModel {
	if pm == nil {
		return types.PricingModelUndefined
	}
	switch *pm {
	case PricingModelCpm:
		return types.PricingModelCPM
	case PricingModelCpc:
		return types.PricingModelCPC
	case PricingModelCpa:
		return types.PricingModelCPA
	}
	return types.PricingModelUndefined
}
