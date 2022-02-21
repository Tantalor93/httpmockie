export GO111MODULE := on

EXECUTABLE = go-http-mock
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

all: check test build

MAKEFLAGS += --no-print-directory

prepare:
	@echo "Downloading tools"
ifeq (, $(shell which go-junit-report))
	go get github.com/jstemmer/go-junit-report
endif
ifeq (, $(shell which gocov))
	go get github.com/axw/gocov/gocov
endif
ifeq (, $(shell which gocov-xml))
	go get github.com/AlekSi/gocov-xml
endif

check:
ifeq (, $(shell which golangci-lint))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.30.0
endif
	golangci-lint run
	go mod tidy

test:
	@echo "Running tests"
	mkdir -p report
	go test -race -v ./... -coverprofile=report/coverage.txt | tee report/report.txt
	gocov convert report/coverage.txt | gocov-xml > report/coverage.xml
	go mod tidy

generate:
	@echo "Running generate"
	go generate

build: generate
	@echo "Running build"
	go build -o bin/$(EXECUTABLE)
