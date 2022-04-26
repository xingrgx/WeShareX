/*
SQLyog Community v13.1.9 (64 bit)
MySQL - 8.0.17 : Database - we_sharex
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`we_sharex` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `we_sharex`;

/*Table structure for table `wsx_file` */

DROP TABLE IF EXISTS `wsx_file`;

CREATE TABLE `wsx_file` (
  `id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件ID',
  `user_id` int(10) unsigned DEFAULT NULL COMMENT '用户ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '文件名称',
  `parent_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '父级目录ID',
  `path` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '路径',
  `dir` tinyint(1) DEFAULT NULL COMMENT '是否为文件夹',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `type` varchar(10) DEFAULT NULL COMMENT '文件类型（扩展名）',
  `size` bigint(20) DEFAULT NULL COMMENT '文件大小',
  `times` bigint(20) DEFAULT NULL COMMENT '下载次数',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `wsx_file_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `wsx_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `wsx_share` */

DROP TABLE IF EXISTS `wsx_share`;

CREATE TABLE `wsx_share` (
  `id` varchar(32) NOT NULL COMMENT '分享ID',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(255) NOT NULL COMMENT '分享名',
  `update_time` timestamp NOT NULL COMMENT '更新时间',
  `create_time` timestamp NOT NULL COMMENT '创建时间',
  `code` varchar(15) DEFAULT NULL COMMENT '提取码',
  `times` int(10) NOT NULL COMMENT '下载次数',
  `never_expire` tinyint(1) NOT NULL COMMENT '1:不过期；0:过期',
  `expire_time` timestamp NOT NULL COMMENT '过期时间',
  `type` tinyint(1) DEFAULT NULL COMMENT '1:文件夹；0:文件',
  `file_id` varchar(32) NOT NULL COMMENT '文件文件夹ID',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  KEY `file_id` (`file_id`),
  CONSTRAINT `wsx_share_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `wsx_user` (`id`),
  CONSTRAINT `wsx_share_ibfk_2` FOREIGN KEY (`file_id`) REFERENCES `wsx_file` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `wsx_user` */

DROP TABLE IF EXISTS `wsx_user`;

CREATE TABLE `wsx_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `passport` varchar(45) NOT NULL COMMENT '账号',
  `password` char(32) NOT NULL COMMENT '密码',
  `nickname` varchar(45) DEFAULT 'User_id' COMMENT '昵称',
  `email` varchar(20) NOT NULL COMMENT '邮箱',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg' COMMENT '头像图床地址',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态（0:启用 1:禁用）',
  `gender` tinyint(1) DEFAULT NULL COMMENT '性别（0:未设置 1:男 2:女）',
  `created_at` datetime DEFAULT NULL COMMENT '注册时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
