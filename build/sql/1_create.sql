DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE IF NOT EXISTS users
(
    uid uuid not null unique,
    login text NOT NULL UNIQUE,
    passwordhash text NOT NULL,
    description text,
    img text,
    salt text
);