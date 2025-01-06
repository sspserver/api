package client

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	archivariuspb "github.com/geniusrabbit/archivarius/internal/server/grpc"
)

// Define the API client types alias
type (
	StatisticRequest  = archivariuspb.StatisticRequest
	Filter            = archivariuspb.Filter
	FilterCondition   = archivariuspb.FilterCondition
	Order             = archivariuspb.Order
	OrderingKey       = archivariuspb.OrderingKey
	Key               = archivariuspb.Key
	StatisticResponse = archivariuspb.StatisticResponse
	Item              = archivariuspb.Item
	ItemKey           = archivariuspb.ItemKey
	Value             = archivariuspb.Value
	Condition         = archivariuspb.Condition
)

const (
	Condition_UNKNOWN = archivariuspb.Condition_UNKNOWN
	Condition_EQ      = archivariuspb.Condition_EQ
	Condition_NE      = archivariuspb.Condition_NE
	Condition_GT      = archivariuspb.Condition_GT
	Condition_GE      = archivariuspb.Condition_GE
	Condition_LT      = archivariuspb.Condition_LT
	Condition_LE      = archivariuspb.Condition_LE
	Condition_IN      = archivariuspb.Condition_IN
	Condition_NI      = archivariuspb.Condition_NI
	Condition_BT      = archivariuspb.Condition_BT
	Condition_NB      = archivariuspb.Condition_NB
	Condition_LI      = archivariuspb.Condition_LI
	Condition_NL      = archivariuspb.Condition_NL
)

const (
	Key_UNKNOWN_KEY         = archivariuspb.Key_UNKNOWN_KEY
	Key_KEY_DATEMARK        = archivariuspb.Key_KEY_DATEMARK
	Key_KEY_TIMEMARK        = archivariuspb.Key_KEY_TIMEMARK
	Key_KEY_CLUSTER         = archivariuspb.Key_KEY_CLUSTER
	Key_KEY_PROJECT_ID      = archivariuspb.Key_KEY_PROJECT_ID
	Key_KEY_ACCOUNT_ID      = archivariuspb.Key_KEY_ACCOUNT_ID
	Key_KEY_PUB_ACCOUNT_ID  = archivariuspb.Key_KEY_PUB_ACCOUNT_ID
	Key_KEY_ADV_ACCOUNT_ID  = archivariuspb.Key_KEY_ADV_ACCOUNT_ID
	Key_KEY_SOURCE_ID       = archivariuspb.Key_KEY_SOURCE_ID
	Key_KEY_ACCESS_POINT_ID = archivariuspb.Key_KEY_ACCESS_POINT_ID
	Key_KEY_PLATFORM        = archivariuspb.Key_KEY_PLATFORM
	Key_KEY_DOMAIN          = archivariuspb.Key_KEY_DOMAIN
	Key_KEY_APP_ID          = archivariuspb.Key_KEY_APP_ID
	Key_KEY_ZONE_ID         = archivariuspb.Key_KEY_ZONE_ID
	Key_KEY_CAMPAIGN_ID     = archivariuspb.Key_KEY_CAMPAIGN_ID
	Key_KEY_AD_ID           = archivariuspb.Key_KEY_AD_ID
	Key_KEY_FORMAT_ID       = archivariuspb.Key_KEY_FORMAT_ID
	Key_KEY_JUMPER_ID       = archivariuspb.Key_KEY_JUMPER_ID
	Key_KEY_CARRIER_ID      = archivariuspb.Key_KEY_CARRIER_ID
	Key_KEY_COUNTRY         = archivariuspb.Key_KEY_COUNTRY
	Key_KEY_CITY            = archivariuspb.Key_KEY_CITY
	Key_KEY_LANGUAGE        = archivariuspb.Key_KEY_LANGUAGE
	Key_KEY_IP              = archivariuspb.Key_KEY_IP
	Key_KEY_DEVICE_TYPE     = archivariuspb.Key_KEY_DEVICE_TYPE
	Key_KEY_DEVICE_ID       = archivariuspb.Key_KEY_DEVICE_ID
	Key_KEY_OS_ID           = archivariuspb.Key_KEY_OS_ID
	Key_KEY_BROWSER_ID      = archivariuspb.Key_KEY_BROWSER_ID
)

const (
	OrderingKey_UNKNOWN_ORDERING_KEY         = archivariuspb.OrderingKey_UNKNOWN_ORDERING_KEY
	OrderingKey_ORDERING_KEY_DATEMARK        = archivariuspb.OrderingKey_ORDERING_KEY_DATEMARK
	OrderingKey_ORDERING_KEY_TIMEMARK        = archivariuspb.OrderingKey_ORDERING_KEY_TIMEMARK
	OrderingKey_ORDERING_KEY_CLUSTER         = archivariuspb.OrderingKey_ORDERING_KEY_CLUSTER
	OrderingKey_ORDERING_KEY_PROJECT_ID      = archivariuspb.OrderingKey_ORDERING_KEY_PROJECT_ID
	OrderingKey_ORDERING_KEY_PUB_ACCOUNT_ID  = archivariuspb.OrderingKey_ORDERING_KEY_PUB_ACCOUNT_ID
	OrderingKey_ORDERING_KEY_ADV_ACCOUNT_ID  = archivariuspb.OrderingKey_ORDERING_KEY_ADV_ACCOUNT_ID
	OrderingKey_ORDERING_KEY_SOURCE_ID       = archivariuspb.OrderingKey_ORDERING_KEY_SOURCE_ID
	OrderingKey_ORDERING_KEY_ACCESS_POINT_ID = archivariuspb.OrderingKey_ORDERING_KEY_ACCESS_POINT_ID
	OrderingKey_ORDERING_KEY_PLATFORM        = archivariuspb.OrderingKey_ORDERING_KEY_PLATFORM
	OrderingKey_ORDERING_KEY_DOMAIN          = archivariuspb.OrderingKey_ORDERING_KEY_DOMAIN
	OrderingKey_ORDERING_KEY_APP_ID          = archivariuspb.OrderingKey_ORDERING_KEY_APP_ID
	OrderingKey_ORDERING_KEY_ZONE_ID         = archivariuspb.OrderingKey_ORDERING_KEY_ZONE_ID
	OrderingKey_ORDERING_KEY_CAMPAIGN_ID     = archivariuspb.OrderingKey_ORDERING_KEY_CAMPAIGN_ID
	OrderingKey_ORDERING_KEY_AD_ID           = archivariuspb.OrderingKey_ORDERING_KEY_AD_ID
	OrderingKey_ORDERING_KEY_FORMAT_ID       = archivariuspb.OrderingKey_ORDERING_KEY_FORMAT_ID
	OrderingKey_ORDERING_KEY_JUMPER_ID       = archivariuspb.OrderingKey_ORDERING_KEY_JUMPER_ID
	OrderingKey_ORDERING_KEY_CARRIER_ID      = archivariuspb.OrderingKey_ORDERING_KEY_CARRIER_ID
	OrderingKey_ORDERING_KEY_COUNTRY         = archivariuspb.OrderingKey_ORDERING_KEY_COUNTRY
	OrderingKey_ORDERING_KEY_CITY            = archivariuspb.OrderingKey_ORDERING_KEY_CITY
	OrderingKey_ORDERING_KEY_LANGUAGE        = archivariuspb.OrderingKey_ORDERING_KEY_LANGUAGE
	OrderingKey_ORDERING_KEY_IP              = archivariuspb.OrderingKey_ORDERING_KEY_IP
	OrderingKey_ORDERING_KEY_DEVICE_TYPE     = archivariuspb.OrderingKey_ORDERING_KEY_DEVICE_TYPE
	OrderingKey_ORDERING_KEY_DEVICE_ID       = archivariuspb.OrderingKey_ORDERING_KEY_DEVICE_ID
	OrderingKey_ORDERING_KEY_OS_ID           = archivariuspb.OrderingKey_ORDERING_KEY_OS_ID
	OrderingKey_ORDERING_KEY_BROWSER_ID      = archivariuspb.OrderingKey_ORDERING_KEY_BROWSER_ID
	OrderingKey_ORDERING_KEY_SPENT           = archivariuspb.OrderingKey_ORDERING_KEY_SPENT
	OrderingKey_ORDERING_KEY_PROFIT          = archivariuspb.OrderingKey_ORDERING_KEY_PROFIT
	OrderingKey_ORDERING_KEY_BID_PRICE       = archivariuspb.OrderingKey_ORDERING_KEY_BID_PRICE
	OrderingKey_ORDERING_KEY_REQUESTS        = archivariuspb.OrderingKey_ORDERING_KEY_REQUESTS
	OrderingKey_ORDERING_KEY_IMPRESSIONS     = archivariuspb.OrderingKey_ORDERING_KEY_IMPRESSIONS
	OrderingKey_ORDERING_KEY_VIEWS           = archivariuspb.OrderingKey_ORDERING_KEY_VIEWS
	OrderingKey_ORDERING_KEY_DIRECTS         = archivariuspb.OrderingKey_ORDERING_KEY_DIRECTS
	OrderingKey_ORDERING_KEY_CLICKS          = archivariuspb.OrderingKey_ORDERING_KEY_CLICKS
	OrderingKey_ORDERING_KEY_LEADS           = archivariuspb.OrderingKey_ORDERING_KEY_LEADS
	OrderingKey_ORDERING_KEY_BIDS            = archivariuspb.OrderingKey_ORDERING_KEY_BIDS
	OrderingKey_ORDERING_KEY_WINS            = archivariuspb.OrderingKey_ORDERING_KEY_WINS
	OrderingKey_ORDERING_KEY_SKIPS           = archivariuspb.OrderingKey_ORDERING_KEY_SKIPS
	OrderingKey_ORDERING_KEY_NOBIDS          = archivariuspb.OrderingKey_ORDERING_KEY_NOBIDS
	OrderingKey_ORDERING_KEY_ERRORS          = archivariuspb.OrderingKey_ORDERING_KEY_ERRORS
	OrderingKey_ORDERING_KEY_CTR             = archivariuspb.OrderingKey_ORDERING_KEY_CTR
	OrderingKey_ORDERING_KEY_ECPM            = archivariuspb.OrderingKey_ORDERING_KEY_ECPM
	OrderingKey_ORDERING_KEY_ECPC            = archivariuspb.OrderingKey_ORDERING_KEY_ECPC
	OrderingKey_ORDERING_KEY_ECPA            = archivariuspb.OrderingKey_ORDERING_KEY_ECPA
)

// TimeNew creates a new timestamppb.Timestamp from time.Time
func TimeNew(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

// TimeNow creates a new timestamppb.Timestamp from time.Now
func TimeNow() *timestamppb.Timestamp {
	return timestamppb.Now()
}
