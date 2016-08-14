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

CREATE TABLE dancer (
    id BIGSERIAL PRIMARY KEY NOT NULL,

    code VARCHAR(256) NOT NULL,--TODO make unique

    name VARCHAR (256) NOT NULL,
    surname VARCHAR (256) NOT NULL,
    patronymic VARCHAR (256),
    sex VARCHAR(256) NOT NULL,

    pair_class VARCHAR(256) NOT NULL,
    jnj_class VARCHAR(256) NOT NULL,

    prev_surname VARCHAR(256),
    source VARCHAR(256)
);

CREATE TABLE dancer_club (
    dancer_id BIGINT NOT NULL REFERENCES dancer,
    club_id BIGINT NOT NULL REFERENCES club
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE dancer_club;
DROP TABLE club;
DROP TABLE dancer;
