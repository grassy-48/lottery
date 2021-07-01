
CREATE DATABASE IF NOT EXISTS `ymtk` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `ymtk`;

CREATE TABLE `lt_codes` (
  `id` int(64) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `uniq_key` varchar(64) NOT NULL UNIQUE,
  `path` text NOT NULL,
  `owner` int(64) NOT NULL,
  `point_id` int(64) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `lt_gifts` (
  `id` int(64) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `name` text,
  `grade` varchar(8) DEFAULT NULL,
  `booth` text,
  `is_stock` tinyint(1) NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `lt_gift_histories` (
  `id` int(64) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` int(64) NOT NULL,
  `gift_id` int(64) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  UNIQUE(user_id, gift_id)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `lt_points` (
  `id` int(64) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `type` varchar(128) DEFAULT '',
  `point` int(8) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `lt_point_histories` (
  `id` int(64) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `user_id` int(64) NOT NULL,
  `code_id` int(64) NULL DEFAULT NULL,
  `plus` int(8) DEFAULT '0',
  `minus` int(8) DEFAULT '0',
  `result` int(32) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `lt_users` (
  `id` int(64) NOT NULL PRIMARY KEY AUTO_INCREMENT,
  `mail` varchar(64) NOT NULL UNIQUE,
  `name` varchar(32) DEFAULT '',
  `circle` varchar(32) DEFAULT '',
  `is_participant` tinyint(1) NOT NULL DEFAULT '0',
  `is_creator` tinyint(1) NOT NULL DEFAULT '0',
  `place` varchar(8) NOT NULL DEFAULT '',
  `point` int(16) NOT NULL DEFAULT '0',
  `total` int(16) NOT NULL DEFAULT '0',
  `minus` int(16) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

