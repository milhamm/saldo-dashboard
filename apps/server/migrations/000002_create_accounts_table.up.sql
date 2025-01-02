CREATE TABLE IF NOT EXISTS accounts (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR NOT NULL,
    user_id UUID NOT NULL,
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
    ),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE INDEX IF NOT EXISTS idx_accounts_user_id ON accounts (user_id);
CREATE INDEX IF NOT EXISTS idx_accounts_created_at ON accounts (created_at);
CREATE INDEX IF NOT EXISTS idx_accounts_updated_at ON accounts (updated_at);