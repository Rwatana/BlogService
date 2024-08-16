# MySQL 8.0をベースにする
FROM mysql:latest

# 環境変数を設定する（オプション）
ENV MYSQL_ROOT_PASSWORD=mysql
COPY create_db.sql /docker-entrypoint-initdb.d/

# ポートを公開する（オプション）
EXPOSE 3306
