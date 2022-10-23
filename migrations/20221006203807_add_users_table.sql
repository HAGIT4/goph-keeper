-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE SCHEMA IF NOT EXISTS keeper;
CREATE TABLE keeper.user (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    passhash VARCHAR(255) NOT NULL
);

CREATE TABLE keeper.loginpass (
    id UUID PRIMARY KEY,
    keyword VARCHAR(255) NOT NULL UNIQUE,
    login VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    meta TEXT
)

CREATE TABLE keeper.text (
    id UUID PRIMARY KEY,
    keyword VARCHAR(255) NOT NULL,
    text TEXT NOT NULL,
    meta TEXT
)

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
