CREATE DATABASE IF NOT EXISTS `go_db`;

USE `go_db`;

CREATE TABLE IF NOT EXISTS `characters` (
    `id`         INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name`       VARCHAR(256) NOT NULL,
    `reality`    INT          NOT NULL,
    `created_at` TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

INSERT INTO characters(name, reality) VALUE('金村美玖', 5);
INSERT INTO characters(name, reality) VALUE('河田陽菜', 5);
INSERT INTO characters(name, reality) VALUE('小坂菜緒', 5);
INSERT INTO characters(name, reality) VALUE('富田鈴花', 5);
INSERT INTO characters(name, reality) VALUE('丹生明里', 4);
INSERT INTO characters(name, reality) VALUE('濱岸ひより', 4);
INSERT INTO characters(name, reality) VALUE('松田好花', 4);
INSERT INTO characters(name, reality) VALUE('宮田愛萌', 4);
INSERT INTO characters(name, reality) VALUE('渡邊美穂', 3);
INSERT INTO characters(name, reality) VALUE('上村ひなの', 3);
INSERT INTO characters(name, reality) VALUE('髙橋未来虹', 3);
INSERT INTO characters(name, reality) VALUE('森本茉莉', 3);
INSERT INTO characters(name, reality) VALUE('山口陽世', 3);
