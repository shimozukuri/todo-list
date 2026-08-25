include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	docker compose up -d todolist-postgres

env-down:
	docker compose down todolist-postgres

env-cleanup:
	@read -p "Clear all volume environment files? Risk of data loss. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
	  docker compose down todolist-postgres && \
	  rm -rf out/pgdata && \
	  echo "Environment files is clear"; \
	else \
	  echo "Environment files cleanup is cancel"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
  		echo "The required parameter 'seq' is missing. Example: make migrate-create seq=init"; \
  		exit 1; \
  	fi; \
	docker compose run --rm todolist-postgres-migrate \
		create \
		-ext sql \
		-dir ./migrations \
		-seq "$(seq)"