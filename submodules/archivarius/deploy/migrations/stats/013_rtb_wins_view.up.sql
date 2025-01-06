CREATE MATERIALIZED VIEW IF NOT EXISTS stats.v_rtb_wins TO stats.rtb_wins
  AS SELECT
    toDate(datemark) AS datemark
  , delay
  , duration
  , service
  , cluster
  , aucid
  , source
  , network
  , access_point
  FROM stats.events_local
  WHERE event IN ('src.win', 'ap.win');
