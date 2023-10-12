-- +goose Up
CREATE TABLE users (
    id number NOT NULL PRIMARY KEY,
    username varchar2(100),
    name varchar2(100),
    surname varchar2(100)
);

INSERT INTO users VALUES
(0, 'root', '', ''),
(1, 'vojtechvitek', 'Vojtech', 'Vitek');

-- +goose Down
DROP TABLE users;
