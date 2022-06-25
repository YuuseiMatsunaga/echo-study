
-- +migrate Up

SET AUTOCOMMIT = 0;
BEGIN;
create table users
(
    id int
    auto_increment primary key, 
    name varchar
    (255), 
    email varchar
    (255), 
    created_at timestamp
);

    INSERT INTO users
        (id, name, email, created_at)
    VALUES
        (1, "Yamada", "yamada@example.com", Now());
    INSERT INTO users
        (id, name, email, created_at)
    VALUES
        (2, "Tanaka", "tanaka@example.com", Now());
COMMIT;

-- +migrate Down
DROP TABLE IF EXISTS users;