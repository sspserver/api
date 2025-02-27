---------------------------------------------------------------------
-- OS
---------------------------------------------------------------------

-- Match expressions
-- By default, the match expressions is equalities case insensitive
-- For example, for the OS "Windows", the match expression is "Windows"
-- For the OS "Android", the match expression is "Android"
-- For the OS "iOS", the match expression is "ios"
-- Supported formats by prefix:
-- $regex$ - Regular expression
-- $wk$ - Wildcard expression

-- Match version expressions
-- The version of the OS is a string in the format "major.minor.patch[-pre-release][+build-metadata]"
-- For example, for the OS "Windows 10", the match version expression is "10"
-- For the OS "Android 10", the match version expression is "10"
-- Supported formats by prefix:
-- $regex$ - Regular expression
-- $wk$ - Wildcard expression

CREATE TABLE IF NOT EXISTS type_os
( id                     BIGSERIAL                  PRIMARY KEY
, name                   VARCHAR(255)               NOT NULL
, version                VARCHAR(255)               NOT NULL
, description            TEXT                       NOT NULL

, year_release           INTEGER                    NOT NULL      DEFAULT 0
, year_end_support       INTEGER                    NOT NULL      DEFAULT 0

-- Match Name Expression
, match_name_exp         VARCHAR(255)               NOT NULL
-- Match User Agent Expression
, match_ua_exp           VARCHAR(255)               NOT NULL
-- Version of lower border in standard format (major.minor.patch[-pre-release][+build-metadata])
, match_ver_min_exp      VARCHAR(255)               NOT NULL
-- Version of upper border in standard format (major.minor.patch[-pre-release][+build-metadata])
, match_ver_max_exp      VARCHAR(255)               NOT NULL

-- Is Active OS
, active                 ActiveStatus               NOT NULL      DEFAULT 'pause' -- ActiveStatus 'active' or 'pause'

, parent_id              BIGINT                     REFERENCES type_os(id) MATCH SIMPLE
                                                        ON UPDATE NO ACTION
                                                        ON DELETE SET NULL
                                                    CHECK (parent_id <> id)

-- Time marks
, created_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL      DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE TRIGGER updated_at_trigger BEFORE UPDATE
    ON type_os FOR EACH ROW EXECUTE PROCEDURE updated_at_column();

CREATE TRIGGER notify_update_event_trigger
AFTER INSERT OR UPDATE OR DELETE ON type_os
    FOR EACH ROW EXECUTE PROCEDURE notify_update_event();
