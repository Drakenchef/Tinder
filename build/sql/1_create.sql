DROP TABLE IF EXISTS profile CASCADE;
CREATE TABLE IF NOT EXISTS profile
(
    id uuid NOT NULL PRIMARY KEY,
    login text NOT NULL UNIQUE,
    description text,
    img text NOT NULL,
    passwordhash bytea NOT NULL
);