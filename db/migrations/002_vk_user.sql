-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE vk_user
(
    id BIGINT PRIMARY KEY NOT NULL,
    first_name VARCHAR(256) NOT NULL,
    last_name VARCHAR(256) NOT NULL
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE vk_user;
