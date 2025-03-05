FROM mysql:9.2

ENV MYSQL_ROOT_PASSWORD=1234

ADD ./test_db-master/ /docker-entrypoint-initdb.d/

EXPOSE 3307

CMD ["mysqld"]