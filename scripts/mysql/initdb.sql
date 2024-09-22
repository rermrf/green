CREATE DATABASE IF NOT EXISTS `green`;

USE `green`;

CREATE TABLE `users` (
                         `id` bigint NOT NULL AUTO_INCREMENT,
                         `email` varchar(191) DEFAULT NULL,
                         `password` longtext,
                         `phone` varchar(191) DEFAULT NULL,
                         `nickname` longtext,
                         `about_me` longtext,
                         `birthday` bigint,
                         `ctime` bigint DEFAULT NULL,
                         `utime` bigint DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `uni_users_email` (`email`),
                         UNIQUE KEY `uni_users_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
