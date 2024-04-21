-- create_database.sql

-- データベースを作成する
CREATE DATABASE IF NOT EXISTS test_db;

-- データベースを使用する
USE test_db;

-- テーブルを作成する
CREATE TABLE IF NOT EXISTS `test_log` (
  `date` INT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "Date",
  `api` VARCHAR(100) NOT NULL COMMENT "API",
  `error` VARCHAR(200) NOT NULL COMMENT "Error_Sentence"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
