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
    source VARCHAR(256),

    CONSTRAINT dancer__code_unique UNIQUE (code),
    CONSTRAINT dancer__sex_values_check CHECK (sex in ('m', 'f')),
    CONSTRAINT dancer__pair_class_values_check CHECK (pair_class in ('A', 'B', 'C', 'D', 'E')),
    CONSTRAINT dancer__jnj_class_values_check CHECK (jnj_class in ('BG', 'RS', 'M', 'S', 'Ch'))
);

CREATE TABLE dancer_club (
    dancer_id BIGINT NOT NULL REFERENCES dancer,
    club_id BIGINT NOT NULL REFERENCES club
);

CREATE TABLE competition (
       id BIGSERIAL PRIMARY KEY NOT NULL,
       title VARCHAR(256) NOT NULL,
       date TIMESTAMP WITHOUT TIME ZONE NOT NULL,
       site VARCHAR(256)
);

CREATE TABLE nomination (
       id BIGSERIAL PRIMARY KEY NOT NULL,
       competition_id BIGINT NOT NULL REFERENCES competition,

       value VARCHAR(256) NOT NULL,

       male_count INT NOT NULL,
       female_count INT NOT NULL,
       type VARCHAR(10) NOT NULL,
       min_class VARCHAR(10),
       max_class VARCHAR(10),
       min_jnj_class VARCHAR(10),
       max_jnj_class VARCHAR(10),

       CONSTRAINT nomination__type_check CHECK (type in ('OLD_JNJ', 'NEW_JNJ', 'CLASSIC')),
       CONSTRAINT nomination__min_class_check CHECK (min_class in ('A', 'B', 'C', 'D', 'E')),
       CONSTRAINT nomination__max_class_check CHECK (max_class in ('A', 'B', 'C', 'D', 'E')),
       CONSTRAINT nomination__min_jnj_class_check CHECK (min_jnj_class in ('BG', 'RS', 'M', 'S', 'Ch')),
       CONSTRAINT nomination__max_jnj_class_check CHECK (max_jnj_class in ('BG', 'RS', 'M', 'S', 'Ch'))
);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE dancer_club;
DROP TABLE club;
DROP TABLE dancer;
DROP TABLE competition;
