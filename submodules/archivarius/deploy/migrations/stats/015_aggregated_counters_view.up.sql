CREATE MATERIALIZED VIEW IF NOT EXISTS stats.v_aggregated_counters_view TO stats.aggregated_counters_local
  AS SELECT
      timemark AS t
    , cluster AS cluster
    -- Accounts link information
    , project AS project_id
    , pub_account AS pub_account_id
    , adv_account AS adv_account_id
    -- Sources/Targets
    , source AS source_id
    , access_point AS access_point_id
    -- Targeting
    , platform
    , domain
    , app AS app_id
    , zone AS zone_id
    , campaign AS campaign_id
    , ad AS ad_id
    , format AS format_id
    , jumper AS jumper_id
    -- Wide targeting information
    , carrier AS carrier_id
    , country
    , city
    , latitude
    , longitude
    , language
    , ip
    , device_type
    , device AS device_id
    , os AS os_id
    , browser AS browser_id
    -- Money
    , pricing_model
    , CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('view', 'direct') THEN
          view_price * sign
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END AS adv_spend
    , CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('view', 'direct') THEN
          potential_view_price * sign
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          potential_click_price * sign
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          potential_lead_price * sign
        ELSE 0
      END AS adv_potential_spend
    , CASE
        WHEN pricing_model = 1 AND status = 2 AND event IN ('impression', 'direct') THEN
          view_price * sign
        WHEN pricing_model = 2 AND status = 2 AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status = 2 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END AS adv_failed_spend
    , CASE
        WHEN pricing_model = 1 AND status = 3 AND event IN ('impression', 'direct') THEN
          view_price * sign
        WHEN pricing_model = 2 AND status = 3 AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status = 3 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END AS adv_compromised_spend
    , CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('view', 'direct') THEN
          purchase_view_price * sign
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          purchase_click_price * sign
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          purchase_lead_price * sign
        ELSE 0
      END AS pub_revenue
    , IF(event = 'ap.bid' , view_price * sign, 0) AS sales_budget
    , IF(event = 'src.bid', view_price * sign, 0) AS buying_budget

    -- Counters
    , CAST(IF(event = 'impression'               , sign, 0) AS UInt64) AS imps
    , CAST(IF(event = 'impression' AND status = 1, sign, 0) AS UInt64) AS success_imps
    , CAST(IF(event = 'impression' AND status = 2, sign, 0) AS UInt64) AS failed_imps
    , CAST(IF(event = 'impression' AND status = 3, sign, 0) AS UInt64) AS compromised_imps
    -- When display custom advertisement in case if no ads with required conditions
    , CAST(IF(event = 'impression' AND status = 4, sign, 0) AS UInt64) AS custom_imps
    , CAST(IF(event = 'impression' AND backup = 1, sign, 0) AS UInt64) AS backup_imps
    , CAST(IF(event = 'view'                     , sign, 0) AS UInt64) AS views
    , CAST(IF(event = 'view'       AND status = 1, sign, 0) AS UInt64) AS success_views
    , CAST(IF(event = 'view'       AND status = 2, sign, 0) AS UInt64) AS failed_views
    , CAST(IF(event = 'view'       AND status = 3, sign, 0) AS UInt64) AS compromised_views
    , CAST(IF(event = 'view'       AND status = 4, sign, 0) AS UInt64) AS custom_views
    , CAST(IF(event = 'view'       AND backup = 1, sign, 0) AS UInt64) AS backup_views
    , CAST(IF(event = 'direct'                   , sign, 0) AS UInt64) AS directs
    , CAST(IF(event = 'direct'     AND status = 1, sign, 0) AS UInt64) AS success_directs
    , CAST(IF(event = 'direct'     AND status = 2, sign, 0) AS UInt64) AS failed_directs
    , CAST(IF(event = 'direct'     AND status = 3, sign, 0) AS UInt64) AS compromised_directs
    , CAST(IF(event = 'direct'     AND status = 4, sign, 0) AS UInt64) AS custom_directs
    , CAST(IF(event = 'direct'     AND backup = 1, sign, 0) AS UInt64) AS backup_directs
    , CAST(IF(event = 'click'                    , sign, 0) AS UInt64) AS clicks
    , CAST(IF(event = 'click'      AND status = 1, sign, 0) AS UInt64) AS success_clicks
    , CAST(IF(event = 'click'      AND status = 2, sign, 0) AS UInt64) AS failed_clicks
    , CAST(IF(event = 'click'      AND status = 3, sign, 0) AS UInt64) AS compromised_clicks
    , CAST(IF(event = 'click'      AND status = 4, sign, 0) AS UInt64) AS custom_clicks
    , CAST(IF(event = 'click'      AND backup = 1, sign, 0) AS UInt64) AS backup_clicks
    , CAST(IF(event = 'lead'                     , sign, 0) AS UInt64) AS leads
    , CAST(IF(event = 'lead'       AND status = 1, sign, 0) AS UInt64) AS success_leads
    , CAST(IF(event = 'lead'       AND status = 2, sign, 0) AS UInt64) AS failed_leads
    , CAST(IF(event = 'lead'       AND status = 3, sign, 0) AS UInt64) AS compromised_leads

    , CAST(IF(event = 'src.bid'                  , sign, 0) AS UInt64) AS src_bid_requests
    , CAST(IF(event = 'src.win'                  , sign, 0) AS UInt64) AS src_bid_wins
    , CAST(IF(event = 'src.skip'                 , sign, 0) AS UInt64) AS src_bid_skips
    , CAST(IF(event = 'src.nobid'                , sign, 0) AS UInt64) AS src_bid_nobids
    , CAST(IF(event = 'src.fail'                 , sign, 0) AS UInt64) AS src_bid_errors
    , CAST(IF(event = 'ap.bid'                   , sign, 0) AS UInt64) AS ap_bid_requests
    , CAST(IF(event = 'ap.win'                   , sign, 0) AS UInt64) AS ap_bid_wins
    , CAST(IF(event = 'ap.skip'                  , sign, 0) AS UInt64) AS ap_bid_skips
    , CAST(IF(event = 'ap.nobid'                 , sign, 0) AS UInt64) AS ap_bid_nobids
    , CAST(IF(event = 'ap.fail'                  , sign, 0) AS UInt64) AS ap_bid_errors

    , CAST(IF(adblock = 1, sign, 0) AS UInt64) AS adblocks
    , CAST(IF(private = 1, sign, 0) AS UInt64) AS privates
    , CAST(IF(robot   = 1, sign, 0) AS UInt64) AS robots
    , CAST(IF(backup  = 1, sign, 0) AS UInt64) AS backups
FROM stats.events_local;
