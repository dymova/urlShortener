DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS urls;

CREATE TABLE users
(
    id       INT AUTO_INCREMENT NOT NULL,
    login    VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE urls
(
    id    INT AUTO_INCREMENT NOT NULL,
    shortened VARCHAR(255) NOT NULL UNIQUE,
    full  TEXT         NOT NULL,
    owner INT         NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`owner`) REFERENCES users (`id`)
);