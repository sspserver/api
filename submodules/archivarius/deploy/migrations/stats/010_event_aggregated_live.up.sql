
CREATE MATERIALIZED VIEW IF NOT EXISTS stats.v_event_aggregated_live
ENGINE = ReplicatedSummingMergeTree(
  '/clickhouse/tables/{shard}/stats/v_event_aggregated_live',
  '{replica}',
  (
    delay,
    duration,
    revenue,
    failed_revenue,
    compromised_revenue,
    expenses,
    potential,
    earnings,
    -- Counters
    total,
    count,
    success,
    failed,
    compromised,
    custom,
    failovers,
    adblocks,
    privates,
    robots,
    backups
  )
)
PARTITION BY toYYYYMM(datemark)
ORDER BY (
    datehourmark
  , datemark
  , event
  , service
  , cluster
  --- Accounts link information
  , project
  , pub_account
  , adv_account
  -- source
  , source
  , network
  , access_point
  -- State Location
  , platform
  , domain
  , app
  , zone
  , campaign
  , ad
  , format
  , url
  , jumper
  -- Money
  , pricing_model
  -- Targeting
  , carrier
  , country
  , city
  , language
  , device_type
  , device
  , os
  , browser
  , categories
)
SETTINGS index_granularity = 8192
AS 
SELECT 
      datehourmark
    , datemark
    , delay
    , duration
    , event
    , service
    , cluster
    --- Accounts link information
    , project
    , pub_account
    , adv_account
    -- source
    , source
    , network
    , access_point
    -- State Location
    , platform
    , domain
    , app
    , zone
    , campaign
    , ad
    , format
    , url
    , jumper
    -- Money
    , pricing_model
    , CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('impression', 'direct') THEN 
          view_price * sign
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END AS revenue
    , CASE
        WHEN pricing_model = 1 AND status = 2 AND event IN ('impression', 'direct') THEN 
          view_price * sign
        WHEN pricing_model = 2 AND status = 2 AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status = 2 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END AS failed_revenue
    , CASE
        WHEN pricing_model = 1 AND status = 3 AND event IN ('impression', 'direct') THEN 
          view_price * sign
        WHEN pricing_model = 2 AND status = 3 AND event IN ('click') THEN
          click_price * sign
        WHEN pricing_model = 3 AND status = 3 AND event IN ('lead') THEN
          lead_price * sign
        ELSE 0
      END AS compromised_revenue
    , CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('impression', 'direct') THEN 
          purchase_view_price * sign
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          purchase_click_price * sign
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          purchase_lead_price * sign
        ELSE 0
      END AS expenses
    , CASE
        WHEN pricing_model = 1 AND status = 1 AND event IN ('impression', 'direct') THEN 
          potential_view_price * sign
        WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
          potential_click_price * sign
        WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
          potential_lead_price * sign
        ELSE 0
      END AS potential
    , revenue - expenses - potential AS earnings -- minus expenses
    -- Counters
    , CAST(1 * sign AS UInt64) AS total
    , CAST(CASE WHEN status = 0   THEN 1 * sign ELSE 0 END AS UInt64) AS count
    , CAST(CASE WHEN status = 1   THEN 1 * sign ELSE 0 END AS UInt64) AS success
    , CAST(CASE WHEN status = 2   THEN 1 * sign ELSE 0 END AS UInt64) AS failed
    , CAST(CASE WHEN status = 3   THEN 1 * sign ELSE 0 END AS UInt64) AS compromised
    , CAST(CASE WHEN status = 4   THEN 1 * sign ELSE 0 END AS UInt64) AS custom
    , CAST(CASE WHEN status = 5   THEN 1 * sign ELSE 0 END AS UInt64) AS failovers
    , CAST(CASE WHEN adblock = 1  THEN 1 * sign ELSE 0 END AS UInt64) AS adblocks
    , CAST(CASE WHEN private = 1  THEN 1 * sign ELSE 0 END AS UInt64) AS privates
    , CAST(CASE WHEN robot = 1    THEN 1 * sign ELSE 0 END AS UInt64) AS robots
    , CAST(CASE WHEN backup = 1   THEN 1 * sign ELSE 0 END AS UInt64) AS backups
    -- Targeting
    , carrier
    , country
    , city
    , language
    , device_type
    , device
    , os
    , browser
    , categories
FROM stats.events_local;
