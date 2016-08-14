-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE club
(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR (100) NOT NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE club;
