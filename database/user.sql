CREATE TABLE `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) NOT NULL UNIQUE,
    PRIMARY KEY (`id`)
  );