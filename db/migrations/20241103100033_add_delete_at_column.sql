-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER table users ADD column deleted_at timestamptz default null;
ALTER table conversations ADD column deleted_at timestamptz default null;
ALTER table messages ADD column deleted_at timestamptz default null;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
