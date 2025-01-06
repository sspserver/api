CREATE MATERIALIZED VIEW IF NOT EXISTS stats.v_accesspoint_state_live TO stats.accesspoint_state_local
AS SELECT
   toStartOfFifteenMinutes(timemark) AS timemark
 , datemark
 , '' AS protocol
 , access_point AS id
 , SUM(CASE
     WHEN pricing_model = 1 AND status = 1 AND event IN ('view', 'direct') THEN 
       view_price * sign
     WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
       click_price * sign
     WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
       lead_price * sign
     ELSE 0
  END) AS spent
 , SUM(CASE
     WHEN pricing_model = 1 AND status = 1 AND event IN ('view', 'direct') THEN
       multiIf(purchase_view_price > 0,  purchase_view_price * sign,  view_price * sign)
     WHEN pricing_model = 2 AND status IN (0, 1) AND event IN ('click') THEN
       multiIf(purchase_click_price > 0, purchase_click_price * sign, click_price * sign)
     WHEN pricing_model = 3 AND status IN (0, 1) AND merged = 1 AND event IN ('lead') THEN
       multiIf(purchase_lead_price > 0,  purchase_lead_price * sign,  lead_price * sign)
     ELSE 0
   END) AS profit -- Revenue of the publisher
 , SUM(CAST(CASE WHEN event = 'impression' AND status IN (0, 1) THEN sign ELSE 0 END AS UInt64)) AS imps
 , SUM(CAST(CASE WHEN event = 'view'       AND status IN (0, 1) THEN sign ELSE 0 END AS UInt64)) AS views
 , SUM(CAST(CASE WHEN event = 'direct'     AND status IN (0, 1) THEN sign ELSE 0 END AS UInt64)) AS directs
 , SUM(CAST(CASE WHEN event = 'click'      AND status IN (0, 1) THEN sign ELSE 0 END AS UInt64)) AS clicks
 , SUM(CAST(CASE WHEN event = 'lead'       AND status IN (0, 1) THEN sign ELSE 0 END AS UInt64)) AS leads
 , SUM(CAST(CASE WHEN event = 'src.bid'                         THEN sign ELSE 0 END AS UInt64)) AS bids
 , SUM(CAST(CASE WHEN event = 'src.win'                         THEN sign ELSE 0 END AS UInt64)) AS wins
 , SUM(CAST(CASE WHEN event = 'src.skip'   AND status IN (0, 1) THEN sign ELSE 0 END AS UInt64)) AS skips
 , SUM(CAST(CASE WHEN event = 'src.nobid'  AND status IN (0, 1) THEN sign ELSE 0 END AS UInt64)) AS nobids
 , SUM(CAST(CASE WHEN event = 'src.fail'                        THEN sign ELSE 0 END AS UInt64)) AS errors
FROM stats.events_local
GROUP BY (
   datemark
 , access_point
 , timemark
);
