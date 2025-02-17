package models

import (
	qmodels "github.com/geniusrabbit/blaze-api/server/graphql/models"

	"github.com/sspserver/api/models"
)

func FromApproveStatus(status models.ApproveStatus) qmodels.ApproveStatus {
	switch status {
	case models.StatusApproved:
		return qmodels.ApproveStatusApproved
	case models.StatusRejected:
		return qmodels.ApproveStatusRejected
	default:
		return qmodels.ApproveStatusPending
	}
}

func ApproveStatusFrom(status qmodels.ApproveStatus) models.ApproveStatus {
	switch status {
	case qmodels.ApproveStatusApproved:
		return models.StatusApproved
	case qmodels.ApproveStatusRejected:
		return models.StatusRejected
	}
	return models.StatusPending
}

func ApproveStatusPtr(status *qmodels.ApproveStatus) *models.ApproveStatus {
	if status == nil {
		return nil
	}
	s := ApproveStatusFrom(*status)
	return &s
}

func FromActiveStatus(status models.ActiveStatus) ActiveStatus {
	switch status {
	case models.StatusActive:
		return qmodels.ActiveStatusActive
	case models.StatusPause:
		return qmodels.ActiveStatusPaused
	default:
		return qmodels.ActiveStatusPaused
	}
}

func ActiveStatusFrom(status ActiveStatus) models.ActiveStatus {
	switch status {
	case qmodels.ActiveStatusActive:
		return models.StatusActive
	case qmodels.ActiveStatusPaused:
		return models.StatusPause
	}
	return models.StatusPause
}

func ActiveStatusPtr(status *ActiveStatus) *models.ActiveStatus {
	if status == nil {
		return nil
	}
	s := ActiveStatusFrom(*status)
	return &s
}

func FromRTBRequestType(t models.RTBRequestType) RTBRequestFormatType {
	switch t {
	case models.RTBRequestTypeJSON:
		return RTBRequestFormatTypeJSON
	case models.RTBRequestTypeXML:
		return RTBRequestFormatTypeXML
	}
	return RTBRequestFormatTypeUndefined
}

func (e RTBRequestFormatType) RequestType() models.RTBRequestType {
	switch e {
	case RTBRequestFormatTypeJSON:
		return models.RTBRequestTypeJSON
	case RTBRequestFormatTypeXML:
		return models.RTBRequestTypeXML
	}
	return models.RTBRequestTypeUndefined
}

func (e *RTBRequestFormatType) RequestTypePtr() *models.RTBRequestType {
	if e == nil {
		return nil
	}
	t := e.RequestType()
	return &t
}

func FromAuctionType(t models.AuctionType) AuctionType {
	switch t {
	case models.AuctionTypeFirstPrice:
		return AuctionTypeFirstPrice
	case models.AuctionTypeSecondPrice:
		return AuctionTypeSecondPrice
	}
	return AuctionTypeUndefined
}

func (e AuctionType) AuctionType() models.AuctionType {
	switch e {
	case AuctionTypeFirstPrice:
		return models.AuctionTypeFirstPrice
	case AuctionTypeSecondPrice:
		return models.AuctionTypeSecondPrice
	}
	return models.AuctionTypeUndefined
}

func (e *AuctionType) AuctionTypePtr() *models.AuctionType {
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

func (e AnyOnlyExclude) Int() int {
	switch e {
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

func (e AnyIPv4IPv6) Int() int {
	switch e {
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

func FromPrivateStatus(status models.PrivateStatus) PrivateStatus {
	switch status {
	case models.StatusPrivate:
		return PrivateStatusPrivate
	case models.StatusPublic:
		return PrivateStatusPublic
	}
	return PrivateStatusPublic
}

func (e *PrivateStatus) ModelStatus() models.PrivateStatus {
	if e == nil {
		return models.StatusPublic
	}
	switch *e {
	case PrivateStatusPrivate:
		return models.StatusPrivate
	case PrivateStatusPublic:
		return models.StatusPublic
	}
	return models.StatusPublic
}

func (e *PrivateStatus) ModelStatusPtr() *models.PrivateStatus {
	if e == nil {
		return nil
	}
	t := e.ModelStatus()
	return &t
}

func FromApplicationType(tp models.ApplicationType) ApplicationType {
	switch tp {
	case models.ApplicationSite:
		return ApplicationTypeSite
	case models.ApplicationApp:
		return ApplicationTypeApp
	case models.ApplicationGame:
		return ApplicationTypeGame
	}
	return ApplicationTypeUndefined
}

func (e ApplicationType) ModelType() models.ApplicationType {
	switch e {
	case ApplicationTypeSite:
		return models.ApplicationSite
	case ApplicationTypeApp:
		return models.ApplicationApp
	case ApplicationTypeGame:
		return models.ApplicationGame
	}
	return models.ApplicationUndefined
}

func FromPlatformType(tp models.PlatformType) PlatformType {
	switch tp {
	case models.PlatformWeb:
		return PlatformTypeWeb
	case models.PlatformDesktop:
		return PlatformTypeDesktop
	case models.PlatformMobile:
		return PlatformTypeMobile
	case models.PlatformSmartPhone:
		return PlatformTypeSmartPhone
	case models.PlatformTablet:
		return PlatformTypeTablet
	case models.PlatformSmartTV:
		return PlatformTypeSmartTv
	case models.PlatformGameStation:
		return PlatformTypeGameStation
	case models.PlatformSmartWatch:
		return PlatformTypeSmartWatch
	case models.PlatformVR:
		return PlatformTypeVr
	case models.PlatformSmartGlasses:
		return PlatformTypeSmartGlasses
	case models.PlatformSmartBillboard:
		return PlatformTypeSmartBillboard
	}
	return PlatformTypeUndefined
}

func (e PlatformType) ModelType() models.PlatformType {
	switch e {
	case PlatformTypeWeb:
		return models.PlatformWeb
	case PlatformTypeDesktop:
		return models.PlatformDesktop
	case PlatformTypeMobile:
		return models.PlatformMobile
	case PlatformTypeSmartPhone:
		return models.PlatformSmartPhone
	case PlatformTypeTablet:
		return models.PlatformTablet
	case PlatformTypeSmartTv:
		return models.PlatformSmartTV
	case PlatformTypeGameStation:
		return models.PlatformGameStation
	case PlatformTypeSmartWatch:
		return models.PlatformSmartWatch
	case PlatformTypeVr:
		return models.PlatformVR
	case PlatformTypeSmartGlasses:
		return models.PlatformSmartGlasses
	case PlatformTypeSmartBillboard:
		return models.PlatformSmartBillboard
	}
	return models.PlatformUndefined
}

func FromPricingModel(pm models.PricingModel) PricingModel {
	switch pm {
	case models.PricingModelCPM:
		return PricingModelCpm
	case models.PricingModelCPC:
		return PricingModelCpc
	case models.PricingModelCPA:
		return PricingModelCpa
	}
	return PricingModelUndefined
}

func PricingModelFrom(pm PricingModel) models.PricingModel {
	return pm.ModelType()
}

func (pm PricingModel) ModelType() models.PricingModel {
	switch pm {
	case PricingModelCpm:
		return models.PricingModelCPM
	case PricingModelCpc:
		return models.PricingModelCPC
	case PricingModelCpa:
		return models.PricingModelCPA
	}
	return models.PricingModelUndefined
}
