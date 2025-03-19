CREATE TABLE IF NOT EXISTS adv_traffic_router
( id                     BIGSERIAL                  PRIMARY KEY
, account_id             BIGINT                     NOT NULL      REFERENCES account_base(id) MATCH SIMPLE
                                                                       ON UPDATE NO ACTION
                                                                       ON DELETE RESTRICT

, percent                NUMERIC                    NOT NULL       DEFAULT 100
                                                        CHECK (percent >= 0 AND percent <= 100)

, description            TEXT                       NOT NULL

-- Is Active traffic router
, active                 ActiveStatus               NOT NULL      DEFAULT 'pause'

, rtb_source_ids         BIGINT[]                   NOT NULL      CHECK (array_length(rtb_source_ids, 1) > 0)

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
, domains                TEXT[]
, apps                   BIGINT[]
, zones                  BIGINT[]
, secure                 INT                        NOT NULL      DEFAULT 0   CHECK (secure IN (0, 1, 2))
, adblock                INT                        NOT NULL      DEFAULT 0   CHECK (adblock IN (0, 1, 2))
, private_browsing       INT                        NOT NULL      DEFAULT 0   CHECK (private_browsing IN (0, 1, 2))
, ip                     INT                        NOT NULL      DEFAULT 0   CHECK (ip IN (0, 1, 2))

-- Time marks
, created_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE OR REPLACE TRIGGER updated_at_trigger
  BEFORE UPDATE ON adv_traffic_router
  FOR EACH ROW EXECUTE PROCEDURE updated_at_column();

CREATE OR REPLACE TRIGGER notify_update_event_trigger
  AFTER INSERT OR UPDATE OR DELETE ON adv_traffic_router
  FOR EACH ROW EXECUTE PROCEDURE notify_update_event();
