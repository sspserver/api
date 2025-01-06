
CREATE TABLE IF NOT EXISTS stats.leads_local (
   sign               Int8        default 1
 , timemark           DateTime
 , datehourmark       DateTime    default toStartOfHour(timemark)
 , datemark           Date        default toDate(timemark)
 , service            String
 , merged             Int8        default 0       -- Is merged CPA
 -- source
 , aucid              FixedString(16)             -- Internal Auction ID
 , impid              FixedString(16)             -- Sub ID of request for paticular impression spot
 , impadid            FixedString(16)             -- Specific ID for paticular ad impression
 , extaucid           String                      -- External Request/Response ID (RTB)
 , extimpid           String                      -- External Imp ID (RTB)
 , source             UInt64                      -- Advertisement Source ID
 , network            String                      -- Source Network Name or Domain (Cross sails)
 , access_point       UInt64                      -- Access Point ID to own Advertisement
 -- State Location
 , project            UInt64                      -- Project ID
 , pub_account        UInt64                      -- Publisher Company ID
 , adv_account        UInt64                      -- Advertiser Company ID
 , platform           UInt8                       -- Where displaid? 0 – undefined, 1 – web site, 2 – native app, 3 – game, 4 - Video Player
 , domain             String                      -- If not web site then "bundle"
 , app                UInt64                      -- application ID (registered in the system)
 , zone               UInt64                      -- application Zone ID
 , campaign           UInt64                      -- campaign ID
 , src_url            String                      -- Advertisement source URL (iframe, image, video, direct)
 , click_url          String                      -- Click target URL
 , win_url            String                      -- Win URL used for RTB confirmation
 , url                String                      -- Non modified target URL
 , ad                 UInt64                      -- Advertisement ID
 , ad_w               UInt32                      -- Area Width
 , ad_h               UInt32                      -- Area Height
 , jumper             UInt64                      -- Jumper Page ID
 , ad_type            UInt8                       -- Specified containt type
 -- Money
 , pricing_model      UInt8                       -- Display As CPM/CPC/CPA/CPI
 , price              UInt64                      -- Number
 , cpmbid             UInt64                      -- Price as CPM value for analise
 , revenue            UInt64                      -- In percents Percent * 100 (three dementions after point)
 , potential          UInt64                      -- Percent of avaited descripancy (three dementions after point)

 , subid              UInt32      default 0
 , created_at         DateTime    default now()
)
ENGINE = ReplicatedCollapsingMergeTree(sign)
PARTITION BY toYYYYMM(datemark)
ORDER BY (aucid, datemark, service, pub_account, adv_account, zone, impadid)
SETTINGS index_granularity = 8192;
