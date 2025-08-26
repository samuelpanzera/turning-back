-- Initial database setup for Turning Back API

-- Create extensions if they don't exist
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create initial tables (these will be managed by GORM migrations in the application)
-- This file is mainly for database initialization and extensions

-- You can add initial data here if needed
-- INSERT INTO users (email, username, first_name, last_name, password, role) 
-- VALUES ('admin@example.com', 'admin', 'Admin', 'User', crypt('admin123', gen_salt('bf')), 'admin');

-- Create indexes for better performance (if not handled by GORM)
-- CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
-- CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);

-- Log the initialization
DO $$
BEGIN
    RAISE NOTICE 'Database initialized successfully for Turning Back API';
END $$;