include .env

export MAKEFLAGS="-s"

DB_STRING := "host=postgres user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable"

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

start: ### Run containers
	docker compose up --build -d
.PHONY: start

stop: ### Down containers
	docker compose down --remove-orphans
.PHONY: stop

logs: ### View output from containers
	docker compose logs -f
.PHONY: logs

clear: ### Remove docker volumes
	docker volume rm $(shell docker volume ls -q -f name=hw-5_)
.PHONY: clear

migrate-up: ### Migrate up
	docker exec goose \
  goose postgres $(DB_STRING) up
.PHONY: migrate-up

migrate-down: ### Migrate down
	docker exec goose \
  goose postgres $(DB_STRING) down
.PHONY: migrate-down

test-data: ### Fetch test data
	docker exec -d postgres mkdir -p /test_data
	docker cp -q ./tests/test_data.sql postgres:/test_data/test_data.sql
	docker exec postgres psql -q -U $(DB_USER) -d $(DB_NAME) -f /test_data/test_data.sql
.PHONY: test-data

###

# TODO убрать
reload:
	make migrate-down
	make migrate-up
	make test-data
.PHONY: reload
