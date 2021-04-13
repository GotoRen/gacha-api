CREATE DATABASE IF NOT EXISTS `go_db`;

USE `go_db`;

CREATE TABLE IF NOT EXISTS `users` (
    `id`         INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name`       VARCHAR(256) NOT NULL,
    `token`      VARCHAR(256) NOT NULL,
    `created_at` TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;

INSERT INTO users(name, token) VALUE('RenGoto', 'abc');
INSERT INTO users(name, token) VALUE('KobayashiFumiaki', 'def');
