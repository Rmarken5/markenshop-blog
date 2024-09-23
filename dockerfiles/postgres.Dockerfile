FROM postgres:16.4-alpine3.20

ENV POSTGRES_USER=marken
ENV POSTGRES_PASSWORD=shop
ENV POSTGRES_DB=blog
ENV PGDATA /var/lib/postgresql/data

VOLUME $PGDATA

EXPOSE 5432

CMD ["postgres"]