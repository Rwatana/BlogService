-- create_database.sql

-- データベースを作成する
CREATE DATABASE IF NOT EXISTS test_db;

-- データベースを使用する
USE test_db;

-- テーブルを作成する
CREATE TABLE IF NOT EXISTS `test_log3` (
  `log_level` VARCHAR(100) NOT NULL COMMENT "Log_Level",
  `date` TIMESTAMP NOT NULL DEFAULT  CURRENT_TIMESTAMP COMMENT "Date",
  `current_service` VARCHAR(100) NOT NULL COMMENT "Current_API",
  `source_service` VARCHAR(100) NOT NULL COMMENT "Source_API",
  `type_of_request` VARCHAR(100) NOT NULL COMMENT "Request_Type",
  `content` VARCHAR(200) NOT NULL COMMENT "Error_Sentence"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE test_log ALTER COLUMN Date TYPE TIMESTAMP
