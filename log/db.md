# Docker HubからMySQLのイメージをダウンロードする
docker pull mysql

# MySQLのコンテナを起動し、MYSQL_ROOT_PASSWORDにパスワードを設定する
docker run -it --name test-world-mysql -e MYSQL_ROOT_PASSWORD=mysql -d mysql:latest

# MySQLのコンテナにログインする
docker exec -it test-world-mysql mysql -u root -p

# パスワードを入力後、MySQLのシェルが起動する

# データベースを作成する
CREATE DATABASE test_db;

# データベースを使用する
USE test_db;

# テーブルを作成する
CREATE TABLE `test_log` (
  `date` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "Date",
  `api` VARCHAR(100) NOT NULL COMMENT "API",
  `error` VARCHAR(200) NOT NULL COMMENT "Error_Sentence"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
