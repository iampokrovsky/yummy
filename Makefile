include .env

export MAKEFLAGS="-s"

DB_STRING := "host=postgres user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSL)"

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ### Run application
	go run ./cmd/app -t menu -a create -d '{"restaurantID": 10, "name": "Chicken", "price": 150000}'
	#go run ./cmd/app -t menu -a get -d '{"id": 101}'
	#go run ./cmd/app -t menu -a list -d '{"restaurantID": 10}'
	#go run ./cmd/app -t menu -a update -d '{"id": 101, "name": "Potato", "price": 250000}'
	#go run ./cmd/app -t menu -a delete -d '{"id": 101}'
	#go run ./cmd/app -t menu -a restore -d '{"id": 101}'
	#go run ./cmd/app -t restaurant -a create -d '{"name": "Mama Mia", "address": "Moscow", "cuisine": "Italian"}'
	#go run ./cmd/app -t restaurant -a get -d '{"id": 21}'
	#go run ./cmd/app -t restaurant -a list -d '{}'
	#go run ./cmd/app -t restaurant -a update -d '{"id": 21, "name": "Dolmio", "address": "Voronezh", "cuisine": "Greek"}'
	#go run ./cmd/app -t restaurant -a delete -d '{"id": 21}'
	#go run ./cmd/app -t restaurant -a restore -d '{"id": 21}'
.PHONY: run

compose-up: ### Run containers
	docker compose up --build -d
.PHONY: start

compose-down: ### Stop containers
	docker compose down --remove-orphans
.PHONY: stop

compose-logs: ### View output from containers
	docker compose logs -f
.PHONY: logs

rm-volumes: ### Remove docker volumes
	docker volume rm $(shell docker volume ls -q -f name=hw-5_)
.PHONY: clean

migrate-up: ### Migrate up
	docker exec goose \
  goose postgres $(DB_STRING) up
.PHONY: migrate-up

migrate-down: ### Migrate down
	docker exec goose \
  goose postgres $(DB_STRING) down
.PHONY: migrate-down

test-data: ### Fetch test data
	docker exec -d postgres mkdir -p /test
	docker cp -q ./test/test_data.sql postgres:/test/test_data.sql
	docker exec postgres psql -q -U $(DB_USER) -d $(DB_NAME) -f /test/test_data.sql
.PHONY: test-data

# TODO убрать
reload:
	make migrate-down
	make migrate-up
	make test-data
.PHONY: reload
