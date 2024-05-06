# Download MySQL image from Docker Hub
docker pull mysql

# Start a MySQL container and set the MYSQL_ROOT_PASSWORD
docker run -it --name test-world-mysql -e MYSQL_ROOT_PASSWORD=mysql -d mysql:latest

# Login to the MySQL container
docker exec -it test-wolrd-mysql bash -p

# Login to the MySQL container
mysql -u root -p -h 127.0.0.1

# After entering the password, the MySQL shell will start

# Create a database
CREATE DATABASE test_db;

# Use the database
USE test_db;

# Create a table
CREATE TABLE `test_log` (
  `date` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "Date",
  `api` VARCHAR(100) NOT NULL COMMENT "API",
  `error` VARCHAR(200) NOT NULL COMMENT "Error_Sentence"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
