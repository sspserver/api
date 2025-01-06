
CREATE TABLE IF NOT EXISTS stats.user_info_local (
   sign               Int8              default 1
 , timemark           DateTime
 , datehourmark       DateTime          default toStartOfHour(timemark)
 , datemark           Date              default toDate(timemark)
 , aucid              FixedString(16)             -- Internal Auction ID

 -- User IDENTITY
 , udid               String                      -- Unique Device ID (IDFA)
 , uuid               FixedString(16)
 , sessid             FixedString(16)

 -- User personal information
 , age                UInt8
 , gender             UInt8
 , search_gender      UInt8
 , email              String
 , phone              String
 , messanger_type     String
 , messanger          String
 , zip                String
 , facebook           String
 , twitter            String
 , linkedin           String

 -- Targeting
 , carrier            UInt64                      -- Carrier ID
 , country            FixedString(2)              -- Country Code
 , city               String                      -- City Code
 , latitude           Float64                     -- Geo latitude
 , longitude          Float64                     -- Geo longitude
 , language           FixedString(5)              -- en-US

 , created_at         DateTime          default now()
)
ENGINE = ReplicatedCollapsingMergeTree(sign)
PARTITION BY toYYYYMM(datemark)
ORDER BY (datemark, aucid, gender, email, facebook, twitter, linkedin)
SETTINGS index_granularity = 8192;
