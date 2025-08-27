CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255),
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
