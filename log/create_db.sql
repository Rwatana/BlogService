-- create_database.sql

-- データベースを作成する
CREATE DATABASE IF NOT EXISTS test_db;

-- データベースを使用する
USE test_db;

-- テーブルを作成する
CREATE TABLE IF NOT EXISTS `test_log` (
  `date` TIMESTAMP NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "Date",
  `service` VARCHAR(100) NOT NULL COMMENT "API",
  `content` VARCHAR(200) NOT NULL COMMENT "Error_Sentence"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE test_log ALTER COLUMN Date TYPE TIMESTAMP
