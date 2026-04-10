-- AI学习社区 - Comments表结构
-- 设计版本: v1.0
-- 创建时间: 2026-04-10
-- 支持树形层级结构，parent_id自关联

CREATE DATABASE IF NOT EXISTS `ai_study_community` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE `ai_study_community`;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `post_id` bigint unsigned NOT NULL COMMENT '所属帖子ID',
  `parent_id` bigint unsigned DEFAULT NULL COMMENT '父评论ID(NULL为一级评论)',
  `content` longtext NOT NULL COMMENT '评论内容',
  `author_id` bigint unsigned NOT NULL COMMENT '评论者用户ID',
  `reply_to_user_id` bigint unsigned DEFAULT NULL COMMENT '被回复的用户ID(NULL则是一级评论)',
  `depth` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '评论层级深度(0=一级)',
  `like_count` int unsigned NOT NULL DEFAULT 0 COMMENT '点赞数',
  `status` tinyint NOT NULL DEFAULT 1 COMMENT '状态: 0-删除 1-正常',
  `is_deleted` tinyint NOT NULL DEFAULT 0 COMMENT '软删除: 0-未删除 1-已删除',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_post_id` (`post_id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_status` (`status`),
  KEY `idx_is_deleted` (`is_deleted`),
  CONSTRAINT `fk_comments_parent` FOREIGN KEY (`parent_id`) REFERENCES `comments` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `fk_comments_post` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表(支持树形层级)';
