FROM postgres
ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB world
COPY database/seed.sql /docker-entrypoint-initdb.d/