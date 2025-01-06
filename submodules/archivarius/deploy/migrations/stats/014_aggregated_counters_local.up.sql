CREATE TABLE IF NOT EXISTS stats.aggregated_counters_local (
  t                        DateTime                                          TTL t + INTERVAL 1 DAY
, time5min_mark            DateTime    DEFAULT toStartOfFiveMinutes(t)       TTL time5min_mark  + INTERVAL 3 DAYS
, time15min_mark           DateTime    DEFAULT toStartOfFifteenMinutes(t)    TTL time15min_mark + INTERVAL 10 DAYS
, timehour_mark            DateTime    DEFAULT toStartOfHour(t)              TTL timehour_mark  + INTERVAL 1 MONTH
, datemark                 Date        DEFAULT toDate(t)
, timemark                 DateTime    ALIAS GREATEST(t, time5min_mark, time15min_mark, timehour_mark, toDateTime(datemark))
, cluster                  FixedString(2)
-- Accounts link information
, project_id               UInt64                    -- Project ID
, pub_account_id           UInt64                    -- Publisher Account ID
, adv_account_id           UInt64                    -- Advertiser Account ID
-- source
, source_id                UInt64                    -- Advertisement Source ID
, access_point_id          UInt64                    -- Access Point ID to own Advertisement
-- Targeting
, platform                 UInt8                     -- Where displaid? 0 – undefined, 1 – web site, 2 – native app, 3 – game, 4 - Targeting
, domain                   String                    -- If not web site then "bundle"
, app_id                   UInt64                    -- application ID (registered in the system)
, zone_id                  UInt64                    -- application Zone ID
, campaign_id              UInt64                    -- campaign ID
, ad_id                    UInt64                    -- Advertisement ID
, format_id                UInt32                    -- Advertisement format ID
, jumper_id                UInt64                    -- Jumper Page ID
-- Wide targeting information
, carrier_id               UInt64                    -- Carrier ID
, country                  FixedString(2)            -- Country Code
, city                     String                    -- City Code
, latitude                 Float64                   -- Geo latitude
, longitude                Float64                   -- Geo longitude
, language                 FixedString(5)            -- en-US
, ip                       IPv6
, device_type              UInt32                    -- Device type 0 - Undefined, 1 - Desktop, etc.
, device_id                UInt32                    -- Device ID
, os_id                    UInt32                    -- OS ID
, browser_id               UInt32                    -- Browser ID
-- Money
, pricing_model            UInt8                     -- Display As CPM/CPC/CPA/CPI
, adv_spend                UInt64
, adv_potential_spend      UInt64
, adv_failed_spend         UInt64
, adv_compromised_spend    UInt64
, pub_revenue              UInt64
, sales_budget             UInt64                   -- Sales budget
, buying_budget            UInt64                   -- Buying budget
, network_revenue          UInt64    ALIAS adv_spend - pub_revenue
-- Counters
, imps                     UInt64
, success_imps             UInt64
, failed_imps              UInt64
, compromised_imps         UInt64
, custom_imps              UInt64
, backup_imps              UInt64
, views                    UInt64
, failed_views             UInt64
, compromised_views        UInt64
, custom_views             UInt64
, backup_views             UInt64
, directs                  UInt64
, success_directs          UInt64
, failed_directs           UInt64
, compromised_directs      UInt64
, custom_directs           UInt64
, backup_directs           UInt64
, clicks                   UInt64
, failed_clicks            UInt64
, compromised_clicks       UInt64
, custom_clicks            UInt64
, backup_clicks            UInt64
, leads                    UInt64
, success_leads            UInt64
, failed_leads             UInt64
, compromised_leads        UInt64

, src_bid_requests         UInt64
, src_bid_wins             UInt64
, src_bid_skips            UInt64
, src_bid_nobids           UInt64
, src_bid_errors           UInt64
, ap_bid_requests          UInt64
, ap_bid_wins              UInt64
, ap_bid_skips             UInt64
, ap_bid_nobids            UInt64
, ap_bid_errors            UInt64

, adblocks                 UInt64
, privates                 UInt64
, robots                   UInt64
, backups                  UInt64
) ENGINE = ReplicatedSummingMergeTree()
PARTITION BY toYYYYMM(datemark)
PRIMARY KEY (datemark, cluster, project_id, pub_account_id, adv_account_id, source_id, access_point_id)
ORDER BY (
  datemark
, cluster
-- Accounts link information
, project_id
, pub_account_id
, adv_account_id
-- source
, source_id
, access_point_id
-- Targeting
, platform
, domain
, app_id
, zone_id
, campaign_id
, ad_id
, format_id
, jumper_id
-- Wide targeting information
, carrier_id
, country
, city
, latitude
, longitude
, language
, ip
, device_type
, device_id
, os_id
, browser_id
-- Money
, pricing_model
)
SETTINGS index_granularity = 8192;
