.PHONY: update
update:
	git pull

.PHONY: down
down:
	docker compose down -v

.PHONY: build-action
build: down
	docker compose up -d --build

.PHONY: up
up:
	docker compose up -d --build