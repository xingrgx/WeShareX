/*
SQLyog Community v13.1.9 (64 bit)
MySQL - 8.0.17 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `wsx_user` (
	`id` int (10),
	`passport` varchar (135),
	`password` char (96),
	`nickname` varchar (135),
	`email` varchar (60),
	`avatar` varchar (600),
	`status` tinyint (4),
	`gender` tinyint (1),
	`created_at` datetime ,
	`update_at` datetime 
); 
