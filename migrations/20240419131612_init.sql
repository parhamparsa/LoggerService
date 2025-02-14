-- +goose Up
-- +goose StatementBegin
-- SELECT 'select 1';
-- TODO: INDEX timestamp and status code
CREATE TABLE audit_logs
(
    id            UUID PRIMARY KEY,
    timestamp     TIMESTAMP   NOT NULL,
    status_code   INT         NOT NULL,
    duration_ms   INT         NOT NULL,
    method        VARCHAR(50) NOT NULL,
    path          TEXT        NOT NULL,
    ip_address    VARCHAR(45) NOT NULL,
    user_agent    TEXT        NOT NULL,
    request_body  TEXT NULL,
    response_body TEXT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'select 2';
drop table audit_logs
-- +goose StatementEnd
