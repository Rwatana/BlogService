CREATE TABLE IF NOT EXISTS `logging` (
    `log_level` VARCHAR(100) NOT NULL COMMENT "Log_Level",
    `date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "Date",
    `current_service` VARCHAR(100) NOT NULL COMMENT "Current_API",
    `source_service` VARCHAR(100) NOT NULL COMMENT "Source_API",
    `type_of_request` VARCHAR(100) NOT NULL COMMENT "Request_Type",
    `content` VARCHAR(200) NOT NULL COMMENT "Error_Sentence"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- change the column name from `date` to `log_date`
ALTER TABLE logging MODIFY COLUMN `date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT "Date";
