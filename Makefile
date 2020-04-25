export PROJECT = viperconf

.DEFAULT_GOAL := help

up: ## Builds, (re)creates, starts, and attaches to Docker containers for a service.
	docker-compose up

down: ## Stops Docker containers and removes containers, networks, volumes, and images created by up.
	docker-compose down

keys: ## Generate private key file to private.pem file
	go run ./cmd/restaurant-admin/main.go keygen private.pem

admin: ## Adds default admin user in DB.
	go run ./cmd/restaurant-admin/main.go --db-disable-tls=1 useradd admin@example.com gophers

migrate: ## Migrate attempts to bring the schema for db up to date with the migrations defined.
	go run ./cmd/restaurant-admin/main.go --db-disable-tls=1 migrate

seed: migrate ## Seed runs the set of seed-data queries against db. The queries are ran in a transaction and rolled back if any fail.
	go run ./cmd/restaurant-admin/main.go --db-disable-tls=1 seed

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
