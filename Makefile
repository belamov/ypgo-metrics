#!/usr/bin/make
# Makefile readme (ru): <http://linux.yaroslavl.ru/docs/prog/gnu_make_3-79_russian_manual.html>
# Makefile readme (en): <https://www.gnu.org/software/make/manual/html_node/index.html#SEC_Contents>

SHELL = /bin/sh

app_container_name := app
docker_bin := $(shell command -v docker 2> /dev/null)
docker_compose_bin := docker compose
docker_compose_yml := docker/docker-compose.yml
user_id := $(shell id -u)

.PHONY : help pull build push login test clean \
         app-pull app app-push\
         sources-pull sources sources-push\
         nginx-pull nginx nginx-push\
         up down restart shell install
.DEFAULT_GOAL := help

# --- [ Development tasks ] -------------------------------------------------------------------------------------------
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

build: ## Build containers
	$(docker_compose_bin) --file "$(docker_compose_yml)" build

up: build ## Run app
	$(docker_compose_bin) --file "$(docker_compose_yml)" up

doc: build ## Run local documentation
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm -p 8090:8090 $(app_container_name) godoc -http=:8090 -goroot="/usr"
	## "http://localhost:8090/pkg/?m=all"

mock: build ## Generate mocks
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) mockgen -destination=internal/app/mocks/repository.go -package=mocks github.com/belamov/ypgo-metrics/internal/app/storage Repository
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) mockgen -destination=internal/app/mocks/service.go -package=mocks github.com/belamov/ypgo-metrics/internal/app/services MetricServiceInterface

proto: ## Generate proto files
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) protoc --go_out=. --go_opt=paths=source_relative \


shell:
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm -it $(app_container_name) sh

lint:
	$(docker_bin) run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run

fieldaligment-fix:
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) fieldalignment -fix ./... || true

gofumpt:
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) gofumpt -l -w .

test: build ## Execute tests
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) go test -v -race ./...

itest:
	## download metrictest binary from latest release https://github.com/Yandex-Practicum/go-autotests/releases/ first
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) go build -v -o ./cmd/server ./cmd/server
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) go build -v -o ./cmd/agent ./cmd/agent
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm -it $(app_container_name) ./metricstest -test.v -test.run=^TestIteration[12345678][ABCDEFGHIGKLMNOPQRSTUVW]*$$ -binary-path=cmd/server/server -agent-binary-path=cmd/agent/agent  -source-path=. -server-port=8088

istatictest:
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) go vet -vettool=statictest ./...

check: build fieldaligment-fix gofumpt istatictest lint test itest   ## Run tests and code analysis

# Prompt to continue
prompt-continue:
	@while [ -z "$$CONTINUE" ]; do \
		read -r -p "Would you like to continue? [y]" CONTINUE; \
	done ; \
	if [ ! $$CONTINUE == "y" ]; then \
        echo "Exiting." ; \
        exit 1 ; \
    fi
