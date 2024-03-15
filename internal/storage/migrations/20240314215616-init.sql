-- +migrate Up
CREATE TABLE access_data (
    time TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT date_trunc('second', NOW()),
    ip VARCHAR NOT NULL
);