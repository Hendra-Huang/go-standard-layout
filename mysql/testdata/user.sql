CREATE TABLE IF NOT EXISTS users (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO users(id, email, name) VALUES (DEFAULT, 'myuser@example.com', 'Myuser');
INSERT INTO users(id, email, name) VALUES (DEFAULT, 'myuser2@example.com', 'Myuser2');
