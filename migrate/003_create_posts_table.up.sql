CREATE TABLE IF NOT EXISTS "post" (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES "user"(user_id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    tags TEXT[],  -- PostgreSQL array type
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
