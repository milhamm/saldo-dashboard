DO $$ BEGIN IF NOT EXISTS (
    SELECT 1
    FROM pg_type
    WHERE typname = 'movement_type'
) THEN CREATE TYPE movement_type AS ENUM (
    'withdraw',
    'transfer',
    'top_up',
    'payment',
    'others'
);
END IF;
END $$;
CREATE TABLE IF NOT EXISTS movements (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    amount BIGINT NOT NULL,
    fee BIGINT NOT NULL,
    type movement_type NOT NULL,
    account_id UUID NOT NULL,
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
    FOREIGN KEY (account_id) REFERENCES accounts (id)
);
CREATE INDEX IF NOT EXISTS idx_movements_amount ON movements (amount);
CREATE INDEX IF NOT EXISTS idx_movements_account_id ON movements (account_id);
CREATE INDEX IF NOT EXISTS idx_movements_created_at ON movements (created_at);
CREATE INDEX IF NOT EXISTS idx_movements_updated_at ON movements (updated_at);