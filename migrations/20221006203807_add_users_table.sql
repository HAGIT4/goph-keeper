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

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
