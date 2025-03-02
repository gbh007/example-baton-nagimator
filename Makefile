.PHONY: run
run:
	go run -trimpath cmd/main.go

.PHONY: up
up:
	docker compose -f deployment/docker-compose.yaml up --build -d --remove-orphans

.PHONY: down
down:
	docker compose -f deployment/docker-compose.yaml down --remove-orphans

.PHONY: restart
restart:
	docker compose -f deployment/docker-compose.yaml restart

.PHONY: jsonnet-build
jsonnet-build:
	cd jsonnet && make build
