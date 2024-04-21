mysql > CREATE TABLE `test_log` (
          `date` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "Date",
          `api` VARCHAR(100) NOT NULL COMMENT "API",
          `error` VARCHAR(200) NOT NULL COMMENT "Error_Sentence",
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
