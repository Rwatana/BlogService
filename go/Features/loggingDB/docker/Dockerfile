# MySQL 8.0をベースにする
FROM mysql:latest

# 環境変数を設定する（オプション）
ENV MYSQL_ROOT_PASSWORD=mysql

# run create_db.sql file
RUN souece create_db.sql
# ポートを公開する（オプション）
EXPOSE 3306
