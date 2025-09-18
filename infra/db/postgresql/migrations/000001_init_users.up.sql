-- create table if table if not exist in database 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; 
CREATE TABLE IF NOT EXISTS users (
    id          UUID PRIMARY KEY,
    username    VARCHAR(150) NOT NULL UNIQUE,
    email       VARCHAR(255) NOT NULL UNIQUE,
    fullname    VARCHAR(255),
    passwd_hash TEXT NOT NULL,
    status      VARCHAR(20) NOT NULL DEFAULT 'active'
                 CHECK (status IN ('active','inactive','banned','disabled','deleted')),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Trigger function: auto update field updated_at
CREATE OR REPLACE FUNCTION set_users_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- trigger : run func if insert or update record 
DROP TRIGGER IF EXISTS trg_set_users_updated_at ON users;
CREATE TRIGGER trg_set_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION set_users_updated_at();
