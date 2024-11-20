.PHONY: pull
pull:
	git pull


.PHONY: update
update: pull
	docker compose up -d --build --force-recreate -V

.PHONY: down
down:
	docker compose down -v

.PHONY: build-action
build: down
	docker compose up -d --build --force-recreate -V

.PHONY: up
up:
	docker compose up -d --build --force-recreate -V