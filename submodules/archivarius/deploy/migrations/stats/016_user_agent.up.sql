CREATE TABLE stats.user_agent
( ua                 String
, cnt                UInt64         DEFAULT 1
) ENGINE = ReplicatedSummingMergeTree
ORDER BY (ua)
PRIMARY KEY (ua)
SETTINGS index_granularity = 8192;