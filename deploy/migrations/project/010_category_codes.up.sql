--------------------------------------------------------------------------------
-- Category taxonomy codes.
-- r0_code is the hierarchy key. c3_1_code is Content Taxonomy 3.1.
-- cattax stores the Ad Product Taxonomy 2.0 id (JSON cattax_id).
-- iab_code stays unique but may be null: new nodes often have no IAB 1.0 code.
-- NOT NULL on r0_code is applied by the fixture after every row is filled.
--------------------------------------------------------------------------------

ALTER TABLE adv_category
    ADD COLUMN r0_code   VARCHAR(128),
    ADD COLUMN c3_1_code VARCHAR(64),
    ADD COLUMN cattax    VARCHAR(32);

ALTER TABLE adv_category
    ALTER COLUMN iab_code DROP NOT NULL;

ALTER TABLE adv_category
    DROP CONSTRAINT adv_category_iab_code_key;

CREATE INDEX adv_category_iab_code_idx
    ON adv_category (iab_code);

CREATE UNIQUE INDEX adv_category_r0_code_uidx
    ON adv_category (r0_code)
    WHERE r0_code IS NOT NULL;
