-- create_database.sql

-- データベースを作成する
CREATE DATABASE IF NOT EXISTS test_db;

-- データベースを使用する
USE test_db;

-- テーブルを作成する
CREATE TABLE IF NOT EXISTS `test_log3` (
  `log_level` VARCHAR(100) NOT NULL COMMENT "Log_Level",
  `date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "Date",
  `current_service` VARCHAR(100) NOT NULL COMMENT "Current_API",
  `source_service` VARCHAR(100) NOT NULL COMMENT "Source_API",
  `type_of_request` VARCHAR(100) NOT NULL COMMENT "Request_Type",
  `content` VARCHAR(200) NOT NULL COMMENT "Error_Sentence"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 日付カラムの型をTIMESTAMPに変更する（既にTIMESTAMPのため不要）
-- ALTER TABLE test_log3 MODIFY COLUMN `date` TIMESTAMP;


INSERT INTO test_log3 (log_level, date, current_service, source_service, type_of_request, content)
VALUES 
('INFO', CURRENT_TIMESTAMP, 'ServiceA', 'ServiceB', 'GET', 'This is a test log entry.');
