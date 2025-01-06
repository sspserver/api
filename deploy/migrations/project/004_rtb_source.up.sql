CREATE TABLE IF NOT EXISTS rtb_source
( id                     BIGSERIAL                  PRIMARY KEY
, account_id             BIGINT                     NOT NULL      REFERENCES account_base(id) MATCH SIMPLE
                                                                       ON UPDATE NO ACTION
                                                                       ON DELETE RESTRICT

, title                  VARCHAR(255)               NOT NULL
, description            TEXT

, status                 ApproveStatus              NOT NULL      DEFAULT 'pending'
, active                 ActiveStatus               NOT NULL      DEFAULT 'pause'
, flags                  JSONB

, protocol               VARCHAR(255)               NOT NULL
, minimal_weight         NUMERIC                    NOT NULL
, url                    TEXT                       NOT NULL
, method                 VARCHAR(10)                NOT NULL      DEFAULT 'POST'
, request_type           RTBRequestType             NOT NULL      DEFAULT 'undefined'
, headers                JSONB
, rps                    INT                        NOT NULL      DEFAULT 0
, timeout                INT                        NOT NULL

-- Money configs
, accuracy                NUMERIC
, price_correction_reduce NUMERIC
, auction_type            AuctionType                NOT NULL      DEFAULT 'undefined'

-- Price limits
, min_bid                NUMERIC
, max_bid                NUMERIC

-- Targeting filters
, formats                TEXT[]
, device_types           BIGINT[]
, devices                BIGINT[]
, os                     BIGINT[]
, browsers               BIGINT[]
, carriers               BIGINT[]
, categories             BIGINT[]
, countries              TEXT[]
, languages              TEXT[]
, apps                   BIGINT[]
, domains                TEXT[]
, zones                  BIGINT[]
, external_zones         BIGINT[]
, secure                 INT                        NOT NULL      DEFAULT 0
, adblock                INT                        NOT NULL      DEFAULT 0
, private_browsing       INT                        NOT NULL      DEFAULT 0
, ip                     INT                        NOT NULL      DEFAULT 0

, config                 JSONB

, created_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE TRIGGER updated_at_trigger 
  BEFORE UPDATE ON rtb_source 
  FOR EACH ROW EXECUTE PROCEDURE updated_at_column();

CREATE TRIGGER notify_update_event_trigger
  AFTER INSERT OR UPDATE OR DELETE ON rtb_source
  FOR EACH ROW EXECUTE PROCEDURE notify_update_event();

COMMENT ON TABLE rtb_source IS 'RTB Source';
COMMENT ON COLUMN rtb_source.secure IS 'Secure flag (0 - any, 1 - only, 2 - exclude)';
COMMENT ON COLUMN rtb_source.adblock IS 'Adblock flag (0 - any, 1 - only, 2 - exclude)';
COMMENT ON COLUMN rtb_source.private_browsing IS 'Private browsing flag (0 - any, 1 - only, 2 - exclude)';
COMMENT ON COLUMN rtb_source.ip IS 'IP flag (0 - any, 1 - only, 2 - exclude)';
