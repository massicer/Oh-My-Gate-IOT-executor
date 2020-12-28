lint:
	golangci-lint run

test-coverage:
	go test ./... -test.v -coverprofile cp.out

test:
	go test ./... -test.v

start:
	go run ./cmd/main.go

up:
	docker-compose up --build

logs:
	docker-compose logs -f 

build-local:
	go build ./cmd/main.go

down:
	docker-compose down

all: lint test build-local

build_docker_local:
	docker build --tag massicer/oh-my-gate-iot-executor:local .

# DEV
.PHONY: build_docker_dev
build_docker_dev:
	docker build --tag massicer/oh-my-gate-iot-executor:dev .

.PHONY: push_docker_dev
push_docker_dev: build_docker_dev
	docker push massicer/oh-my-gate-iot-executor:dev


# LATEST
.PHONY: build_docker_latest
build_docker_latest:
	docker build --tag massicer/oh-my-gate-iot-executor:latest .

.PHONY: push_docker_latest
push_docker_latest: build_docker_latest
	docker push massicer/oh-my-gate-iot-executor:latest

export_env_variables:
	export $(grep -v '^#' .env | xargs)