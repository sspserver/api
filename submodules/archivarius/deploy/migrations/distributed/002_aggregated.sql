
CREATE TABLE IF NOT EXISTS distributed.aggregated (
  datehourmark            DateTime
, datemark                Date
, impression_delay        UInt64
, click_delay             UInt64
, duration                UInt64
, service                 String
-- source
, source                  UInt64            -- Advertisement Source ID
, network                 String            -- Source Network Name or Domain (Cross sails)
, access_point            UInt64            -- Access Point ID to own Advertisement
-- State Location
, project                 UInt64            -- Project ID
, pub_account             UInt64            -- Publisher Company ID
, adv_account             UInt64            -- Advertiser Company ID
, platform                UInt8             -- Where displaid? 0 – undefined, 1 – web site, 2 – native app, 3 – game, 4 - Video Player
, domain                  String            -- If not web site then "bundle"
, app                     UInt64            -- application ID (registered in the system)
, zone                    UInt64            -- application Zone ID
, campaign                UInt64            -- campaign ID
, ad                      UInt64            -- Advertisement ID
, url                     String            -- Non modified target URL
, jumper                  UInt64            -- Jumper Page ID
, ad_type                 UInt8             -- Specified containt type
-- Money
, pricing_model           UInt8                     -- Display As CPM/CPC/CPA/CPI
, purchase_view_price     UInt64                    -- Money paid to the source
, purchase_click_price    UInt64
, purchase_lead_price     UInt64
, potential_view_price    UInt64                    -- Additional price which can we have
, potential_click_price   UInt64
, potential_lead_price    UInt64
, view_price              UInt64                    -- Total price with all expencies per action
, click_price             UInt64
, lead_price              UInt64
, competitor              UInt64                    -- Competitor compaign ID
, competitor_source       UInt64                    -- Competitor source ID
, competitor_ecpm         UInt64                    -- Second price in CPM
-- Counters
, total                   UInt64
, imps                    UInt64            -- Count of Impressions
, success_imps            UInt64            -- Count of Success Impressions
, failed_imps             UInt64            -- Count of Failed Impressions
, compromised_imps        UInt64            -- Count of Compromised Impressions
, custom_imps             UInt64
, failover_imps           UInt64
, views                   UInt64
, failed_views            UInt64
, compromised_views       UInt64
, custom_views            UInt64
, failover_views          UInt64
, directs                 UInt64
, failed_directs          UInt64
, compromised_directs     UInt64
, custom_directs          UInt64
, failover_directs        UInt64
, clicks                  UInt64
, failed_clicks           UInt64
, compromised_clicks      UInt64
, custom_clicks           UInt64
, failover_clicks         UInt64
, leads                   UInt64
, failed_leads            UInt64
, compromised_leads       UInt64
, adblocks                UInt64
, privates                UInt64
, robots                  UInt64
, backups                 UInt64
-- Targeting
, carrier                 UInt64                      -- Carrier ID
, country                 FixedString(2)              -- Country Code
, city                    String                      -- City Code
, latitude                Float64                     -- Geo latitude
, longitude               Float64                     -- Geo longitude
, language                FixedString(5)              -- en-US
, ip                      IPv6
, ref                     String                      -- Referal link
, page_url                String                      -- Page link
, ua                      String                      -- User Agent
, device_type             UInt32                      -- Device type 0 - Undefined, 1 - Desktop, etc.
, device                  UInt32                      -- Device ID
, os                      UInt32                      -- OS ID
, browser                 UInt32                      -- Browser ID
, categories              Array(Int32)                -- Categories list
, adblock                 UInt8
, private                 UInt8                       -- Private Mode
, robot                   UInt8                       -- Robot traffic
, proxy                   UInt8                       -- Proxy traffic
, backup                  UInt8                       -- Backup Display Type
, x                       Int32                       -- X - coord of addisplay or click position
, y                       Int32                       -- Y - coord of addisplay or click position
, w                       Int32                       -- W - available space
, h                       Int32                       -- H - available space

, subid1                  String
, subid2                  String
, subid3                  String
, subid4                  String
, subid5                  String

, subid                   UInt32      default 0
) Engine=Merge(stats, '^(v_)?aggregated_');
