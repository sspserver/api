--------------------------------------------------------------------------------
-- DeviceMaker table - represents a device maker with information about name, description, active status, and models.
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS type_device_maker
( id                     BIGSERIAL                  PRIMARY KEY
, codename               VARCHAR(255)               NOT NULL       UNIQUE
, name                   VARCHAR(255)               NOT NULL
, description            TEXT                       NOT NULL

-- Match expression
-- This expression is used to match the device maker name with the user agent string.
-- By default, it is a simple case-insensitive match.
-- With prefix '$regex$' | '$re$', it is a regular expression.
-- With prefix '$wc$', it is a wildcard expression.
, match_exp              TEXT                       NOT NULL

-- Is Active device maker
, active                 ActiveStatus               NOT NULL      DEFAULT 'pause'

-- Time marks
, created_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE TRIGGER updated_at_trigger BEFORE UPDATE
    ON type_device_maker FOR EACH ROW EXECUTE PROCEDURE updated_at_column();

CREATE TRIGGER notify_update_event_trigger
AFTER INSERT OR UPDATE OR DELETE ON type_device_maker
    FOR EACH ROW EXECUTE PROCEDURE notify_update_event();

--------------------------------------------------------------------------------
-- DeviceModel table - represents a device model with information about name, description, active status, and maker.
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS type_device_model
( id                     BIGSERIAL                  PRIMARY KEY
, codename               VARCHAR(255)               NOT NULL       UNIQUE
, name                   VARCHAR(255)               NOT NULL
, description            TEXT                       NOT NULL

, year_release           INTEGER                    NOT NULL        DEFAULT 0

-- Match expression
-- This expression is used to match the device model name with the user agent string.
-- By default, it is a simple case-insensitive match.
-- With prefix '$regex$' | '$re$', it is a regular expression.
-- With prefix '$wc$', it is a wildcard expression.
, match_exp              TEXT                       NOT NULL

-- Is Active device model
, active                 ActiveStatus               NOT NULL      DEFAULT 'pause'

-- Maker
, maker_codename         VARCHAR(255)               NOT NULL      REFERENCES type_device_maker(codename) MATCH SIMPLE
                                                                       ON UPDATE NO ACTION
                                                                       ON DELETE RESTRICT

-- Device Type
, type_codename          VARCHAR(255)               NOT NULL       CHECK (type_codename <> '')

, parent_id              BIGINT                     REFERENCES type_device_model(id) MATCH SIMPLE
                                                        ON UPDATE NO ACTION
                                                        ON DELETE SET NULL
                                                    CHECK (parent_id <> id)

-- Time marks
, created_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE TRIGGER updated_at_trigger BEFORE UPDATE
    ON type_device_model FOR EACH ROW EXECUTE PROCEDURE updated_at_column();

CREATE TRIGGER notify_update_event_trigger
AFTER INSERT OR UPDATE OR DELETE ON type_device_model
    FOR EACH ROW EXECUTE PROCEDURE notify_update_event();
