-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE f_partition
(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    index INTEGER NOT NULL,
    competition_id BIGINT NOT NULL REFERENCES competition
);

CREATE TABLE f_judge (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    dancer_id BIGINT NOT NULL REFERENCES dancer,
    partition_id BIGINT NOT NULL REFERENCES f_partition,
    letter VARCHAR(1) NOT NULL
);

CREATE TABLE f_nomination (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    partition_id BIGINT NOT NULL REFERENCES f_partition,
    r_nomination_id BIGINT REFERENCES nomination,
    title VARCHAR(256) NOT NULL
);

CREATE TABLE f_place (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    place_from INTEGER NOT NULL,
    place_to INTEGER NOT NULL,
    number INTEGER NOT NULL,
    stage_title VARCHAR(128) NOT NULL,
    nomination_id BIGINT REFERENCES f_nomination,
    dancer1_id BIGINT NOT NULL REFERENCES dancer,
    dancer2_id BIGINT REFERENCES dancer,

    result1_id BIGINT REFERENCES result,
    result2_id BIGINT REFERENCES result
);

CREATE TABLE f_dancer_club (
    dancer_id BIGINT NOT NULL REFERENCES dancer,
    competition_id BIGINT NOT NULL REFERENCES competition,
    club_id BIGINT NOT NULL REFERENCES club,

    CONSTRAINT f_dancer_club_pk PRIMARY KEY(dancer_id, competition_id, club_id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE f_place;
DROP TABLE f_judge;
DROP TABLE f_nomination;
DROP TABLE f_partition;
