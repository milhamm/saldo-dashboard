CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    created_at BIGINT DEFAULT (
        EXTRACT(
            epoch
            FROM now()
        ) * 1000::numeric
    ),
    updated_at BIGINT DEFAULT (
        EXTRACT(
            epoch
            FROM now()
        ) * 1000::numeric
    )
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_phone ON users (phone);
CREATE INDEX IF NOT EXISTS idx_users_name ON users (name);
CREATE INDEX IF NOT EXISTS idx_users_created_at ON users (created_at);
CREATE INDEX IF NOT EXISTS idx_users_updated_at ON users (updated_at);