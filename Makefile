include .env
export


export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d todolist-postgres

env-down:
	@docker compose down todolist-postgres

env-cleanup:
	@read -p "Clear all volume environment files? Risk of data loss. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
	  docker compose down todolist-postgres port-forwarder && \
	  rm -rf ${PROJECT_ROOT}/out/pgdata && \
	  echo "Environment files is clear"; \
	else \
	  echo "Environment files cleanup is cancel"; \
	fi

env-port-forwarder:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
  		echo "The required parameter 'seq' is missing. Example: make migrate-create seq=init"; \
  		exit 1; \
  	fi; \
	docker compose run --rm todolist-postgres-migrate \
		create \
		-ext sql \
		-dir ${PROJECT_ROOT}/migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
  		echo "The required parameter 'action' is missing. Example: make migrate-action action=up"; \
  		exit 1; \
  	fi; \
	docker compose run --rm todolist-postgres-migrate \
        		-path /migrations \
        		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todolist-env-postgres:5432/${POSTGRES_DB}?sslmode=disable \
        		"$(action)"

todolist-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todo-list/main.go