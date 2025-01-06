INSERT INTO stats.events_local2 (
   sign
 , timemark
 , delay              -- Delay of preparation of Ads in Nanosecinds (Load picture before display, etc)
 , duration           -- Duration in Nanosecond
 , service
 , event
 , status             -- Status: 0 - undefined, 1 - success, 2 - failed, 3 - compromised, 4 - Custom
 , merged             -- Is merged CPA (internal)
 -- source
 , aucid              -- Internal Auction ID
 , impid              -- Sub ID of request for paticular impression spot
 , impadid            -- Specific ID for paticular ad impression
 , extaucid           -- External Request/Response ID (RTB)
 , extimpid           -- External Imp ID (RTB)
 , source             -- Advertisement Source ID
 , network            -- Source Network Name or Domain (Cross sails)
 , access_point       -- Access Point ID to own Advertisement
 -- State Location
 , project            -- Project ID
 , pub_account        -- Publisher Company ID
 , adv_account        -- Advertiser Company ID
 , platform           -- Where displaid? 0 – undefined, 1 – web site, 2 – native app, 3 – game, 4 - Video Player
 , domain             -- If not web site then "bundle"
 , app                -- application ID (registered in the system)
 , zone               -- application Zone ID
 , campaign           -- campaign ID
 , src_url            -- Advertisement source URL (iframe, image, video, direct)
 , click_url          -- Click target URL
 , win_url            -- Win URL used for RTB confirmation
 , url                -- Non modified target URL
 , ad                 -- Advertisement ID
 , ad_w               -- Area Width
 , ad_h               -- Area Height
 , jumper             -- Jumper Page ID
 , ad_type            -- Specified containt type
 -- Money
 , pricing_model      -- Display As CPM/CPC/CPA/CPI
 , price              -- Number
 , cpmbid             -- Price as CPM value for analise
 , competitor         -- Competitor compaign ID
 , competitor_source  -- Competitor source ID
 , competitor_cpmbid  -- Second price in CPM
 , revenue            -- In percents Percent * 100 (three dementions after point)
 , potential          -- Percent of avaited descripancy (three dementions after point)
 -- User IDENTITY
 , udid               -- Unique Device ID (IDFA)
 , uuid
 , sessid
 , fingerprint
 , etag
 -- Targeting
 , carrier            -- Carrier ID
 , country            -- Country Code
 , city               -- City Code
 , latitude           -- Geo latitude
 , longitude          -- Geo longitude
 , language           -- en-US
 , ip
 , ref                -- Referal link
 , page               -- Page link
 , ua                 -- User Agent
 , device_type        -- Device type 0 - Undefined, 1 - Desktop, etc.
 , device             -- Device ID
 , os                 -- OS ID
 , browser            -- Browser ID
 , categories         -- Categories list
 , adblock
 , private            -- Private Mode
 , robot              -- Robot traffic
 , proxy              -- Proxy traffic
 , backup             -- Backup Display Type
 , x                  -- X - coord of addisplay or click position
 , y                  -- Y - coord of addisplay or click position
 , w                  -- W - available space
 , h                  -- H - available space

 , subid              -- SubID number
 , created_at
) SELECT
   sign
 , timemark
 , delay              -- Delay of preparation of Ads in Nanosecinds (Load picture before display, etc)
 , duration           -- Duration in Nanosecond
 , service
 , event
 , status             -- Status: 0 - undefined, 1 - success, 2 - failed, 3 - compromised, 4 - Custom
 , merged             -- Is merged CPA (internal)
 -- source
 , aucid              -- Internal Auction ID
 , impid              -- Sub ID of request for paticular impression spot
 , impadid            -- Specific ID for paticular ad impression
 , extaucid           -- External Request/Response ID (RTB)
 , extimpid           -- External Imp ID (RTB)
 , source             -- Advertisement Source ID
 , network            -- Source Network Name or Domain (Cross sails)
 , access_point       -- Access Point ID to own Advertisement
 -- State Location
 , project            -- Project ID
 , pub_account        -- Publisher Company ID
 , adv_account        -- Advertiser Company ID
 , platform           -- Where displaid? 0 – undefined, 1 – web site, 2 – native app, 3 – game, 4 - Video Player
 , domain             -- If not web site then "bundle"
 , app                -- application ID (registered in the system)
 , zone               -- application Zone ID
 , campaign           -- campaign ID
 , src_url            -- Advertisement source URL (iframe, image, video, direct)
 , click_url          -- Click target URL
 , win_url            -- Win URL used for RTB confirmation
 , url                -- Non modified target URL
 , ad                 -- Advertisement ID
 , ad_w               -- Area Width
 , ad_h               -- Area Height
 , jumper             -- Jumper Page ID
 , ad_type            -- Specified containt type
 -- Money
 , pricing_model      -- Display As CPM/CPC/CPA/CPI
 , price              -- Number
 , cpmbid             -- Price as CPM value for analise
 , competitor         -- Competitor compaign ID
 , competitor_source  -- Competitor source ID
 , competitor_cpmbid  -- Second price in CPM
 , revenue            -- In percents Percent * 100 (three dementions after point)
 , potential          -- Percent of avaited descripancy (three dementions after point)
 -- User IDENTITY
 , udid               -- Unique Device ID (IDFA)
 , uuid
 , sessid
 , fingerprint
 , etag
 -- Targeting
 , carrier            -- Carrier ID
 , country            -- Country Code
 , city               -- City Code
 , latitude           -- Geo latitude
 , longitude          -- Geo longitude
 , language           -- en-US
 , ip
 , ref                -- Referal link
 , page               -- Page link
 , ua                 -- User Agent
 , device_type        -- Device type 0 - Undefined, 1 - Desktop, etc.
 , device             -- Device ID
 , os                 -- OS ID
 , browser            -- Browser ID
 , categories         -- Categories list
 , adblock
 , private            -- Private Mode
 , robot              -- Robot traffic
 , proxy              -- Proxy traffic
 , backup             -- Backup Display Type
 , x                  -- X - coord of addisplay or click position
 , y                  -- Y - coord of addisplay or click position
 , w                  -- W - available space
 , h                  -- H - available space

 , subid              -- SubID number
 , created_at
FROM stats.events_local;

RENAME TABLE stats.events_local TO stats.events_local_x2;
RENAME TABLE stats.events_local2 TO stats.events_local;
