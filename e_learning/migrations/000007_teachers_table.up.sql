CREATE TABLE IF NOT EXISTS teachers (
                                      id bigserial PRIMARY KEY,
                                      name text NOT NULL,
                                      surname text NOT NULL,
                                      subject text NOT NULL,
                                      salary text NOT NULL
);