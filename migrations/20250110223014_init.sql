-- +goose Up
-- +goose StatementBegin
-- SELECT 'select 1';
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_indexes
        WHERE tablename = 'audit_logs'
          AND indexname = 'idx_time_stamp'
    ) THEN
CREATE INDEX idx_time_stamp ON audit_logs (timestamp);
END IF;
END $$;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'select 2';
DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM pg_indexes
        WHERE tablename = 'audit_logs'
          AND indexname = 'idx_time_stamp'
    ) THEN
DROP INDEX idx_time_stamp;
END IF;
END $$;
-- +goose StatementEnd
