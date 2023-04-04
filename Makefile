include .env
export

DB_STRING := "host=postgres user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable"

run: ### Run containers
	docker compose up --build -d
.PHONY: run

stop: ### Down containers
	docker compose down --remove-orphans
.PHONY: stop

logs: ### View output from containers
	docker compose logs -f
.PHONY: logs

# TODO Добавить полное удаление всего
clean: ### Remove docker volumes
	docker volume rm $(shell docker volume ls -q -f name=hw-5_)
.PHONY: clean

# TODO Добавить описание
migrate-up:
	docker exec goose \
  goose postgres $(DB_STRING) up
.PHONY: migrate-up

# TODO Добавить описание
migrate-down:
	docker exec goose \
  goose postgres $(DB_STRING) down
.PHONY: migrate-down
