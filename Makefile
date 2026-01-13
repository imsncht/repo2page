# repo2page Makefile

BINARY_NAME=repo2page
VERSION=$(shell git describe --tags --always --dirty || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD || echo "unknown")
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS=-ldflags "-X 'repo2page/internal/version.Version=${VERSION}' -X 'repo2page/internal/version.Commit=${COMMIT}' -X 'repo2page/internal/version.Date=${DATE}'"

.PHONY: all build clean test run install fmt vet

all: build

build:
	@echo "Building ${BINARY_NAME}..."
	go build ${LDFLAGS} -o bin/${BINARY_NAME} ./cmd/repo2page

clean:
	@echo "Cleaning..."
	go clean
	rm -rf bin/

test:
	@echo "Running tests..."
	go test ./...

run:
	@go run ./cmd/repo2page

install:
	@echo "Installing..."
	go install ${LDFLAGS} ./cmd/repo2page

fmt:
	@echo "Formatting..."
	go fmt ./...

vet:
	@echo "Vetting..."
	go vet ./...
