include .env

DB_STRING := "host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSL)"

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

bin-deps: ### Install binary dependencies
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.10.0
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0
.PHONY: bin-deps

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
	docker volume rm $(shell docker volume ls -q -f name=yummy_)
.PHONY: clean

migrate-up: ### Migrate up
	goose -dir ./migrations postgres $(DB_STRING) up
.PHONY: migrate-up

migrate-down: ### Migrate down
	goose -dir ./migrations postgres $(DB_STRING) down
.PHONY: migrate-down

test-data: ### Fetch test data
	docker exec -d postgres mkdir -p /test
	docker cp -q ./test/test_data.sql postgres:/test/test_data.sql
	docker exec postgres psql -q -U $(DB_USER) -d $(DB_NAME) -f /test/test_data.sql
.PHONY: test-data

# TODO Придумать более подходящие название
run-cli: ### Run CLI application
	go run ./cmd/app_cli
.PHONY: run-cli

mock: ### Run mockgen
	mockgen -source ./pkg/postgres/interfaces.go -destination ./pkg/postgres/mock/mock.go
.PHONY: mock

test-cover: ### Run mockgen
	go test -v ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
.PHONY: test-cover

# TODO убрать
reload:
	make migrate-down
	make migrate-up
	make test-data
.PHONY: reload

# TODO Добавить команды для сборки бинарников разных типов
