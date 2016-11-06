-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE f_competition ADD COLUMN approved_ash BOOL NOT NULL DEFAULT false;
ALTER TABLE f_competition ADD COLUMN raw_text TEXT;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE f_competition DROP COLUMN approved_ash;
ALTER TABLE f_competition DROP COLUMN raw_text;

