CREATE EXTENSION pgcrypto;

CREATE SCHEMA data;

CREATE SCHEMA api;

CREATE TABLE data.users
(
    id       SERIAL PRIMARY KEY,
    name     TEXT NOT NULL,
    email    TEXT NOT NULL UNIQUE,
    pwd_hash TEXT NOT NULL
);
