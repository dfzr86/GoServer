CREATE DATABASE IF NOT EXISTS mydatabase DEFAULT CHARACTER SET = utf8 DEFAULT COLLATE = utf8_unicode_ci;

use mydatabase;

CREATE TABLE IF NOT EXISTS `t_user` (
  `openId`    varchar (255)   NOT NULL,
  `nick`      varchar (255)   DEFAULT NULL,
  `avatar`    varchar (255)   DEFAULT NULL,
  `level`     varchar (255)   DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS `t_article` (
  `id`              int   NOT NULL AUTO_INCREMENT,
  `title`           varchar (255)   NOT NULL,
  `titleDesc`       varchar (255)   DEFAULT NULL,
  `coverImage`      varchar (255)   DEFAULT NULL,
  `contentPath`     varchar (255)   DEFAULT NULL,
  PRIMARY KEY (`id`)
);

REPLACE INTO t_user (openId, nick, level) VALUES ("0", "zimusjw", "-100");
REPLACE INTO t_article (id, title, titleDesc) VALUES ("0", "第一篇文章", "第一篇文章简介");




