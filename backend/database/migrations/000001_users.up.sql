CREATE TABLE IF NOT EXISTS users (
    id          BIGSERIAL NOT NULL primary key UNIQUE,
    name        VARCHAR(255) DEFAULT NULL,
    email       VARCHAR(255) DEFAULT NULL UNIQUE,
    created_at  TIMESTAMP default now(),
    updated_at  TIMESTAMP default now()
);