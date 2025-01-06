CREATE TABLE IF NOT EXISTS stats.model_state_local (
   sign               Int8        default 1
 , timemark           DateTime
 , datemark           Date        default toDate(timemark)
 , model              String
 , id                 UInt64      default 0
 , spent              Int64       default 0
 , imps               Int64       default 0
 , clicks             Int64       default 0
 , leads              Int64       default 0
 , profit             Int64       default 0
 , created_at         DateTime    default now()
)
Engine = ReplicatedCollapsingMergeTree(sign)
PARTITION BY toYYYYMM(datemark)
ORDER BY (id, model, datemark)
SETTINGS index_granularity = 8192;
