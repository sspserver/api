CREATE TABLE IF NOT EXISTS adv_format
( id              BIGSERIAL              PRIMARY KEY
, codename        VARCHAR(255)           NOT NULL
, type            VARCHAR(255)           NOT NULL
, title           VARCHAR(255)           NOT NULL
, description     TEXT                   NOT NULL       DEFAULT ''
, active          ActiveStatus           NOT NULL

, width           INT
, height          INT
, min_width       INT
, min_height      INT
, config          JSONB

, created_at      TIMESTAMP              NOT NULL      DEFAULT NOW()
, updated_at      TIMESTAMP              NOT NULL      DEFAULT NOW()
, deleted_at      TIMESTAMP
);

CREATE TRIGGER updated_at_triger BEFORE UPDATE
    ON adv_format FOR EACH ROW EXECUTE PROCEDURE updated_at_column();

CREATE TRIGGER notify_update_event_trigger
AFTER INSERT OR UPDATE OR DELETE ON adv_format
    FOR EACH ROW EXECUTE PROCEDURE notify_update_event();
