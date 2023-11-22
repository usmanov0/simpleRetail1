-include .env
export

DOCKER_COMPOSE_FILE=docker-compose.yaml



.PHONY: start
start:
	@echo "Start Containers"
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d ${DOCKER_SERVICES}
	sleep 2
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

.PHONY: stop
stop:
	@echo "Stop Containers"
	docker-compose -f ${DOCKER_COMPOSE_FILE} stop ${DOCKER_SERVICES}
	sleep 2
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

.PHONY: stop
rm: stop
	@echo "Remove Containers"
	docker-compose -f ${DOCKER_COMPOSE_FILE} rm -v -f ${DOCKER_SERVICES}

.PHONY: migrate-generate
migrate-generate:
	@echo "Generating $(name) migrations..."
	docker-compose run --rm migrate create -seq -ext=.sql -dir=./migrations $(name)

.PHONY: migrate-up
migrate-up:
	@echo "Migrating up..."
	docker-compose run --rm migrate -path=./migrations -database='${PROJECT_DSN}' upS_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' up

.PHONY: migrate-down
migrate-down:
	@echo "Migrating down..."
	docker-compose run --rm migrate -path=./migrations -database='${PROJECT_DSN}' down
