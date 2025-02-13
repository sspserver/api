--------------------------------------------------------------------------------
-- Category table - represents a category with information about name, description, and hierarchy.
--------------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS adv_category
( id                     BIGSERIAL                  PRIMARY KEY
, name                   VARCHAR(255)               NOT NULL
, description            TEXT                       NOT NULL

, iab_code               VARCHAR(255)               NOT NULL        UNIQUE -- IAB code

, parent_id              BIGINT                     REFERENCES adv_category(id) MATCH SIMPLE
                                                                       ON UPDATE NO ACTION
                                                                       ON DELETE SET NULL
                                                    CHECK (parent_id <> id)
, position               BIGINT                     NOT NULL

-- Is Active advertisement
, active                 ActiveStatus               NOT NULL        DEFAULT 'pause'

-- Time marks
, created_at             TIMESTAMP                  NOT NULL        DEFAULT NOW()
, updated_at             TIMESTAMP                  NOT NULL        DEFAULT NOW()
, deleted_at             TIMESTAMP
);

CREATE TRIGGER updated_at_trigger BEFORE UPDATE
    ON adv_category FOR EACH ROW EXECUTE PROCEDURE updated_at_column();

CREATE TRIGGER notify_update_event_trigger
AFTER INSERT OR UPDATE OR DELETE ON adv_category
    FOR EACH ROW EXECUTE PROCEDURE notify_update_event();
