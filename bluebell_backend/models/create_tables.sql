use bluebell;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `age` int(20) collate utf8mb4_general_ci NULL,
    `sex` varchar(20) collate utf8mb4_general_ci NULL,
    `address` varchar(64) COLLATE utf8mb4_general_ci NULL,
    `postId` int(20) collate utf8mb4_general_ci NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `captcha_id` varchar(4) DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `user_status` varchar(64) DEFAULT '未申请' COLLATE utf8mb4_general_ci,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DROP TABLE IF EXISTS `community`;
CREATE TABLE `community`(
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `community_id` int(10) unsigned NOT NULL,
  `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_community_id` (`community_id`),
  UNIQUE KEY `idx_community_name` (`community_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
INSERT INTO `community` VALUES ('1', '1', 'Go', 'Golang', '2016-11-01 08:10:10', '2016-11-01 08:10:10');
INSERT INTO `community` VALUES ('2', '2', 'leetcode', '刷题刷题刷题', '2020-01-01 08:00:00', '2020-01-01 08:00:00');
INSERT INTO `community` VALUES ('3', '3', 'PUBG', '大吉大利，今晚吃鸡。', '2018-08-07 08:30:00', '2018-08-07 08:30:00');
INSERT INTO `community` VALUES ('4', '4', 'LOL', '欢迎来到英雄联盟!', '2016-01-01 08:00:00', '2016-01-01 08:00:00');

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` bigint(20) NOT NULL COMMENT '帖子id',
  `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
  `author_id` bigint(20) NOT NULL COMMENT '作者的用户id',
  `community_id` bigint(20) NOT NULL COMMENT '所属社区',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '帖子状态',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_post_id` (`post_id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `comment_id` bigint(20) unsigned NOT NULL,
  `content` text COLLATE utf8mb4_general_ci NOT NULL,
  `post_id` bigint(20) NOT NULL,
  `author_id` bigint(20) NOT NULL,
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_comment_id` (`comment_id`),
  KEY `idx_author_Id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `project`;
CREATE TABLE `project`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `project_user_name` varchar(255) NOT NUll,
    `project_university` varchar(255) NOT NULL,
    `project_college` varchar(255) NOT NULL,
    `project_major` varchar(255) NOT NULL,
    `project_email` varchar(255) NOT NULL,
    `project_phone` bigint NOT NULL,
    `projectDirection` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET= utf8mb4 COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `project1`;
CREATE TABLE `project1`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `project_captain` varchar(255) NOT NULL,
    `project_detail_sort` varchar(255) NOT NULL,
    `project_detail_name` varchar(255) NOT NULL,
    `project_detail_person` text NOT NULL,
    `project_detail_intro` varchar(255) NOT NULL,
    `project_detail_idea` varchar(255) NOT NULL ,
    `project_detail_adv` varchar(255) NOT NULL ,
    `project_detail_teacher` text NOT NULL ,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_ProjectDetail_Name` (project_detail_name)
)ENGINE = InnoDB DEFAULT CHARSET =utf8mb4 COLLATE = utf8mb4_general_ci;

DROP TABLE IF EXISTS `projects`;
CREATE TABLE projects (
                          id bigint(20) NOT NULL AUTO_INCREMENT,
                          name VARCHAR(255) NOT NULL  COMMENT '组长名字',
                          university VARCHAR(255) NOT NULL COMMENT '学校',
                          college VARCHAR(255) NOT NULL Comment '学院',
                          major TEXT NOT NULL COMMENT '专业',
                          email TEXT NOT NULL COMMENT '邮箱',
                          phone TEXT NOT NULL COMMENT '手机号码',
                          project_direction VARCHAR(255) NOT NULL COMMENT  '申报类型',
                          project_sort varchar(255) NOT NULL COMMENT '项目类型',
                          project_name varchar(255) NOT NULL COMMENT  '项目名称',
                          member varchar(255) NOT NULL,
                          introduction varchar(255) NOT NULL,
                          creativity varchar(255) NOT NULL ,
                          advantage varchar(255) NOT NULL,
                          instructor varchar(255) NOT NULL,
<<<<<<< HEAD
                          status varchar(255) ,
                          content varchar(255),
=======
                          status varchar(255) NOT NULL ,
                          content varchar(255) ,
>>>>>>> a1518829ab829eb2efd1ee7fcfd2989a02d795a2
                          PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET =utf8mb4 COLLATE = utf8mb4_general_ci;

use bluebell;
CREATE TABLE IF NOT EXISTS admins (
                                      user_id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
                                      username VARCHAR(255) NOT NULL,
                                      password VARCHAR(255) NOT NULL,
                                      PRIMARY KEY(`user_id`)
)ENGINE = InnoDB DEFAULT CHARSET =utf8mb4 COLLATE = utf8mb4_general_ci;