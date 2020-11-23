FROM postgres

COPY init_db.sql /docker-entrypoint-initdb.d/init_db.sql

ARG postgres_user
ENV POSTGRES_USER=$postgres_user

ARG postgres_password
ENV POSTGRES_PASSWORD=$postgres_password

ARG postgres_db
ENV POSTGRES_DB=$postgres_db

EXPOSE 5432
