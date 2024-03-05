# its doesnt stop the containers after running the app
# TODO - add a stop target to stop the containers after running the app

.PHONY: create_network up_db start_db stop_db up_redis start_redis stop_redis run default

NETWORK_NAME=chat_network
PG_CONTAINER_NAME=postgres_db
REDIS_CONTAINER_NAME=redis_db
RUN_CMD=go run .

create_network:
	@if ! docker network ls --format "{{.Name}}" | grep -q $(NETWORK_NAME); then \
		docker network create $(NETWORK_NAME); \
	fi

up_db:
	@if ! docker ps -a --format "{{.Names}}" | grep -q $(PG_CONTAINER_NAME); then \
		docker run --network $(NETWORK_NAME) --name $(PG_CONTAINER_NAME) \
		-e POSTGRES_PASSWORD=$(PG_PASS) \
		-e POSTGRES_DB=$(PG_DB) \
		-p $(PG_PORT):5432 -d postgres; \
	fi

start_db:
	@docker start $(PG_CONTAINER_NAME) || true

stop_db:
	@docker stop $(PG_CONTAINER_NAME) || true

up_redis:
	@if ! docker ps -a --format "{{.Names}}" | grep -q $(REDIS_CONTAINER_NAME); then \
		docker run --network $(NETWORK_NAME) --name $(REDIS_CONTAINER_NAME) -p $(REDIS_PORT):6379 -d redis; \
	fi

start_redis:
	@docker start $(REDIS_CONTAINER_NAME) || true

stop_redis:
	@docker stop $(REDIS_CONTAINER_NAME) || true

run: up_db start_db up_redis start_redis
	@sleep 3
	@$(RUN_CMD)
	@make stop_db
	@make stop_redis

default: up_db start_db up_redis start_redis
	@sleep 3
	@watchexec -e go -r -- $(RUN_CMD)
	@make stop_db
	@make stop_redis
