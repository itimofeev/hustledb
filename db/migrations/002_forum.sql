-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE partition
(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    index INTEGER NOT NULL,
    competition_id BIGINT NOT NULL REFERENCES competition
);

CREATE TABLE judge (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    dancer_id BIGINT NOT NULL REFERENCES dancer,
    partition_id BIGINT NOT NULL REFERENCES partition,
    letter VARCHAR(1) NOT NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE judge;
DROP TABLE partition;
