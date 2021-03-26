BUILDX_PLATFORMS = linux/amd64,linux/arm64,linux/arm/v7,linux/arm/v6

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test-coverage
test-coverage:
	go test ./... -test.v -coverprofile cp.out

.PHONY: test
test:
	go test ./... -test.v

.PHONY: start
start:
	go run ./cmd/main.go

.PHONY: up
up:
	docker-compose up --build

.PHONY: logs
logs:
	docker-compose logs -f 

.PHONY: build-local
build-local:
	go build ./cmd/main.go

.PHONY: down
down:
	docker-compose down

.PHONY: all
all: lint test build-local

.PHONY: package_docker
package_docker:
	docker build --tag massicer/oh-my-gate-iot-executor:${IMAGE_TAG} .


.PHONY: push_docker
push_docker: package_docker
	docker push massicer/oh-my-gate-iot-executor:${IMAGE_TAG}

.PHONY: cross-build
cross-build:
	docker context create remotecontext
	@docker buildx create --name mybuilder --use
	@docker buildx build --platform ${BUILDX_PLATFORMS} -t ${PROD_IMAGE} --push ./app

.PHONY: export_env_variables
export_env_variables:
	export $(grep -v '^#' .env | xargs)