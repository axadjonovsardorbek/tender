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

swag-init:
	swag init -g api/api.go -o api/docs

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1