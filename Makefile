BIN         = $(GOPATH)/bin
GOLINT      = $(BIN)/golint
SRC         = $(shell find . -type f -name '*.go' -not -path "./vendor/*" )
export GO111MODULE = on

prepare:
	@echo "Downloading tools"
	go get -u github.com/jstemmer/go-junit-report
	go get golang.org/x/lint/golint

check: prepare
	gofmt -s -d -e $(SRC)
	@test -z $(shell gofmt -l ${SRC} | tee /dev/stderr)
	$(GOLINT) -set_exit_status $$(go list ./...)
	go vet  ./...

test: prepare
	@echo "Running tests"
	mkdir -p report
	go test -v ./... | tee report/report.txt
	go-junit-report -set-exit-code < report/report.txt > report/report.xml


build:
	@echo "Running build"
	CGO_ENABLED=0 go build -o bin/go-http-mock