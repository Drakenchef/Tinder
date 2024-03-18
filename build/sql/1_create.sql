DROP TABLE IF EXISTS images;
DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    uid uuid PRIMARY KEY,
    login text NOT NULL UNIQUE,
    passwordhash text NOT NULL,
    description text,
    salt text
);

CREATE TABLE images
(
    id SERIAL PRIMARY KEY,
    url text NOT NULL,
    user_id uuid REFERENCES users(uid) ON DELETE CASCADE
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