CREATE TABLE IF NOT EXISTS notes
(
    id          BIGSERIAL    NOT NULL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);