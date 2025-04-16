-- migrations/0001_create_tables.up.sql

CREATE TABLE disciplines (
                             id SERIAL PRIMARY KEY,
                             name TEXT NOT NULL UNIQUE
);

CREATE TABLE tests (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       discipline_id INTEGER NOT NULL UNIQUE REFERENCES disciplines(id)
);

CREATE TABLE grades (
                        id SERIAL PRIMARY KEY,
                        score DOUBLE NOT NULL,
                        test_id INTEGER NOT NULL UNIQUE REFERENCES tests(id)
);
