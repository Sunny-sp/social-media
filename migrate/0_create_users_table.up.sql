CREATE TABLE "user" (
    user_id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    mobile TEXT,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- -- 1) Create the trigger function (only once per table)
-- CREATE OR REPLACE FUNCTION update_updated_on_user_task()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.updated_at = now();
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- -- 2) Create the trigger
-- CREATE TRIGGER update_user_task_updated_on
-- BEFORE UPDATE ON user_task
-- FOR EACH ROW
-- EXECUTE FUNCTION update_updated_on_user_task();
