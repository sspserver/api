package connectors

import (
	"context"

	"github.com/demdxx/gocast/v2"
	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"

	"github.com/sspserver/api/internal/repository/adformat"
	"github.com/sspserver/api/internal/repository/application"
	"github.com/sspserver/api/internal/repository/browser"
	"github.com/sspserver/api/internal/repository/category"
	"github.com/sspserver/api/internal/repository/devicemaker"
	"github.com/sspserver/api/internal/repository/devicemodel"
	"github.com/sspserver/api/internal/repository/os"
	"github.com/sspserver/api/internal/repository/rtbsource"
	"github.com/sspserver/api/internal/repository/zone"
	gqlmodels "github.com/sspserver/api/internal/server/graphql/models"
)

// RTBSourceConnection implements collection accessor interface with pagination.
type RTBSourceConnection = connectors.CollectionConnection[gqlmodels.RTBSource, gqlmodels.RTBSourceEdge]

// NewRTBSourceConnection based on query object
func NewRTBSourceConnection(ctx context.Context, rtbSourceAccessor rtbsource.Usecase, filter *gqlmodels.RTBSourceListFilter, order *gqlmodels.RTBSourceListOrder, page *gqlmodels.Page) *RTBSourceConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.RTBSource, gqlmodels.RTBSourceEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.RTBSource, error) {
			list, err := rtbSourceAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromRTBSourceModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return rtbSourceAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.RTBSource) *gqlmodels.RTBSourceEdge {
			return &gqlmodels.RTBSourceEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}

// AdFormatConnection implements collection accessor interface with pagination.
type AdFormatConnection = connectors.CollectionConnection[gqlmodels.AdFormat, gqlmodels.AdFormatEdge]

// NewAdFormatConnection based on query object
func NewAdFormatConnection(ctx context.Context, adFormatAccessor adformat.Usecase, filter *gqlmodels.AdFormatListFilter, order *gqlmodels.AdFormatListOrder, page *gqlmodels.Page) *AdFormatConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.AdFormat, gqlmodels.AdFormatEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.AdFormat, error) {
			list, err := adFormatAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromAdFormatModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return adFormatAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.AdFormat) *gqlmodels.AdFormatEdge {
			return &gqlmodels.AdFormatEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}

// CategoryConnection implements collection accessor interface with pagination.
type CategoryConnection = connectors.CollectionConnection[gqlmodels.Category, gqlmodels.CategoryEdge]

// NewCategoryConnection based on query object
func NewCategoryConnection(ctx context.Context, categoryAccessor category.Usecase, filter *gqlmodels.CategoryListFilter, order *gqlmodels.CategoryListOrder, page *gqlmodels.Page) *CategoryConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.Category, gqlmodels.CategoryEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Category, error) {
			list, err := categoryAccessor.FetchList(ctx,
				&repository.PreloadOption{Fields: []string{`Parent`}},
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromCategoryModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return categoryAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.Category) *gqlmodels.CategoryEdge {
			return &gqlmodels.CategoryEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}

// OSConnection implements collection accessor interface with pagination.
type OSConnection = connectors.CollectionConnection[gqlmodels.Os, gqlmodels.OSEdge]

// NewOSConnection based on query object
func NewOSConnection(ctx context.Context, osAccessor os.Usecase, filter *gqlmodels.OSListFilter, order *gqlmodels.OSListOrder, page *gqlmodels.Page) *OSConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.Os, gqlmodels.OSEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Os, error) {
			list, err := osAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromOSModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return osAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.Os) *gqlmodels.OSEdge {
			return &gqlmodels.OSEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}

// BrowserConnection implements collection accessor interface with pagination.
type BrowserConnection = connectors.CollectionConnection[gqlmodels.Browser, gqlmodels.BrowserEdge]

// NewBrowserConnection based on query object
func NewBrowserConnection(ctx context.Context, browserAccessor browser.Usecase, filter *gqlmodels.BrowserListFilter, order *gqlmodels.BrowserListOrder, page *gqlmodels.Page) *BrowserConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.Browser, gqlmodels.BrowserEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Browser, error) {
			list, err := browserAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromBrowserModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return browserAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.Browser) *gqlmodels.BrowserEdge {
			return &gqlmodels.BrowserEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}

// DeviceMakerConnection implements collection accessor interface with pagination.
type DeviceMakerConnection = connectors.CollectionConnection[gqlmodels.DeviceMaker, gqlmodels.DeviceMakerEdge]

// NewDeviceMakerConnection based on query object
func NewDeviceMakerConnection(ctx context.Context, deviceMakerAccessor devicemaker.Usecase, filter *gqlmodels.DeviceMakerListFilter, order *gqlmodels.DeviceMakerListOrder, page *gqlmodels.Page) *DeviceMakerConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.DeviceMaker, gqlmodels.DeviceMakerEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.DeviceMaker, error) {
			list, err := deviceMakerAccessor.FetchList(ctx,
				&repository.PreloadOption{Fields: []string{`Models`}},
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromDeviceMakerModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return deviceMakerAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.DeviceMaker) *gqlmodels.DeviceMakerEdge {
			return &gqlmodels.DeviceMakerEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}

// DeviceModelConnection implements collection accessor interface with pagination.
type DeviceModelConnection = connectors.CollectionConnection[gqlmodels.DeviceModel, gqlmodels.DeviceModelEdge]

// NewDeviceModelConnection based on query object
func NewDeviceModelConnection(ctx context.Context, deviceModelAccessor devicemodel.Usecase, filter *gqlmodels.DeviceModelListFilter, order *gqlmodels.DeviceModelListOrder, page *gqlmodels.Page) *DeviceModelConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.DeviceModel, gqlmodels.DeviceModelEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.DeviceModel, error) {
			list, err := deviceModelAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromDeviceModelModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return deviceModelAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.DeviceModel) *gqlmodels.DeviceModelEdge {
			return &gqlmodels.DeviceModelEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}

// ApplicationConnection implements collection accessor interface with pagination.
type ApplicationConnection = connectors.CollectionConnection[gqlmodels.Application, gqlmodels.ApplicationEdge]

// NewApplicationConnection based on query object
func NewApplicationConnection(ctx context.Context, applicationAccessor application.Usecase, filter *gqlmodels.ApplicationListFilter, order *gqlmodels.ApplicationListOrder, page *gqlmodels.Page) *ApplicationConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.Application, gqlmodels.ApplicationEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Application, error) {
			list, err := applicationAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromApplicationModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return applicationAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.Application) *gqlmodels.ApplicationEdge {
			return &gqlmodels.ApplicationEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}

// ZoneConnection implements collection accessor interface with pagination.
type ZoneConnection = connectors.CollectionConnection[gqlmodels.Zone, gqlmodels.ZoneEdge]

// NewZoneConnection based on query object
func NewZoneConnection(ctx context.Context, zoneAccessor zone.Usecase, filter *gqlmodels.ZoneListFilter, order *gqlmodels.ZoneListOrder, page *gqlmodels.Page) *ZoneConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[gqlmodels.Zone, gqlmodels.ZoneEdge]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Zone, error) {
			list, err := zoneAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromZoneModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return zoneAccessor.Count(ctx, filter.Filter())
		},
		ConvertToEdgeFunc: func(obj *gqlmodels.Zone) *gqlmodels.ZoneEdge {
			return &gqlmodels.ZoneEdge{
				Cursor: gocast.Str(obj.ID),
				Node:   obj,
			}
		},
	}, page)
}
