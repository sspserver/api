CREATE TABLE stats.rtb_wins
( timemark            DateTime
, delay               UInt64      default 0     -- Delay of preparation of Ads in Nanosecinds (Load picture before display, etc)
, duration            UInt64      default 0     -- Duration in Nanosecond
, service             String
, cluster             FixedString(2)
-- source
, aucid               FixedString(16)           -- Internal Auction ID
, source              UInt64                    -- Advertisement Source ID
, network             String                    -- Source Network Name or Domain (Cross sails)
, access_point        UInt64                    -- Access Point ID to own Advertisement
)
ENGINE = ReplicatedMergeTree
PARTITION BY toYYYYMM(timemark)
ORDER BY (timemark, cluster, source, access_point)
SETTINGS index_granularity = 8192;
