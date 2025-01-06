
INSERT INTO stats.events_local (
    sign
  , timemark
  , delay              -- Delay of preparation of Ads in Nanosecinds (Load picture before display, etc)
  , duration           -- Duration in Nanosecond
  , service
  , event
  , status        -- Status: 0 - undefined, 1 - success, 2 - failed, 3 - compromised, 4 - Custom
  , merged        -- Is merged CPA (internal)
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
  , lead_price         -- Number
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

  , subid              -- SubID number
  ) SELECT
    1 AS sign
  , timemark
  , delay              -- Delay of preparation of Ads in Nanosecinds (Load picture before display, etc)
  , duration           -- Duration in Nanosecond
  , service
  , 'lead' AS event
  , 1 AS status        -- Status: 0 - undefined, 1 - success, 2 - failed, 3 - compromised, 4 - Custom
  , 1 AS merged        -- Is merged CPA (internal)
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
  , a.lead_price       -- Number
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

  , subid              -- SubID number
  FROM
  (
    SELECT aucid, lead_price FROM stats.events_local FINAL
      WHERE event = 'lead' AND datemark = today()
      GROUP BY aucid, lead_price
      HAVING SUM(merged) = 0
  ) AS a
  ALL INNER JOIN
  (
    SELECT * FROM stats.events_local FINAL WHERE event IN ('click', 'direct')
  ) AS b USING (aucid);

--
-- Separate leads table
--

INSERT INTO stats.leads_local (
    sign
  , timemark
  , service
  , merged             -- Is merged CPA
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
  , revenue            -- In percents Percent * 100 (three dementions after point)
  , potential          -- Percent of avaited descripancy (three dementions after point)

  , subid
  ) SELECT
    1 AS sign
  , timemark
  , service
  , 1 AS merged        -- Is merged CPA (internal)
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
  , lead_price AS price-- Number
  , cpmbid             -- Price as CPM value for analise
  , revenue            -- In percents Percent * 100 (three dementions after point)
  , potential          -- Percent of avaited descripancy (three dementions after point)

  , subid              -- SubID number
  FROM stats.events_local FINAL
  WHERE event = 'lead' AND merged = 1 AND datemark = today();
