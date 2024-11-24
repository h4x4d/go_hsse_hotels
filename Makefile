DOCKER_COMPOSE := $(shell command -v docker-compose >/dev/null 2>&1 && echo docker-compose || echo docker compose)

.PHONY: pull
pull:
	git pull

.PHONY: pull
generate:
	./scripts/generate_from_swagger.sh

.PHONY: update
update: pull
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V

.PHONY: down
down:
	$(DOCKER_COMPOSE) down -v

.PHONY: build-action
build: down
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V

.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V