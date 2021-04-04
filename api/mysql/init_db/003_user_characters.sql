CREATE DATABASE IF NOT EXISTS `go_db`;

USE `go_db`;

CREATE TABLE IF NOT EXISTS `user_characters` (
    `id`          INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `userId`      INT NOT NULL,
    `characterId` INT NOT NULL,
    CONSTRAINT fk_user_id      FOREIGN KEY (`userId`)      REFERENCES users (`id`)      ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_charatcer_id FOREIGN KEY (`characterId`) REFERENCES characters (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4;
