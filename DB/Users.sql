CREATE TABLE Users (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `user_name` varchar(45) COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '会員氏名',
  `mail` varchar(255) COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'メールアドレス',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '作成時間',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新時間',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '削除時間',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='ユーザー一覧テーブル';