--------------------------------------------------------------------------------
-- Browser table - represents a browser with information about name, active status, and versions.
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS type_browser
( id                     BIGSERIAL                  PRIMARY KEY
, name                   VARCHAR(255)               NOT NULL
, version                VARCHAR(255)               NOT NULL
, description            TEXT                       NOT NULL

, year_release           INTEGER                    NOT NULL        DEFAULT 0
, year_end_support       INTEGER                    NOT NULL        DEFAULT 0

-- Match Name Expression
, match_name_exp         TEXT                       NOT NULL        DEFAULT ''
-- Match User Agent Expression
, match_ua_exp           TEXT                       NOT NULL        DEFAULT ''
-- Version of lower border in standard format (major.minor.patch[-pre-release][+build-metadata])
, match_ver_min_exp      TEXT                       NOT NULL        DEFAULT ''
-- Version of upper border in standard format (major.minor.patch[-pre-release][+build-metadata])
, match_ver_max_exp      TEXT                       NOT NULL        DEFAULT ''

-- Is Active Browser
, active                 ActiveStatus               NOT NULL        DEFAULT 'pause' -- ActiveStatus 'active' or 'pause'

, parent_id              BIGINT                     REFERENCES type_browser(id) MATCH SIMPLE
                                                        ON UPDATE NO ACTION
                                                        ON DELETE SET NULL
                                                    CHECK (parent_id <> id)

-- Time marks
, created_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE TRIGGER updated_at_trigger BEFORE UPDATE
    ON type_browser FOR EACH ROW EXECUTE PROCEDURE updated_at_column();

CREATE TRIGGER notify_update_event_trigger
AFTER INSERT OR UPDATE OR DELETE ON type_browser
    FOR EACH ROW EXECUTE PROCEDURE notify_update_event();
