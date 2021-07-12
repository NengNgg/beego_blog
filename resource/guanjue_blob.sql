/*
Navicat MySQL Data Transfer

Source Server         : localhost_test
Source Server Version : 50528
Source Host           : localhost:3306
Source Database       : guanjue_blob

Target Server Type    : MYSQL
Target Server Version : 50528
File Encoding         : 65001

Date: 2021-07-12 21:16:44
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '帖子标题',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '帖子描述',
  `content` varchar(4000) NOT NULL DEFAULT '' COMMENT '帖子内容',
  `cover` varchar(255) NOT NULL DEFAULT 'static/upload/no_pic.jpg' COMMENT '帖子封面图',
  `read_num` int(11) NOT NULL DEFAULT '0' COMMENT '帖子阅读数',
  `star_num` int(11) NOT NULL DEFAULT '0' COMMENT '帖子点赞数',
  `author_id` int(11) NOT NULL COMMENT '帖子作者',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES ('4', 'beego', 'beego_study', '', 'static/upload/16260833101619515910bq3.png', '7', '0', '1', '2021-07-12 17:48:30');

-- ----------------------------
-- Table structure for sys_post_comment
-- ----------------------------
DROP TABLE IF EXISTS `sys_post_comment`;
CREATE TABLE `sys_post_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(4000) NOT NULL DEFAULT '' COMMENT '评论内容',
  `post_id` int(11) NOT NULL COMMENT '帖子外键',
  `p_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级评论',
  `author_id` int(11) NOT NULL COMMENT '评论人',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_post_comment
-- ----------------------------
INSERT INTO `sys_post_comment` VALUES ('7', '真棒', '4', '0', '6', '2021-07-12 17:49:46');
INSERT INTO `sys_post_comment` VALUES ('8', '真棒\n', '4', '0', '5', '2021-07-12 19:05:12');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `is_admin` int(11) NOT NULL DEFAULT '2' COMMENT '1是管理员，2是普通用户',
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `cover` varchar(255) NOT NULL DEFAULT 'static/upload/bq3.png' COMMENT '头像',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'admin', '25d55ad283aa400af464c76d713c07ad', '1', '2021-07-12 17:36:50', 'static/upload/bq3.png');
INSERT INTO `sys_user` VALUES ('5', '黄春', '25f9e794323b453885f5181f1b624d0b', '2', '2021-07-12 17:29:00', 'static/upload/bq3.png');
INSERT INTO `sys_user` VALUES ('6', '熊凯棣', '25d55ad283aa400af464c76d713c07ad', '2', '2021-07-12 17:48:52', 'static/upload/bq3.png');
