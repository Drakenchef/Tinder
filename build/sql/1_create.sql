DROP TABLE IF EXISTS profile CASCADE;
CREATE TABLE IF NOT EXISTS profile
(
    uid uuid NOT NULL PRIMARY KEY,
    login text NOT NULL UNIQUE,
    description text,
    img text NOT NULL
);

DROP TABLE IF EXISTS users CASCADE;
CREATE TABLE IF NOT EXISTS users
(
    uid serial not null unique,
    login text NOT NULL UNIQUE,
    passwordhash bytea NOT NULL
);