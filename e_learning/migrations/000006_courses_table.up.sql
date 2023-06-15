CREATE TABLE IF NOT EXISTS courses (
                                       id bigserial PRIMARY KEY,
                                       name text NOT NULL,
                                       credits integer NOT NULL,
                                       price text NOT NULL
);