IMG ?= ghcr.io/kerezsiz42/kinko:latest
CONTAINER_TOOL ?= docker

.PHONY: all
all: test run

.PHONY: build
build: fmt
	go build -o bin/kinko cmd/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: gen
gen:
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: run
run: build
	./bin/kinko

.PHONY: docker-build
docker-build:
	$(CONTAINER_TOOL) build -t ${IMG} .

.PHONY: docker-push
docker-push:
	$(CONTAINER_TOOL) push ${IMG}

.PHONY: deploy
deploy:
	@echo "Implement your deploy script here."