GIT_COMMIT=$(shell git rev-parse --short HEAD)
IMAGE=joaofnds/foo
TAG=$(GIT_COMMIT)

.PHONY: test

install-deps:
	go mod download

test:
	go test ./...

build:
	docker build -t $(IMAGE):$(GIT_COMMIT) .

push:
	docker push $(IMAGE):$(GIT_COMMIT)