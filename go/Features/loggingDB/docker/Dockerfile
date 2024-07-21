FROM mysql:8.0

# set environment variables
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=logging
ENV MYSQL_USER=user
ENV MYSQL_PASSWORD=password

# copy the database schema to the /docker-entrypoint-initdb.d directory
COPY ./db/migrations /docker-entrypoint-initdb.d/

# set the entrypoint script
ENTRYPOINT ["docker-entrypoint.sh"]

# start the mysql server
CMD ["mysqld"]
