--------------------------------------------------------------------------------
-- Browser table - represents a browser with information about name, active status, and versions.
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS type_browser
( id                     BIGSERIAL                  PRIMARY KEY
, name                   VARCHAR(255)               NOT NULL
, description            TEXT                       NOT NULL

, match_exp              TEXT                       NOT NULL

-- Is Active browser
, active                 ActiveStatus               NOT NULL      DEFAULT 'pause'

-- Versions
, versions               JSONB

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
