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
DROP TABLE IF EXISTS likes CASCADE;
CREATE TABLE IF NOT EXISTS likes
(
    ID SERIAL PRIMARY KEY,
    uidfirstlike uuid NOT NULL,
    uidsecondlike uuid NOT NULL,
    Date TIMESTAMP NOT NULL,
    Mutual BOOLEAN NOT NULL
);