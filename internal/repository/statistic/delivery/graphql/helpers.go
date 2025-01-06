package graphql

import (
	"github.com/geniusrabbit/archivarius/client"

	"github.com/sspserver/api/internal/server/graphql/models"
)

// itemFromPb convert statistic item from client item
func itemFromPb(item *client.Item) (*models.StatisticAdItem, error) {
	var keys []*models.StatisticItemKey

	for _, k := range item.Keys {
		keys = append(keys, &models.StatisticItemKey{
			Key:   keyFromPb(k.Key),
			Value: client.ValueExtFrom(k.Value.Value),
		})
	}

	return &models.StatisticAdItem{
		Keys:        keys,
		Spent:       item.Spent,
		Profit:      item.Profit,
		BidPrice:    item.BidPrice,
		Requests:    item.Requests,
		Impressions: item.Impressions,
		Views:       item.Views,
		Directs:     item.Directs,
		Clicks:      item.Clicks,
		Leads:       item.Leads,
		Bids:        item.Bids,
		Wins:        item.Wins,
		Skips:       item.Skips,
		Nobids:      item.Nobids,
		Errors:      item.Errors,
		Ctr:         item.Ctr,
		ECpm:        item.Ecpm,
		ECpc:        item.Ecpc,
		ECpa:        item.Ecpa,
	}, nil
}

// keyFromPb convert client key to key
func keyFromPb(key client.Key) models.StatisticKey {
	switch key {
	case client.Key_KEY_DATEMARK:
		return models.StatisticKeyDatemark
	case client.Key_KEY_TIMEMARK:
		return models.StatisticKeyTimemark
	case client.Key_KEY_CLUSTER:
		return models.StatisticKeyCluster
	case client.Key_KEY_PROJECT_ID:
		return models.StatisticKeyProjectID
	case client.Key_KEY_PUB_ACCOUNT_ID:
		return models.StatisticKeyPubAccountID
	case client.Key_KEY_ADV_ACCOUNT_ID:
		return models.StatisticKeyAdvAccountID
	case client.Key_KEY_SOURCE_ID:
		return models.StatisticKeySourceID
	case client.Key_KEY_ACCESS_POINT_ID:
		return models.StatisticKeyAccessPointID
	case client.Key_KEY_PLATFORM:
		return models.StatisticKeyPlatform
	case client.Key_KEY_DOMAIN:
		return models.StatisticKeyDomain
	case client.Key_KEY_APP_ID:
		return models.StatisticKeyAppID
	case client.Key_KEY_ZONE_ID:
		return models.StatisticKeyZoneID
	case client.Key_KEY_CAMPAIGN_ID:
		return models.StatisticKeyCampaignID
	case client.Key_KEY_AD_ID:
		return models.StatisticKeyAdID
	case client.Key_KEY_FORMAT_ID:
		return models.StatisticKeyFormatID
	case client.Key_KEY_JUMPER_ID:
		return models.StatisticKeyJumperID
	case client.Key_KEY_CARRIER_ID:
		return models.StatisticKeyCarrierID
	case client.Key_KEY_COUNTRY:
		return models.StatisticKeyCountry
	case client.Key_KEY_CITY:
		return models.StatisticKeyCity
	case client.Key_KEY_LANGUAGE:
		return models.StatisticKeyLanguage
	case client.Key_KEY_IP:
		return models.StatisticKeyIP
	case client.Key_KEY_DEVICE_TYPE:
		return models.StatisticKeyDeviceType
	case client.Key_KEY_DEVICE_ID:
		return models.StatisticKeyDeviceID
	case client.Key_KEY_OS_ID:
		return models.StatisticKeyOsID
	case client.Key_KEY_BROWSER_ID:
		return models.StatisticKeyBrowserID
	}

	panic("unknown key")
}
