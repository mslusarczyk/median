.DEFAULT_GOAL: all

.PHONY: all test build fmt ci

all: fmt build test

fmt:
	@echo "Formating source code"
	goimports -l -w .

test:
	@echo "Running tests"
	go test -v ./... && echo "TESTS PASSED"

ci: build test

build:
	@echo "Building sources"
	go build -v ./...
