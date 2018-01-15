CREATE TABLE IF NOT EXISTS article (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    user_id bigint(20) unsigned NOT NULL,
    title varchar(100) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO article(id, user_id, title) VALUES (DEFAULT, 1, 'Title 1');
INSERT INTO article(id, user_id, title) VALUES (DEFAULT, 1, 'Title 2');
