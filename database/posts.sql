-- AI学习社区 - Posts表结构
-- 设计版本: v1.0
-- 创建时间: 2026-04-10

CREATE DATABASE IF NOT EXISTS `ai_study_community` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `ai_study_community`;

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '帖子ID',
  `title` varchar(255) NOT NULL COMMENT '帖子标题',
  `content` longtext NOT NULL COMMENT '帖子内容(Markdown/富文本)',
  `author_id` bigint unsigned NOT NULL COMMENT '作者用户ID',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 0-删除 1-正常 2-置顶 3-精华',
  `view_count` int unsigned NOT NULL DEFAULT 0 COMMENT '浏览量',
  `like_count` int unsigned NOT NULL DEFAULT 0 COMMENT '点赞数',
  `comment_count` int unsigned NOT NULL DEFAULT 0 COMMENT '评论数',
  `tags` varchar(500) DEFAULT NULL COMMENT '标签,多个逗号分隔',
  `cover_image` varchar(500) DEFAULT NULL COMMENT '封面图片URL',
  `is_deleted` tinyint NOT NULL DEFAULT 0 COMMENT '软删除: 0-未删除 1-已删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_status` (`status`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_is_deleted` (`is_deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='帖子表';
