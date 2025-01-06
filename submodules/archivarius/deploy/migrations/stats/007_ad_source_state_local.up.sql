CREATE TABLE IF NOT EXISTS stats.source_state_local (
   timemark           DateTime
 , datemark           Date
 , protocol           String
 , id                 UInt64      default 0

 , spent              Int64       default 0
 , profit             Int64       default 0
 , bid_price          Int64       default 0 -- Sum of all bid prices on the auction
 , potential          Int64       default 0 -- Sum of all potential prices on the auction

 , imps               UInt64      default 0
 , views              UInt64      default 0
 , directs            UInt64      default 0
 , clicks             UInt64      default 0
 , leads              UInt64      default 0
 , bids               UInt64      default 0
 , wins               UInt64      default 0
 , skips              UInt64      default 0
 , nobids             UInt64      default 0
 , errors             UInt64      default 0
)
Engine = ReplicatedSummingMergeTree
PARTITION BY toYYYYMM(datemark)
PRIMARY KEY (id, intHash32(id), protocol, datemark)
ORDER BY (id, intHash32(id), protocol, datemark)
SAMPLE BY intHash32(id);
