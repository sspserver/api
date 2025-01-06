CREATE MATERIALIZED VIEW IF NOT EXISTS stats.v_aggregated_counters_local_live Engine = AggregatingMergeTree()
  PARTITION BY toYYYYMM(datehourmark) ORDER BY (campaign, ad, datehourmark)
AS SELECT
      MAX(timemark) AS timemark
    , datehourmark
    , datemark
    , cluster
    --- Accounts link information
    , project
    , pub_account
    , adv_account
    -- source
    , source
    , access_point
    -- State
    , platform
    , app
    , zone
    , campaign
    , ad
    , format
    , jumper
    -- Money
    , pricing_model
    , sumState(CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('view', 'direct') THEN 
          view_price * sign
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END) AS spend -- Spend of the network
    , sumState(multiIf(
        pricing_model = 1 AND status = 0 AND event IN ('view', 'direct'),
        view_price * sign, 0
      )) AS potential_spend -- Spend of the network
    , sumState(CASE
        WHEN pricing_model = 1 AND status = 2 AND event IN ('view', 'direct') THEN 
          view_price * sign
        WHEN pricing_model = 2 AND status = 2 AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status = 2 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END) AS failed_spend
    , sumState(CASE
        WHEN pricing_model = 1 AND status = 3 AND event IN ('view', 'direct') THEN 
          view_price * sign
        WHEN pricing_model = 2 AND status = 3 AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status = 3 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END) AS compromised_spend
    , sumState(CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('view', 'direct') THEN
          multiIf(purchase_view_price > 0,  purchase_view_price * sign,  CAST((view_price / 100) * (revenue / 100) AS UInt64) * sign)
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          multiIf(purchase_click_price > 0, purchase_click_price * sign, CAST((click_price / 100) * (revenue / 100) AS UInt64) * sign)
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          multiIf(purchase_lead_price > 0,  purchase_lead_price * sign,  CAST((lead_price / 100) * (revenue / 100) AS UInt64) * sign)
        ELSE 0
      END) AS revenue -- Revenue of the publisher
    , sumState(CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('view', 'direct') THEN
          multiIf(purchase_view_price > 0,  (view_price - purchase_view_price) * sign,   CAST((view_price / 100) * (potential / 100) AS UInt64) * sign)
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          multiIf(purchase_click_price > 0, (click_price - purchase_click_price) * sign, CAST((click_price / 100) * (potential / 100) AS UInt64) * sign)
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          multiIf(purchase_lead_price > 0,  (lead_price - purchase_lead_price) * sign,   CAST((lead_price / 100) * (potential / 100) AS UInt64) * sign)
        ELSE 0
      END) AS potential -- Revenue potential of publisher
    -- Counters
    , sumState(CAST(1 * sign AS UInt64)) AS total
    , sumState(CAST(CASE WHEN event = 'impression' AND status = 0   THEN 1 * sign ELSE 0 END AS UInt64)) AS imps
    , sumState(CAST(CASE WHEN event = 'impression' AND status = 1   THEN 1 * sign ELSE 0 END AS UInt64)) AS success_imps
    , sumState(CAST(CASE WHEN event = 'impression' AND status = 2   THEN 1 * sign ELSE 0 END AS UInt64)) AS failed_imps
    , sumState(CAST(CASE WHEN event = 'impression' AND status = 3   THEN 1 * sign ELSE 0 END AS UInt64)) AS compromised_imps
    , sumState(CAST(CASE WHEN event = 'impression' AND status = 4   THEN 1 * sign ELSE 0 END AS UInt64)) AS custom_imps
    , sumState(CAST(CASE WHEN event = 'impression' AND status = 5   THEN 1 * sign ELSE 0 END AS UInt64)) AS failover_imps
    , sumState(CAST(CASE WHEN event = 'view' AND status IN (0, 1)   THEN 1 * sign ELSE 0 END AS UInt64)) AS views
    , sumState(CAST(CASE WHEN event = 'view' AND status = 2         THEN 1 * sign ELSE 0 END AS UInt64)) AS failed_views
    , sumState(CAST(CASE WHEN event = 'view' AND status = 3         THEN 1 * sign ELSE 0 END AS UInt64)) AS compromised_views
    , sumState(CAST(CASE WHEN event = 'view' AND status = 4         THEN 1 * sign ELSE 0 END AS UInt64)) AS custom_views
    , sumState(CAST(CASE WHEN event = 'view' AND status = 5         THEN 1 * sign ELSE 0 END AS UInt64)) AS failover_views
    , sumState(CAST(CASE WHEN event = 'direct' AND status = 0       THEN 1 * sign ELSE 0 END AS UInt64)) AS directs
    , sumState(CAST(CASE WHEN event = 'direct' AND status = 1       THEN 1 * sign ELSE 0 END AS UInt64)) AS success_directs
    , sumState(CAST(CASE WHEN event = 'direct' AND status = 2       THEN 1 * sign ELSE 0 END AS UInt64)) AS failed_directs
    , sumState(CAST(CASE WHEN event = 'direct' AND status = 3       THEN 1 * sign ELSE 0 END AS UInt64)) AS compromised_directs
    , sumState(CAST(CASE WHEN event = 'direct' AND status = 4       THEN 1 * sign ELSE 0 END AS UInt64)) AS custom_directs
    , sumState(CAST(CASE WHEN event = 'direct' AND status = 5       THEN 1 * sign ELSE 0 END AS UInt64)) AS failover_directs
    , sumState(CAST(CASE WHEN event = 'click'  AND status IN (0, 1) THEN 1 * sign ELSE 0 END AS UInt64)) AS clicks
    , sumState(CAST(CASE WHEN event = 'click'  AND status = 2       THEN 1 * sign ELSE 0 END AS UInt64)) AS failed_clicks
    , sumState(CAST(CASE WHEN event = 'click'  AND status = 3       THEN 1 * sign ELSE 0 END AS UInt64)) AS compromised_clicks
    , sumState(CAST(CASE WHEN event = 'click'  AND status = 4       THEN 1 * sign ELSE 0 END AS UInt64)) AS custom_clicks
    , sumState(CAST(CASE WHEN event = 'click'  AND status = 5       THEN 1 * sign ELSE 0 END AS UInt64)) AS failover_clicks
    , sumState(CAST(CASE WHEN event = 'lead'   AND status IN (0, 1) THEN 1 * sign ELSE 0 END AS UInt64)) AS leads
    , sumState(CAST(CASE WHEN event = 'lead'   AND status = 2       THEN 1 * sign ELSE 0 END AS UInt64)) AS failed_leads
    , sumState(CAST(CASE WHEN event = 'lead'   AND status = 3       THEN 1 * sign ELSE 0 END AS UInt64)) AS compromised_leads
    , sumState(CAST(CASE WHEN adblock = 1                           THEN 1 * sign ELSE 0 END AS UInt64)) AS adblocks
    , sumState(CAST(CASE WHEN private = 1                           THEN 1 * sign ELSE 0 END AS UInt64)) AS privates
    , sumState(CAST(CASE WHEN robot = 1                             THEN 1 * sign ELSE 0 END AS UInt64)) AS robots
    , sumState(CAST(CASE WHEN backup = 1                            THEN 1 * sign ELSE 0 END AS UInt64)) AS backups
FROM stats.events_local
GROUP BY (
      datehourmark
    , datemark
    , cluster
    --- Accounts link information
    , project
    , pub_account
    , adv_account
    -- source
    , source
    , access_point
    -- State
    , platform
    , app
    , zone
    , campaign
    , ad
    , format
    , jumper
    -- Money
    , pricing_model
);
