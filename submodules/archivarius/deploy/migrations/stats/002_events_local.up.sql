
-- Money is Number * 1_000_000_000

-- TODO: complete review of the table and refactor it
CREATE TABLE IF NOT EXISTS stats.events_local (
   sign                   Int8        DEFAULT 1
 , timemark               DateTime
 , datehourmark           DateTime    DEFAULT toStartOfHour(timemark)
 , datemark               Date        DEFAULT toDate(timemark)
 , delay                  UInt64      DEFAULT 0       -- Delay of preparation of Ads in Nanosecinds (Load picture before display, etc)
 , duration               UInt64      DEFAULT 0       -- Duration in Nanosecond
 , service                String
 , cluster                FixedString(2)
 , event                  LowCardinality(String)
 , status                 UInt8                       -- Status: 0 - undefined, 1 - success, 2 - failed, 3 - compromised, 4 - Custom, 5 – Failover
 , merged                 Int8        DEFAULT 0       -- Is merged CPA (internal)
 --- Accounts link information
 , project                UInt64                      -- Project ID
 , pub_account            UInt64                      -- Publisher Account ID
 , adv_account            UInt64                      -- Advertiser Account ID
 -- source
 , aucid                  FixedString(16)             -- Internal Auction ID
 , auctype                UInt8                       -- 1 - First Price, 2 - Second Price, 3 - Fixed Price, 0 - Default (Second Price)
 , impid                  FixedString(16)             -- Sub ID of request for paticular impression spot
 , impadid                FixedString(16)             -- Specific ID for paticular ad impression
 , extaucid               String                      -- External Request/Response ID (RTB)
 , extimpid               String                      -- External Imp ID (RTB)
 , extzoneid              String                      -- External zone ID
 , source                 UInt64                      -- Advertisement Source ID
 , network                String                      -- Source Network Name or Domain (Cross sails)
 , access_point           UInt64                      -- Access Point ID to own Advertisement
 -- State Location
 , platform               UInt8                       -- Where displaid? 0 – undefined, 1 – web site, 2 – native app, 3 – game, 4 - Video Player
 , domain                 String                      -- If not web site then "bundle"
 , app                    UInt64                      -- application ID (registered in the system)
 , zone                   UInt64                      -- application Zone ID
 , pixel                  UInt64                      -- Pixel ID
 , campaign               UInt64                      -- campaign ID
 , format                 UInt32                      -- Advertisement format ID
 , ad                     UInt64                      -- Advertisement ID
 , ad_w                   UInt32                      -- Area Width
 , ad_h                   UInt32                      -- Area Height
 , src_url                String                      -- Advertisement source URL (iframe, image, video, direct)
 , win_url                String                      -- Win URL used for RTB confirmation
 , url                    String                      -- Non modified target URL
 , jumper                 UInt64                      -- Jumper Page ID
 -- Money
 , pricing_model          UInt8                     -- Display As CPM/CPC/CPA/CPI
 , purchase_view_price    UInt64                    -- Money paid to the source
 , purchase_click_price   UInt64
 , purchase_lead_price    UInt64
 , potential_view_price   UInt64                    -- Additional price which can we have
 , potential_click_price  UInt64
 , potential_lead_price   UInt64
 , view_price             UInt64                    -- Total price with all expencies per action
 , click_price            UInt64
 , lead_price             UInt64
 , competitor             UInt64                    -- Competitor compaign ID
 , competitor_source      UInt64                    -- Competitor source ID
 , competitor_ecpm        UInt64                    -- Second price in CPM
 -- User IDENTITY
 , udid                   String                      -- Unique Device ID (IDFA)
 , uuid                   FixedString(16)
 , sessid                 FixedString(16)
 , fingerprint            String
 , etag                   String
 -- Targeting
 , carrier                UInt64                      -- Carrier ID
 , country                FixedString(2)              -- Country Code
 , city                   String                      -- City Code
 , latitude               Float64                     -- Geo latitude
 , longitude              Float64                     -- Geo longitude
 , language               FixedString(5)              -- en-US
 , ip                     IPv6
 , ref                    String                      -- Referal link
 , page_url               String                      -- Page link
 , ua                     String                      -- User Agent
 , device_type            UInt32                      -- Device type 0 - Undefined, 1 - Desktop, etc.
 , device                 UInt32                      -- Device ID
 , os                     UInt32                      -- OS ID
 , browser                UInt32                      -- Browser ID
 , categories             Array(Int32)                -- Categories list
 , adblock                UInt8
 , private                UInt8                       -- Private Mode
 , robot                  UInt8                       -- Robot traffic
 , proxy                  UInt8                       -- Proxy traffic
 , backup                 UInt8                       -- Backup Display Type
 , x                      Int32                       -- X - coord of addisplay or click position
 , y                      Int32                       -- Y - coord of addisplay or click position
 , w                      Int32                       -- W - available space
 , h                      Int32                       -- H - available space

 , subid1                 String
 , subid2                 String
 , subid3                 String
 , subid4                 String
 , subid5                 String

 , created_at             DateTime    DEFAULT now()
)
ENGINE = ReplicatedCollapsingMergeTree(sign)
PARTITION BY toYYYYMM(datemark)
ORDER BY (event, datemark, service, project, pub_account, adv_account, zone, format, aucid, impadid)
SETTINGS index_granularity = 8192;
