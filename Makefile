DOCKER_COMPOSE := $(shell command -v docker-compose >/dev/null 2>&1 && echo docker-compose || echo docker compose)

.PHONY: pull
pull:
	git pull

.PHONY: swagger_generate
swagger_generate:
	./scripts/generate_from_swagger.sh

.PHONY: grpc_generate
grpc_generate:
	protoc -I api/proto api/proto/*.proto \
	  --go_out=hotel/internal/grpc/gen \
	  --go_opt=paths=source_relative \
	  --go_opt=Mhotel.proto=github.com/h4x4d/go_hsse_hotels/hotel/internal/grpc/gen \
	  --go-grpc_out=hotel/internal/grpc/gen \
	  --go-grpc_opt=paths=source_relative \
	  --go-grpc_opt=Mhotel.proto=github.com/h4x4d/go_hsse_hotels/hotel/internal/grpc/gen

	protoc -I api/proto api/proto/*.proto \
	  --go_out=booking/internal/grpc/gen \
	  --go_opt=paths=source_relative \
	  --go_opt=Mhotel.proto=github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/gen \
	  --go-grpc_out=booking/internal/grpc/gen \
	  --go-grpc_opt=paths=source_relative \
	  --go-grpc_opt=Mhotel.proto=github.com/h4x4d/go_hsse_hotels/booking/internal/grpc/gen

.PHONY: codegen
codegen: grpc_generate swagger_generate

.PHONY: update
update: pull
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V --remove-orphans

.PHONY: down
down:
	$(DOCKER_COMPOSE) down -v --remove-orphans

.PHONY: build-action
build: down
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V --remove-orphans

.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d --build --force-recreate -V --remove-orphans

.PHONY: ps
ps:
	$(DOCKER_COMPOSE) ps -a