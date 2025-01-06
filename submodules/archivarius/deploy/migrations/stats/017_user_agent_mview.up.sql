CREATE MATERIALIZED VIEW IF NOT EXISTS stats.v_user_agent TO stats.user_agent
  AS SELECT ua, 1 AS cnt FROM stats.events_local;
