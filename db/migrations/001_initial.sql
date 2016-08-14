-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE club
(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR (256) NOT NULL,
    leader VARCHAR (256),
    site1 VARCHAR (256),
    old_name VARCHAR (256),
    comment VARCHAR (256)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE club;
