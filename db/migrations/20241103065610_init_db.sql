-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    CREATE TABLE IF NOT EXISTS users (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
        login text not null unique ,
        password text not null,
        salt bytea not null,
        created_at timestamptz DEFAULT now()
    );

    CREATE TABLE IF NOT EXISTS conversations (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
        name TEXT,
        created_at timestamptz DEFAULT now()
    );

    CREATE TABLE IF NOT EXISTS participants (
        conversation_id uuid,
        user_id uuid,
        PRIMARY KEY (conversation_id, user_id)
    );

    CREATE TABLE IF NOT EXISTS messages(
        id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
        message TEXT NOT NULL,
        sender_id uuid not null,
        conversation_id uuid not null,
        created_at timestamptz DEFAULT now()
    );

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
