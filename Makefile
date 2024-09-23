build-postgres:
	docker build -f dockerfiles/postgres.Dockerfile -t dev-db .

.Phony:
docker-build-database:
	docker build -f ./dockerfiles/postgres.Dockerfile . -t us-central1-docker.pkg.dev/small-biz-template/markenshop/blog-db:latest

.Phony:
docker-run-database:
	docker run --rm -d -p 5432:5432 -v blog-db:/var/lib/postgresql/data -t us-central1-docker.pkg.dev/small-biz-template/markenshop/blog-db:latest

.Phony:
migrate-up:
	goose -dir ./repo/postgres/migration "${db_string}" up

.Phony:
migrate-down:
	goose -dir ./repo/postgres/migration "${db_string}" down

migrate-create:
	goose -s -dir ./repo/postgres/migration create $(file_name) sql