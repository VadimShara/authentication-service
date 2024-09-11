CREATE TABLE refresh_tokens (
    username VARCHAR(16) UNIQUE NOT NULL,
    refresh_token TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);