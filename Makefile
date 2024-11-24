DOCKER_COMPOSE := $(shell command -v docker-compose >/dev/null 2>&1 && echo docker-compose || echo docker compose)

.PHONY: pull
pull:
	git pull

.PHONY: swagger_generate
swagger_generate:
	./scripts/generate_from_swagger.sh

.PHONY: grpc_generate
grpc_generate:
	protoc -I hotel/api/proto hotel/api/proto/*.proto --go_out=hotel/internal/grpc/gen --go_opt=paths=source_relative  --go-grpc_out=hotel/internal/grpc/gen --go-grpc_opt=paths=source_relative
	protoc -I booking/api/proto booking/api/proto/*.proto --go_out=booking/internal/grpc/gen --go_opt=paths=source_relative  --go-grpc_out=booking/internal/grpc/gen --go-grpc_opt=paths=source_relative

.PHONY: codegen
codegen: grpc_generate swagger_generate

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