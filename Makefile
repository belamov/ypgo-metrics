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

mock: ## Generate mocks

proto: ## Generate proto files
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) protoc --go_out=. --go_opt=paths=source_relative \
                                                                               --go-grpc_out=. --go-grpc_opt=paths=source_relative

lint:
	$(docker_bin) run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run

fieldaligment-fix:
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) fieldalignment -fix ./... || true

gofumpt:
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) gofumpt -l -w .

test: ## Execute tests
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) go test -v -race ./...

fresh-itest:
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) curl -s https://api.github.com/repos/Yandex-Practicum/go-autotests/releases/latest | grep 'http.\+metricstest"' | cut -d : -f 2,3 | tr -d \" | wget -v --show-progress -O metricstest && chmod +x ./metricstest

itest: build
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) go build -v -o ./cmd/server ./cmd/server
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) go build -v -o ./cmd/agent ./cmd/agent
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm -it $(app_container_name) ./metricstest -test.v -test.run=^TestIteration1$$ -binary-path=cmd/server/server -agent-binary-path=cmd/agent/agent

check: build fieldaligment-fix gofumpt lint test itest  ## Run tests and code analysis

staticlint:
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) go build -v -o /usr/src/app/cmd/staticlint/staticlint /usr/src/app/cmd/staticlint
	$(docker_compose_bin) --file "$(docker_compose_yml)" run --rm $(app_container_name) /usr/src/app/cmd/staticlint/staticlint ./...

# Prompt to continue
prompt-continue:
	@while [ -z "$$CONTINUE" ]; do \
		read -r -p "Would you like to continue? [y]" CONTINUE; \
	done ; \
	if [ ! $$CONTINUE == "y" ]; then \
        echo "Exiting." ; \
        exit 1 ; \
    fi
