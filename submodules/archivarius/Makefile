include .env
export

APP_BUILD_TAGS ?= postgres,migrate,clickhouse,redis,jaeger
K8C := kubectl --cluster=do-ams3-geniusrabbit

include deploy/build.mk

PROJECT_WORKSPACE ?= adnet-project
PROJECT_NAME ?= archivarius
DOCKER_COMPOSE := docker compose -p "archivarius" -f deploy/develop/docker-compose.yml
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

.PHONY: cover
cover:
	@mkdir -p $(TMP_ETC)
	@rm -f $(TMP_ETC)/coverage.txt $(TMP_ETC)/coverage.html
	go test -race -coverprofile=$(TMP_ETC)/coverage.txt -coverpkg=./... ./...
	@go tool cover -html=$(TMP_ETC)/coverage.txt -o $(TMP_ETC)/coverage.html
	@echo
	@go tool cover -func=$(TMP_ETC)/coverage.txt | grep total
	@echo
	@echo Open the coverage report:
	@echo open $(TMP_ETC)/coverage.html

.PHONY: generate-pb
generate-pb: ## Generate protobuf files
	protoc -I=. --go_out=internal/server --go-grpc_out=internal/server protocol/archivarius.proto

.PHONY: __eval_srcs
__eval_srcs:
	$(eval SRCS := $(shell find . -not -path 'bazel-*' -not -path '.tmp*' -name '*.go'))

.PHONY: generate-code
generate-code: ## Run codegeneration procedure
	@echo "Generate code"
	@go generate ./...

.PHONY: build
build: ## Build application
	@echo "Build application"
	@rm -rf .build
	@$(call do_build,"cmd/main.go",archivarius)

.PHONY: build-docker-dev
build-docker-dev: build ## Build docker image for development
	echo "Build develop docker image"
	DOCKER_BUILDKIT=${DOCKER_BUILDKIT} docker build -t ${DOCKER_CONTAINER_IMAGE} -f deploy/develop/Dockerfile .

.PHONY: run-archivarius
run-archivarius: build-docker-dev ## Run archivarius service by docker-compose
	@echo "Run Archivarius service http://localhost:${DOCKER_SERVER_HTTP_PORT}"
	$(DOCKER_COMPOSE) up archivarius --remove-orphans

.PHONY: cleanup-archivarius
cleanup-archivarius: ## Remove containers and volumes
	-docker ps | grep "archivarius" | awk '{print $$1}' | xargs --no-run-if-empty docker kill
	-docker ps -a  | grep "archivarius" | awk '{print $$1}' | xargs --no-run-if-empty docker rm
	-docker volume ls  | grep "archivarius" | awk '{print $$2}' | xargs --no-run-if-empty docker volume rm

.PHONY: stop
stop: ## Stop all services
	@echo "Stop all services"
	$(DOCKER_COMPOSE) stop

.PHONY: import-clickhouse-dump
import-clickhouse-dump:
	$(DOCKER_COMPOSE) up clickhouse-dump --remove-orphans

.PHONY: dbcli
dbcli: ## Open development database
	$(DOCKER_COMPOSE) exec $(DOCKER_DATABASE_NAME) psql -U $(DATABASE_USER) $(DATABASE_DB)

.PHONY: chin
chin: ## Connect to dev clickhouse
	$(DOCKER_COMPOSE) exec clickhouse-server clickhouse-client

.PHONY: chink
chin8: ## Connect to prod clickhouse
	${K8C} -n bidswitch-statistic exec -it chi-statistic-statistic-0-0-0 -- clickhouse-client

.PHONY: dbcli8
dbcli8: ## Connect to prod database
	${K8C} -n bidswitch-database exec -it `${K8C} get pods -n bidswitch-database -l app=database -o "jsonpath={.items[0].metadata.name}"` -- psql -U bidsw-unm-db adnet-bidswitch-main

.PHONY: es8
es8: ## Eventstram logs
	${K8C} -n bidswitch logs --follow `${K8C} get pods -n bidswitch -l app=eventstream -o "jsonpath={.items[0].metadata.name}"`

# https://github.com/kubernetes/dashboard
.PHONY: k8dash
k8dash: ## Run k8s dashboard
	kubectl -n kubernetes-dashboard port-forward svc/kubernetes-dashboard-kong-proxy 8443:443

.PHONY: k8s-dev-token
k8s-dev-token: ## Get kube dashboard token
	kubectl apply -f deploy/develop/k8s/user.yaml
	kubectl -n kubernetes-dashboard create token admin-user --duration=8760h

.PHONY: init-submodules
init-submodules: ## Init submodules
	git submodule update --init --recursive

.PHONY: pull-submodules
pull-submodules: ## Pull submodules
	git submodule update --recursive --remote

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help