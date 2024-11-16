.PHONY: run_db run stop_db stop

# Start database services
run_db:
	docker-compose up -d db redis

# Stop database services
stop_db:
	docker-compose stop db redis

# Start the entire application
run:
	docker-compose up --build

# Stop the entire application
stop:
	docker-compose down

CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

run:
	go run cmd/main.go
init:
	go mod init
	go mod tidy 
	go mod vendor

proto-gen:
	./internal/scripts/gen-proto.sh ${CURRENT_DIR}
	
migrate_up:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/tender?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/tender?sslmode=disable -verbose down

migrate_force:
	migrate -path migrations -database postgres://postgres:1111@localhost:5432/tender?sslmode=disable -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

insert_file:
	migrate create -ext sql -dir migrations -seq insert_table

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

lint: ## Run golangci-lint with printing to stdout
	golangci-lint -c .golangci.yaml run --build-tags "musl" ./...

swag-init:
	swag init -g api/api.go -o api/docs

swag-gen:
	~/go/bin/swag init -g ./api/api.go -o api/docs force 1