#!make
POSTGRESQL_URL := postgres://postgres:postgres@localhost:5433/master?sslmode=disable

### GO tools
tools:
	cd tools && go mod tidy && go mod vendor && go mod verify && go generate -tags tools
.PHONY: tools	

vendor:
	go mod tidy && go mod vendor && go mod verify
.PHONY: vendor	

build:
	go build -o ./bin/service ./cmd/service
.PHONY: build	

fmt:
	go vet ./cmd/... && fmt ./internal/...

vet:
	go vet ./cmd/... && fmt ./internal/...

imports:
	bin/goimports -local github.com/lidofinance/mev-boost-monitoring -w -d $(shell find . -type f -name '*.go'| grep -v "/vendor/\|/.git/\|/tools/")

lint:
	bin/golangci-lint run --config=.golangci.yml --fix ./...

full-lint: imports fmt vet lint	
.PHONY: full-lint	

full-lint: imports fmt vet lint
.PHONY: full-lint

### Migrations
.PHONY: rollback
rollback:
	bin/migrate -database ${POSTGRESQL_URL} -path db/migrations down

.PHONY: migrate
migrate:
	bin/migrate -database ${POSTGRESQL_URL} -path db/migrations up

.PHONY: up
up:
	UID_GID="$(id -u):$(id -g)" docker-compose -f docker-compose.yml up -d

.PHONY: up-rebuild
up-rebuild:
	UID_GID="$(id -u):$(id -g)" docker-compose -f docker-compose.yml up -d --build mev-boost-collector


.PHONY: down
down:
	UID_GID="$(id -u):$(id -g)" docker-compose -f docker-compose.yml down


.PHONY: run
run:
	make up && make migrate