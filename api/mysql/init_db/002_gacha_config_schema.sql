CREATE DATABASE IF NOT EXISTS `go_db`;

USE `go_db`;

CREATE TABLE IF NOT EXISTS gacha_configs(
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `reality` INT NOT NULL,
    `probability` FLOAT NOT NULL
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;

INSERT INTO gacha_configs(reality, probability) VALUE(5, 0.01);

INSERT INTO gacha_configs(reality, probability) VALUE(4, 0.1);

INSERT INTO gacha_configs(reality, probability) VALUE(3, 0.89);