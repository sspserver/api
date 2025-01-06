--------------------------------------------------------------------------------
-- Create table adv_application to store information about the ad application (website, mobile app, etc.).
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS adv_application
( id                     BIGSERIAL                  PRIMARY KEY

, title                  VARCHAR(255)               NOT NULL
, description            TEXT

, account_id             BIGINT                     NOT NULL      REFERENCES account_base(id) MATCH SIMPLE
                                                                       ON UPDATE NO ACTION
                                                                       ON DELETE RESTRICT
, creator_id             BIGINT                     NOT NULL      REFERENCES account_user(id) MATCH SIMPLE
                                                                       ON UPDATE NO ACTION
                                                                       ON DELETE RESTRICT

, uri                    VARCHAR(255)               NOT NULL
, type                   ApplicationType            NOT NULL
, platform               PlatformType               NOT NULL
, premium                BOOLEAN                    NOT NULL      DEFAULT FALSE

-- Status
, status                 ApproveStatus              NOT NULL      DEFAULT 'pending'
, active                 ActiveStatus               NOT NULL      DEFAULT 'pause'
, private                PrivateStatus              NOT NULL      DEFAULT 'public'

, categories             BIGINT[]

-- Revenue Share
, revenue_share          NUMERIC(3, 10)             NOT NULL      DEFAULT 0

-- Adsources
, allowed_sources        BIGINT[]
, disallowed_sources     BIGINT[]

, created_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE TRIGGER updated_at_trigger 
  BEFORE UPDATE ON adv_application 
  FOR EACH ROW EXECUTE FUNCTION updated_at_column();

CREATE TRIGGER notify_update_event_trigger
  AFTER INSERT OR UPDATE OR DELETE ON adv_application
  FOR EACH ROW EXECUTE FUNCTION notify_update_event();


--------------------------------------------------------------------------------
-- Create table adv_zone to store information about the ad zone.
--------------------------------------------------------------------------------

CREATE TABLE adv_zone
( id                     BIGSERIAL                  PRIMARY KEY
, codename               VARCHAR(64)                NOT NULL      UNIQUE

, title                  VARCHAR(255)               NOT NULL
, description            TEXT

, account_id             BIGINT                     NOT NULL      REFERENCES account_base(id) MATCH SIMPLE
                                                                       ON UPDATE NO ACTION
                                                                       ON DELETE RESTRICT

, type                   ZoneType                   NOT NULL
, status                 ApproveStatus              NOT NULL      DEFAULT 'pending'
, active                 ActiveStatus               NOT NULL      DEFAULT 'pause'

, default_code           JSONB
, context                JSONB
, min_ecpm               NUMERIC(10, 5)             NOT NULL      DEFAULT 0
, min_ecpm_by_geo        JSONB

-- Price of buying per view
, fixed_purchase_price   NUMERIC(10, 5)             NOT NULL      DEFAULT 0

-- Filtering
, allowed_formats        BIGINT[]
, allowed_types          BIGINT[]
, allowed_sources        BIGINT[]
, disallowed_sources     BIGINT[]
, campaigns              BIGINT[]

, created_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE TRIGGER updated_at_trigger 
  BEFORE UPDATE ON adv_zone 
  FOR EACH ROW EXECUTE FUNCTION updated_at_column();

CREATE TRIGGER notify_update_event_trigger
  AFTER INSERT OR UPDATE OR DELETE ON adv_zone
  FOR EACH ROW EXECUTE FUNCTION notify_update_event();
