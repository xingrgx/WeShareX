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

/*Data for the table `wsx_file` */

insert  into `wsx_file`(`id`,`user_id`,`name`,`parent_id`,`path`,`dir`,`create_at`,`update_at`,`type`,`size`,`times`) values 
('1vlyjjjevw0cjjdwkhwmv80100rp1d0b',1,'A.Dir','1vlyjjjevw0cjjdwno7m748w00sth5ls','/B.Dir/A.Dir',1,'2022-04-25 22:18:42','2022-04-26 11:21:08','',0,0),
('1vlyjjjevw0cjjdwkibo2h8300h1o6wi',1,'A3.Dir','1vlyjjjevw0cjjdwkhwmv80100rp1d0b','/B.Dir/A.Dir/A3.Dir',1,'2022-04-25 22:18:42','2022-04-26 11:43:52','',0,0),
('1vlyjjjevw0cjjdwkipfus4500uw3o6r',1,'新建 Microsoft Word 文档.docx','1vlyjjjevw0cjjdwkl2uqgks004icq8s','/B.Dir/A.Dir/A1.Dir/A11.Dir/新建 Microsoft Word 文档.docx',0,'2022-04-25 22:18:42','2022-04-26 11:31:43','.docx',0,0),
('1vlyjjjevw0cjjdwkivi5cg8003hphec',1,'A2.Dir','1vlyjjjevw0cjjdwkhwmv80100rp1d0b','/B.Dir/A.Dir/A2.Dir',1,'2022-04-25 22:18:42','2022-04-26 11:21:08','',0,0),
('1vlyjjjevw0cjjdwkj8cfpga004m1nc7',1,'A21.md','1vlyjjjevw0cjjdwkivi5cg8003hphec','/B.Dir/A.Dir/A2.Dir/A21.md',0,'2022-04-25 22:18:42','2022-04-26 11:21:08','.md',0,0),
('1vlyjjjevw0cjjdwkjefzesd00oyiqa3',1,'A1.Dir','1vlyjjjevw0cjjdwkhwmv80100rp1d0b','/B.Dir/A.Dir/A1.Dir',1,'2022-04-25 22:18:42','2022-04-26 11:31:42','',0,0),
('1vlyjjjevw0cjjdwkjs5418f00tk045d',1,'A12.Dir','1vlyjjjevw0cjjdwkjefzesd00oyiqa3','/B.Dir/A.Dir/A1.Dir/A12.Dir',1,'2022-04-25 22:18:42','2022-04-26 11:31:43','',0,0),
('1vlyjjjevw0cjjdwkk17auwh00zrvrcp',1,'A122.md','1vlyjjjevw0cjjdwkjs5418f00tk045d','/B.Dir/A.Dir/A1.Dir/A12.Dir/A122.md',0,'2022-04-25 22:18:42','2022-04-26 11:31:43','.md',0,0),
('1vlyjjjevw0cjjdwkkiru8cm00m1e42s',1,'A121.Dir','1vlyjjjevw0cjjdwkjs5418f00tk045d','/B.Dir/A.Dir/A1.Dir/A12.Dir/A121.Dir',1,'2022-04-25 22:18:42','2022-04-26 11:31:43','',0,0),
('1vlyjjjevw0cjjdwkkso93go00o1pld9',1,'A1211.txt','1vlyjjjevw0cjjdwkkiru8cm00m1e42s','/B.Dir/A.Dir/A1.Dir/A12.Dir/A121.Dir/A1211.txt',0,'2022-04-25 22:18:42','2022-04-26 11:31:43','.txt',0,0),
('1vlyjjjevw0cjjdwkl2uqgks004icq8s',1,'A11.Dir','1vlyjjjevw0cjjdwkjefzesd00oyiqa3','/B.Dir/A.Dir/A1.Dir/A11.Dir',1,'2022-04-25 22:18:42','2022-04-26 11:31:43','',0,0),
('1vlyjjjevw0cjjdwklcbnasu003dvt2k',1,'A111.md','1vlyjjjevw0cjjdwkl2uqgks004icq8s','/B.Dir/A.Dir/A1.Dir/A11.Dir/A111.md',0,'2022-04-25 22:18:42','2022-04-26 11:31:43','.md',0,0),
('1vlyjjjevw0cjjdwno7m748w00sth5ls',1,'B.Dir','root','/B.Dir',1,'2022-04-25 22:18:49','2022-04-25 22:18:49','',0,0),
('bwziv208go0ck4n71hlwsu0100uhun73',1,'2022届计算机系毕设答辩计划.pdf','root','/2022届计算机系毕设答辩计划.pdf',0,'2022-05-20 22:01:32','2022-05-20 22:01:32','.pdf',271383,0),
('bwziv209go0ck9b95uoohb83007vr029',2,'A.Dir','root','/A.Dir',1,'2022-05-26 09:43:24','2022-05-26 09:43:24','',0,0),
('bwziv209go0ck9b95v00jl8500om9aay',2,'A3.Dir','bwziv209go0ck9b95uoohb83007vr029','/A.Dir/A3.Dir',1,'2022-05-26 09:43:24','2022-05-26 09:43:24','',0,0),
('bwziv209go0ck9b95vc0zew80004aop2',2,'新建 Microsoft Word 文档.docx','bwziv209go0ck9b95v00jl8500om9aay','/A.Dir/A3.Dir/新建 Microsoft Word 文档.docx',0,'2022-05-26 09:43:24','2022-05-26 09:43:24','.docx',0,0),
('bwziv209go0ck9b95vh2v3ka009qqied',2,'A2.Dir','bwziv209go0ck9b95uoohb83007vr029','/A.Dir/A2.Dir',1,'2022-05-26 09:43:24','2022-05-26 09:43:24','',0,0),
('bwziv209go0ck9b95vsnnykd00cdrq3r',2,'A21.md','bwziv209go0ck9b95vh2v3ka009qqied','/A.Dir/A2.Dir/A21.md',0,'2022-05-26 09:43:24','2022-05-26 09:43:24','.md',0,0),
('bwziv209go0ck9b95vxgfekf00jns2s2',2,'A1.Dir','bwziv209go0ck9b95uoohb83007vr029','/A.Dir/A1.Dir',1,'2022-05-26 09:43:24','2022-05-26 09:43:24','',0,0),
('bwziv209go0ck9b95w871m4h00sie3w1',2,'A12.Dir','bwziv209go0ck9b95vxgfekf00jns2s2','/A.Dir/A1.Dir/A12.Dir',1,'2022-05-26 09:43:24','2022-05-26 09:43:24','',0,0),
('bwziv209go0ck9b95whma9ok00mb3et8',2,'A122.md','bwziv209go0ck9b95w871m4h00sie3w1','/A.Dir/A1.Dir/A12.Dir/A122.md',0,'2022-05-26 09:43:24','2022-05-26 09:43:24','.md',0,0),
('bwziv209go0ck9b95wtrm90o008oab5w',2,'A121.Dir','bwziv209go0ck9b95w871m4h00sie3w1','/A.Dir/A1.Dir/A12.Dir/A121.Dir',1,'2022-05-26 09:43:24','2022-05-26 09:43:24','',0,0),
('bwziv209go0ck9b95x3get4r00kk35xg',2,'A1211.txt','bwziv209go0ck9b95wtrm90o008oab5w','/A.Dir/A1.Dir/A12.Dir/A121.Dir/A1211.txt',0,'2022-05-26 09:43:24','2022-05-26 09:43:24','.txt',0,0),
('bwziv209go0ck9b95xbr7d8u004tzz70',2,'A11.Dir','bwziv209go0ck9b95vxgfekf00jns2s2','/A.Dir/A1.Dir/A11.Dir',1,'2022-05-26 09:43:24','2022-05-26 09:43:24','',0,0),
('bwziv209go0ck9b95xqvmfww004szvor',2,'A111.md','bwziv209go0ck9b95xbr7d8u004tzz70','/A.Dir/A1.Dir/A11.Dir/A111.md',0,'2022-05-26 09:43:24','2022-05-26 09:43:24','.md',11,0),
('zt4iri05k80cka9jkdwc580j00v5btkc',1,'aaaaa','root','/aaaaa',1,'2022-05-27 12:35:37','2022-05-27 12:35:37','',0,0),
('zt4iri05k80ckabp524uq8cn00pgkqpu',1,'IMG_20220518_172622.jpg','root','/IMG_20220518_172622.jpg',0,'2022-05-27 14:16:56','2022-05-27 14:16:56','.jpg',3918946,0);

/*Table structure for table `wsx_friends` */

DROP TABLE IF EXISTS `wsx_friends`;

CREATE TABLE `wsx_friends` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `me` int(10) unsigned NOT NULL COMMENT '用户我ID',
  `friend` int(10) unsigned NOT NULL COMMENT '好友ID',
  `time` datetime NOT NULL COMMENT '添加时间',
  `status` int(1) unsigned zerofill NOT NULL COMMENT '状态 等待同意:1 同意:0 已同意:2',
  PRIMARY KEY (`id`),
  KEY `me` (`me`),
  KEY `friend` (`friend`),
  CONSTRAINT `wsx_friends_ibfk_1` FOREIGN KEY (`me`) REFERENCES `wsx_user` (`id`),
  CONSTRAINT `wsx_friends_ibfk_2` FOREIGN KEY (`friend`) REFERENCES `wsx_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `wsx_friends` */

insert  into `wsx_friends`(`id`,`me`,`friend`,`time`,`status`) values 
(53,2,1,'2022-05-27 20:47:33',2),
(54,1,2,'2022-05-27 20:47:33',2),
(55,2,1,'2022-05-27 20:48:37',2),
(56,1,2,'2022-05-27 20:48:37',2),
(57,33,1,'2022-05-27 23:04:43',2),
(58,1,33,'2022-05-27 23:04:43',2),
(59,33,1,'2022-05-28 08:20:29',1),
(60,1,33,'2022-05-28 08:20:29',0),
(61,33,1,'2022-05-28 08:28:21',1),
(62,1,33,'2022-05-28 08:28:21',0);

/*Table structure for table `wsx_record` */

DROP TABLE IF EXISTS `wsx_record`;

CREATE TABLE `wsx_record` (
  `id` int(32) NOT NULL AUTO_INCREMENT COMMENT '聊天记录ID',
  `sender_id` int(10) unsigned NOT NULL COMMENT '发送者ID',
  `receiver_id` int(10) unsigned NOT NULL COMMENT '接收者ID',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '内容',
  `time` datetime NOT NULL COMMENT '发送时间',
  PRIMARY KEY (`id`),
  KEY `sender` (`sender_id`),
  KEY `receiver` (`receiver_id`),
  CONSTRAINT `wsx_record_ibfk_1` FOREIGN KEY (`sender_id`) REFERENCES `wsx_user` (`id`),
  CONSTRAINT `wsx_record_ibfk_2` FOREIGN KEY (`receiver_id`) REFERENCES `wsx_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `wsx_record` */

insert  into `wsx_record`(`id`,`sender_id`,`receiver_id`,`content`,`time`) values 
(15,1,33,'1111','2022-05-27 23:06:16'),
(16,33,1,'jjj','2022-05-27 23:13:37'),
(17,33,1,'dddddd','2022-05-27 23:19:52');

/*Table structure for table `wsx_relation` */

DROP TABLE IF EXISTS `wsx_relation`;

CREATE TABLE `wsx_relation` (
  `id` varchar(32) NOT NULL COMMENT 'ID',
  `share_id` varchar(32) NOT NULL COMMENT '分享ID',
  `file_id` varchar(32) NOT NULL COMMENT '文件ID',
  PRIMARY KEY (`id`),
  KEY `wsx_relation_ibfk_1` (`share_id`),
  KEY `wsx_relation_ibfk_2` (`file_id`),
  CONSTRAINT `wsx_relation_ibfk_1` FOREIGN KEY (`share_id`) REFERENCES `wsx_share` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `wsx_relation_ibfk_2` FOREIGN KEY (`file_id`) REFERENCES `wsx_file` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `wsx_relation` */

insert  into `wsx_relation`(`id`,`share_id`,`file_id`) values 
('zt4iri05k80ckabs5z9btukr0027hxfm','zt4iri05k80ckabs5z5riekp00l7w5cs','zt4iri05k80ckabp524uq8cn00pgkqpu'),
('zt4iri05k80ckadggf49pc8z0065ybnh','zt4iri05k80ckadggf19yv8x00ezcf6y','bwziv208go0ck4n71hlwsu0100uhun73');

/*Table structure for table `wsx_share` */

DROP TABLE IF EXISTS `wsx_share`;

CREATE TABLE `wsx_share` (
  `id` varchar(32) NOT NULL COMMENT '分享ID',
  `user_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `name` varchar(255) NOT NULL COMMENT '分享名',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `code` varchar(15) DEFAULT NULL COMMENT '提取码',
  `times` int(10) NOT NULL COMMENT '下载次数',
  `never_expire` tinyint(1) NOT NULL COMMENT '1:不过期；0:过期',
  `expire_at` datetime NOT NULL COMMENT '过期时间',
  `type` tinyint(1) DEFAULT NULL COMMENT '1:文件夹；0:文件',
  `nickname` varchar(45) NOT NULL COMMENT '分享者',
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `wsx_share_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `wsx_user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `wsx_share` */

insert  into `wsx_share`(`id`,`user_id`,`name`,`update_at`,`create_at`,`code`,`times`,`never_expire`,`expire_at`,`type`,`nickname`) values 
('zt4iri05k80ckabs5z5riekp00l7w5cs',1,'IMG_20220518_172622.jpg','2022-05-27 14:20:53','2022-05-27 14:20:53','LqFd',0,0,'2022-06-27 14:20:54',0,'zhangsan'),
('zt4iri05k80ckadggf19yv8x00ezcf6y',1,'2022届计算机系毕设答辩计划.pdf','2022-05-27 15:39:38','2022-05-27 15:39:38','PRxX',0,0,'2022-06-27 15:39:38',0,'zhangsan');

/*Table structure for table `wsx_user` */

DROP TABLE IF EXISTS `wsx_user`;

CREATE TABLE `wsx_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `passport` varchar(45) NOT NULL COMMENT '账号',
  `password` char(32) NOT NULL COMMENT '密码',
  `nickname` varchar(45) DEFAULT 'User_id' COMMENT '昵称',
  `email` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '邮箱',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT 'https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg' COMMENT '头像图床地址',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态（0:启用 1:禁用）',
  `gender` tinyint(1) DEFAULT NULL COMMENT '性别（0:未设置 1:男 2:女）',
  `created_at` datetime DEFAULT NULL COMMENT '注册时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `wsx_user` */

insert  into `wsx_user`(`id`,`passport`,`password`,`nickname`,`email`,`avatar`,`status`,`gender`,`created_at`,`update_at`) values 
(1,'xht','xht','zhangsan','xingrgx@foxmail.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',2,0,'2022-05-27 20:23:48','2022-05-27 20:36:12'),
(2,'zhangsan','123456','李四','san.zhang@163.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',1,1,'2022-03-30 20:46:59','2022-05-27 23:05:36'),
(22,'zwk','123456','测试','eddievandeer@163.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',0,1,'2022-04-01 14:49:47','2022-05-28 08:28:58'),
(24,'1','2222','你好','22222222','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',1,0,NULL,'2022-05-28 08:29:02'),
(25,'zheng','','zheng','1255173159@qq.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',0,0,'2022-04-14 21:02:59','2022-05-27 23:20:26'),
(27,'xue.ht','123456','xue.ht','xue.ht@qq.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',1,NULL,'2022-05-27 10:46:42','2022-05-28 08:21:06'),
(28,'lisi','123456','李四','lisi@qq.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',0,NULL,'2022-05-27 10:50:12','2022-05-27 10:50:12'),
(29,'eddievandeer','123456','Eddievandeer','qwer@sina.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',0,NULL,'2022-05-27 15:35:56','2022-05-27 15:35:56'),
(30,'yinpeng','123','士大夫','1669351755@qq.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',0,NULL,'2022-05-27 15:36:19','2022-05-27 15:36:19'),
(31,'a','123456','阿发画家回复','dashfka@asd.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',0,NULL,'2022-05-27 15:59:29','2022-05-27 15:59:29'),
(33,'xuehuitao','123456','zhangsan','xuehuitao@qq.com','https://s3.bmp.ovh/imgs/2022/03/31/fc24a51977107d97.jpg',1,1,'2022-05-27 22:40:05','2022-05-28 08:25:44');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
