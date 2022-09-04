export GO111MODULE := on

EXECUTABLE = httpmockie
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

all: check test build

MAKEFLAGS += --no-print-directory

check:
ifeq (, $(shell which golangci-lint))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.48.0
endif
	golangci-lint run
	go mod tidy

test:
	@echo "Running tests"
	go test -race -v ./...
	go mod tidy

generate:
	@echo "Running generate"
	go generate

build: generate
	@echo "Running build"
	go build -o bin/$(EXECUTABLE)

clean:
	rm -rf bin/ httpmockie
