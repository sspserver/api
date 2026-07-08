package connectors

import (
	"context"

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/xtypes"
	"github.com/geniusrabbit/blaze-api/repository"
	"github.com/geniusrabbit/blaze-api/server/graphql/connectors"

	"github.com/sspserver/api/pkg/repository/adformat"
	"github.com/sspserver/api/pkg/repository/application"
	"github.com/sspserver/api/pkg/repository/browser"
	"github.com/sspserver/api/pkg/repository/category"
	"github.com/sspserver/api/pkg/repository/devicemaker"
	"github.com/sspserver/api/pkg/repository/devicemodel"
	"github.com/sspserver/api/pkg/repository/os"
	"github.com/sspserver/api/pkg/repository/rtbsource"
	"github.com/sspserver/api/pkg/repository/statistic"
	"github.com/sspserver/api/pkg/repository/trafficrouter"
	"github.com/sspserver/api/pkg/repository/zone"
	gqlmodels "github.com/sspserver/api/pkg/server/graphql/models"
)

// RTBSourceConnection implements collection accessor interface with pagination.
type RTBSourceConnection = connectors.CollectionConnection[*gqlmodels.RTBSource]

// NewRTBSourceConnection based on query object
func NewRTBSourceConnection(ctx context.Context, rtbSourceAccessor rtbsource.Usecase, filter *gqlmodels.RTBSourceListFilter, order []*gqlmodels.RTBSourceListOrder, page *gqlmodels.Page) *RTBSourceConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.RTBSource]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.RTBSource, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.RTBSourceListOrder, res *rtbsource.ListOrder) { val.Fill(res) })
			list, err := rtbSourceAccessor.FetchList(ctx,
				filter.Filter(), &newOrder, page.Pagination())
			return gqlmodels.FromRTBSourceModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return rtbSourceAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// AdFormatConnection implements collection accessor interface with pagination.
type AdFormatConnection = connectors.CollectionConnection[*gqlmodels.AdFormat]

// NewAdFormatConnection based on query object
func NewAdFormatConnection(ctx context.Context, adFormatAccessor adformat.Usecase, filter *gqlmodels.AdFormatListFilter, order *gqlmodels.AdFormatListOrder, page *gqlmodels.Page) *AdFormatConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.AdFormat]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.AdFormat, error) {
			list, err := adFormatAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromAdFormatModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return adFormatAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// CategoryConnection implements collection accessor interface with pagination.
type CategoryConnection = connectors.CollectionConnection[*gqlmodels.Category]

// NewCategoryConnection based on query object
func NewCategoryConnection(ctx context.Context, categoryAccessor category.Usecase, filter *gqlmodels.CategoryListFilter, order *gqlmodels.CategoryListOrder, page *gqlmodels.Page) *CategoryConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Category]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Category, error) {
			newFilter := filter.Filter()
			list, err := categoryAccessor.FetchList(ctx,
				&repository.PreloadOption{
					Fields: gocast.IfThen(newFilter.IsChildrensPreload(),
						[]string{`Parent`, `Childrens`}, []string{`Parent`}),
				},
				newFilter, order.Order(), page.Pagination())
			return gqlmodels.FromCategoryModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return categoryAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// OSConnection implements collection accessor interface with pagination.
type OSConnection = connectors.CollectionConnection[*gqlmodels.Os]

// NewOSConnection based on query object
func NewOSConnection(ctx context.Context, osAccessor os.Usecase, filter *gqlmodels.OSListFilter, order []*gqlmodels.OSListOrder, page *gqlmodels.Page) *OSConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Os]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Os, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.OSListOrder, res *os.ListOrder) { val.Fill(res) })
			newFilter := filter.Filter()
			preloadOptions := (*repository.PreloadOption)(nil)
			if newFilter.IsChildrensPreload() {
				preloadOptions = &repository.PreloadOption{Fields: []string{`Versions`}}
			}
			list, err := osAccessor.FetchList(ctx,
				preloadOptions, newFilter, &newOrder, page.Pagination())
			return gqlmodels.FromOSModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return osAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// BrowserConnection implements collection accessor interface with pagination.
type BrowserConnection = connectors.CollectionConnection[*gqlmodels.Browser]

// NewBrowserConnection based on query object
func NewBrowserConnection(ctx context.Context, browserAccessor browser.Usecase, filter *gqlmodels.BrowserListFilter, order []*gqlmodels.BrowserListOrder, page *gqlmodels.Page) *BrowserConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Browser]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Browser, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.BrowserListOrder, res *browser.ListOrder) { val.Fill(res) })
			newFilter := filter.Filter()
			preloadOptions := (*repository.PreloadOption)(nil)
			if newFilter.IsChildrensPreload() {
				preloadOptions = &repository.PreloadOption{Fields: []string{`Versions`}}
			}
			list, err := browserAccessor.FetchList(ctx,
				preloadOptions, filter.Filter(), &newOrder, page.Pagination())
			return gqlmodels.FromBrowserModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return browserAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// DeviceMakerConnection implements collection accessor interface with pagination.
type DeviceMakerConnection = connectors.CollectionConnection[*gqlmodels.DeviceMaker]

// NewDeviceMakerConnection based on query object
func NewDeviceMakerConnection(ctx context.Context, deviceMakerAccessor devicemaker.Usecase, filter *gqlmodels.DeviceMakerListFilter, order []*gqlmodels.DeviceMakerListOrder, page *gqlmodels.Page) *DeviceMakerConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.DeviceMaker]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.DeviceMaker, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.DeviceMakerListOrder, res *devicemaker.ListOrder) { val.Fill(res) })
			list, err := deviceMakerAccessor.FetchList(ctx,
				&repository.PreloadOption{Fields: []string{`Models`}},
				filter.Filter(), &newOrder, page.Pagination())
			return gqlmodels.FromDeviceMakerModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return deviceMakerAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// DeviceModelConnection implements collection accessor interface with pagination.
type DeviceModelConnection = connectors.CollectionConnection[*gqlmodels.DeviceModel]

// NewDeviceModelConnection based on query object
func NewDeviceModelConnection(ctx context.Context, deviceModelAccessor devicemodel.Usecase, filter *gqlmodels.DeviceModelListFilter, order []*gqlmodels.DeviceModelListOrder, page *gqlmodels.Page) *DeviceModelConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.DeviceModel]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.DeviceModel, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.DeviceModelListOrder, res *devicemodel.ListOrder) { val.Fill(res) })
			list, err := deviceModelAccessor.FetchList(ctx,
				filter.Filter(), &newOrder, page.Pagination())
			return gqlmodels.FromDeviceModelModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return deviceModelAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// ApplicationConnection implements collection accessor interface with pagination.
type ApplicationConnection = connectors.CollectionConnection[*gqlmodels.Application]

// NewApplicationConnection based on query object
func NewApplicationConnection(ctx context.Context, applicationAccessor application.Usecase, filter *gqlmodels.ApplicationListFilter, order *gqlmodels.ApplicationListOrder, page *gqlmodels.Page) *ApplicationConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Application]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Application, error) {
			list, err := applicationAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromApplicationModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return applicationAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// ZoneConnection implements collection accessor interface with pagination.
type ZoneConnection = connectors.CollectionConnection[*gqlmodels.Zone]

// NewZoneConnection based on query object
func NewZoneConnection(ctx context.Context, zoneAccessor zone.Usecase, filter *gqlmodels.ZoneListFilter, order *gqlmodels.ZoneListOrder, page *gqlmodels.Page) *ZoneConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.Zone]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.Zone, error) {
			list, err := zoneAccessor.FetchList(ctx,
				filter.Filter(), order.Order(), page.Pagination())
			return gqlmodels.FromZoneModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return zoneAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}

// StatisticAdItemConnection implements collection accessor interface with pagination.
type StatisticAdItemConnection = connectors.CollectionConnection[*gqlmodels.StatisticAdItem]

// NewStatisticAdItemConnection based on query object
func NewStatisticAdItemConnection(ctx context.Context, statisticAccessor statistic.Usecase, filter *gqlmodels.StatisticAdListFilter, group []gqlmodels.StatisticKey, order []*gqlmodels.StatisticAdKeyOrder, page *gqlmodels.Page) *StatisticAdItemConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.StatisticAdItem]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.StatisticAdItem, error) {
			list, err := statisticAccessor.Statistic(ctx,
				filter.Filter(), gqlmodels.StatisticGroup(group),
				gqlmodels.StatisticAdListOrder(order, group), page.Pagination())
			return gqlmodels.FromStatisticAdItemModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return statisticAccessor.Count(ctx, filter.Filter(), gqlmodels.StatisticGroup(group))
		},
	}, page)
}

// TrafficRouterConnection implements collection accessor interface with pagination.
type TrafficRouterConnection = connectors.CollectionConnection[*gqlmodels.TrafficRouter]

// NewTrafficRouterConnection based on query object
func NewTrafficRouterConnection(ctx context.Context, trafficRouterAccessor trafficrouter.Usecase, filter *gqlmodels.TrafficRouterListFilter, order []*gqlmodels.TrafficRouterListOrder, page *gqlmodels.Page) *TrafficRouterConnection {
	return connectors.NewCollectionConnection(ctx, &connectors.DataAccessorFunc[*gqlmodels.TrafficRouter]{
		FetchDataListFunc: func(ctx context.Context) ([]*gqlmodels.TrafficRouter, error) {
			newOrder := xtypes.SliceReduce(order,
				func(val *gqlmodels.TrafficRouterListOrder, res *trafficrouter.ListOrder) { val.Fill(res) })
			list, err := trafficRouterAccessor.FetchList(ctx,
				filter.Filter(), &newOrder, page.Pagination())
			return gqlmodels.FromTrafficRouterModelList(list), err
		},
		CountDataFunc: func(ctx context.Context) (int64, error) {
			return trafficRouterAccessor.Count(ctx, filter.Filter())
		},
	}, page)
}
