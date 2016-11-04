-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE f_competition
(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    url varchar(256) NOT NULL,
    date DATE NOT NULL,
    title varchar(256) NOT NULL,
    description varchar(256),
    city varchar(256)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE f_competition;
