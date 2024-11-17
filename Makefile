CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

.PHONY: run_db run stop_db stop

run_db:
	docker-compose up -d db redis

stop_db:
	docker-compose stop db redis

run:
	docker-compose up --build
	docker compose up

stop:
	docker-compose down

migrate:
	docker-compose run --rm migrate-road

logs:
	docker-compose logs -f

migrate_up:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/pima?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/pima?sslmode=disable -verbose down

migrate_force:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/pima?sslmode=disable -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

insert_file:
	migrate create -ext sql -dir migrations -seq insert_table

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1	