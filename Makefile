include .env

DB_STRING := "host=$(DB_HOST) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=$(DB_SSL)"
PB_PATH := "internal/pkg/api"

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

get-protoc:	### Install protoc and Go plugins
	curl -LO "https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip"
	unzip protoc-3.15.8-linux-x86_64.zip -d "${HOME}/.local"
	export PATH="${PATH}:${HOME}/.local/bin"
	rm protoc-3.15.8-linux-x86_64.zip
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

bin-deps: get-protoc ### Install binary dependencies
	go install github.com/pressly/goose/v3/cmd/goose@v3.10.0
	go install github.com/golang/mock/mockgen@v1.6.0
	export PATH="${PATH}:$(go env GOPATH)/bin"
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
	docker cp -q ./test/sql/test_data.sql postgres:/test/test_data.sql
	docker exec postgres psql -q -U $(DB_USER) -d $(DB_NAME) -f /test/test_data.sql
.PHONY: test-data

run: ### Run app
	go run ./cmd/app_cli
	# TODO
.PHONY: run-cli

mocks: ### Run mockgen
	mockgen -source ./internal/app/menu/repo/interfaces.go -destination ./test/mocks/menu_repo.go -package mocks
.PHONY: mocks

test-unit: ### Run unit tests
	go clean -testcache
	go test -v ./...
.PHONY: test-unit

test-integration: ### Run integration tests
	go clean -testcache
	go test -tags=integration -v ./...
.PHONY: test-integration

test-cover: ### Run tests with coverage
	go clean -testcache
	go test -tags=all,integration -v -coverpkg=./... ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
.PHONY: test-cover


generate: ### Generate gRPC code
	if [ -d $(PB_PATH) ]; then rm -rf $(PB_PATH); fi
	mkdir -p $(PB_PATH)
	protoc \
  --proto_path=api/proto/ \
  --go_out=$(PB_PATH) \
  --go-grpc_out=$(PB_PATH) \
  api/proto/*.proto
 .PHONY: generate

# TODO: remove
reload:
	make migrate-down
	make migrate-up
	make test-data
.PHONY: reload

# TODO: add command for building different types of binaries
