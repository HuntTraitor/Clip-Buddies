-include ../.env

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run: builds the binary and runs the program
.PHONY: run
run: build
	@./bin/api

## build: builds the binary
.PHONY: build
build:
	@echo 'Building cmd/api...'
	go build -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux/_amd64/api ./cmd/api

## seed: seeds the database
.PHONY: seed
seed:
	@docker compose exec -T db psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -f /tmp/seeds.sql

## db-stats: checks the status of the database to see if you are connected
.PHONY: db-status
db-status:
	@docker compose exec -T backend goose -dir=db/migrations postgres "$(MIGRATION_URL)" status

## up: runs the up migrations
.PHONY: up
up:
	@docker compose exec -T backend goose -dir=db/migrations postgres "$(MIGRATION_URL)" up

## reset: resets the migrations
.PHONY: reset
reset:
	@docker compose exec -T backend goose -dir=db/migrations postgres "$(MIGRATION_URL)" reset

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format all .go files
.PHONY: tidy
tidy:
	@echo 'Formatting .go files...'
	go fmt ./...
	@echo 'Tidying module dependencies...'
	go mod tidy

## audit: run quality control checks
audit:
	@echo 'Checking module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Verifying code...'
	go vet ./...
	staticcheck ./...
