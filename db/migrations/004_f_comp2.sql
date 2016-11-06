-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE f_competition ADD COLUMN approved_ash BOOL NOT NULL DEFAULT false;
ALTER TABLE f_competition ADD COLUMN raw_text TEXT;
ALTER TABLE f_competition ADD COLUMN raw_text_changed TEXT;
ALTER TABLE f_competition ADD COLUMN download_date TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW();
ALTER TABLE f_competition ADD COLUMN has_change BOOL NOT NULL DEFAULT FALSE;

ALTER TABLE f_competition ADD CONSTRAINT f_competition__url_unique UNIQUE (url);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE f_competition DROP COLUMN IF EXISTS approved_ash;
ALTER TABLE f_competition DROP COLUMN IF EXISTS raw_text;
ALTER TABLE f_competition DROP COLUMN IF EXISTS raw_text_changed;
ALTER TABLE f_competition DROP COLUMN IF EXISTS download_date;
ALTER TABLE f_competition DROP COLUMN IF EXISTS has_change;

ALTER TABLE f_competition DROP CONSTRAINT IF EXISTS f_competition__url_unique;