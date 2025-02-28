include deploy/develop/.db.env
include .env
export

export GOPRIVATE="github.com/geniusrabbit/*"
APP_BUILD_TAGS ?= postgres,clickhouse,migrate,redis,jaeger
# doctl kubernetes cluster kubeconfig save use_your_cluster_name

include deploy/build.mk

PROJECT_WORKSPACE ?= ssp-project
PROJECT_NAME ?= api
DOCKER_COMPOSE := docker compose -p $(PROJECT_WORKSPACE) -f deploy/develop/docker-compose.yml
DOCKER_CONTAINER_IMAGE := ${PROJECT_WORKSPACE}/${PROJECT_NAME}

.PHONY: all
all: lint cover

.PHONY: lint
lint: golint ## Run linters

.PHONY: golint
golint:
	# golint -set_exit_status ./...
	golangci-lint run -v ./...

.PHONY: fmt
fmt: ## Run formatting code
	@echo "Fix formatting"
	@gofmt -w ${GO_FMT_FLAGS} $$(go list -f "{{ .Dir }}" ./...); if [ "$${errors}" != "" ]; then echo "$${errors}"; fi

.PHONY: fixi
fixi:
	 @echo "Fix imports $(shell go list -m)"
	 goimports -local="$(shell go list -m)" -w .

.PHONY: test
test: ## Run unit tests
	go test -v -tags "${APP_TAGS}" -race ./...

.PHONY: tidy
tidy: ## Run go mod tidy
	go mod tidy

.PHONY: vendor
vendor: ## Run go mod vendor
	go mod vendor

.PHONY: generate-code
generate-code: ## Run codegeneration procedure
	@echo "Generate code"
	@go generate ./...

.PHONY: build-gql
build-gql: ## Build graphql server
	# cd protocol/graphql && go run github.com/99designs/gqlgen
	cd protocol/graphql && gqlgen

.PHONY: build
build: ## Build API application
	@echo "Build application"
	@rm -rf .build
	@$(call do_build,"cmd/api/main.go",api)

.PHONY: build-docker-dev
build-docker-dev: build ## Build docker image for development
	echo "Build develop docker image"
	DOCKER_BUILDKIT=${DOCKER_BUILDKIT} docker build \
		--build-arg TARGETPLATFORM=${LOCAL_TARGETPLATFORM} \
		-t ${DOCKER_CONTAINER_IMAGE} \
		-f deploy/develop/Dockerfile .

.PHONY: run
run: build-docker-dev ## Run API service by docker-compose
	@echo "Run API service http://localhost:${DOCKER_SERVER_HTTP_PORT}"
	$(DOCKER_COMPOSE) up api

.PHONY: stop
stop: ## Stop all services
	@echo "Stop all services"
	$(DOCKER_COMPOSE) stop

.PHONY: db-cli
db-cli: ## Open development database
	$(DOCKER_COMPOSE) exec $(DOCKER_DATABASE_NAME) psql -U $(DATABASE_USER) $(DATABASE_DB)

.PHONY: dbdump
dbdump: ## Dump development database
	$(DOCKER_COMPOSE) exec $(DOCKER_DATABASE_NAME) pg_dump -U $(DATABASE_USER) $(DATABASE_DB)

.PHONY: ch-cli
ch-cli: ## Connect to dev clickhouse
	$(DOCKER_COMPOSE) exec clickhouse-server clickhouse-client

.PHONY: chidump
chidump:
	@$(DOCKER_COMPOSE) exec clickhouse-server clickhouse-client --query="SELECT * FROM stats.events_local" --format SQLInsert
	$(DOCKER_COMPOSE) up api

.PHONY: import-clickhouse-dump
import-clickhouse-dump:
	$(DOCKER_COMPOSE) up clickhouse-dump --remove-orphans

.PHONY: import-postgres-dump
import-postgres-dump:
	$(DOCKER_COMPOSE) up database-dump --remove-orphans

.PHONY: init-submodules
init-submodules: ## Init submodules
	git submodule update --init --recursive

.PHONY: pull-submodules
pull-submodules: ## Pull submodules
	git submodule update --recursive --remote

.PHONY: reset-dev-env
reset-dev-env: ## Reset dev environment
	@${DOCKER_COMPOSE} down -v --rmi all

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
