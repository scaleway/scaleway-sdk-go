
# Makefile for scaleway-sdk-go

WORKDIR := $(shell pwd)

# Go versions to test against (matching GitHub Actions)
GO_VERSIONS := 1.23.x 1.24.x

# List all Go packages (excluding vendor and testdata)
PKGS := $(shell go list ./... | grep -v /testdata)

# Build architectures to test
ARCHS := 386 amd64 arm arm64

# Platforms to test
PLATFORMS := windows-latest macos-latest

.PHONY: all build install-dependencies format format-check lint test test-coverage doc prebuild generate-alias generate-packages generate-global-sdk-package publish format-generated clean check-tokens build-arch-test build-platform-test unit-test

all: build

build:
	go build -v ./...

install-dependencies:
	go mod tidy
	go mod download

format:
	go fmt ./...
	./scripts/lint.sh --write

format-check:
	@echo "Checking code formatting..."
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "Files not formatted:"; \
		gofmt -l .; \
		exit 1; \
	fi
	@echo "Code formatting is correct"

lint:
	./scripts/lint.sh

test:
	CGO_ENABLED=0 go test -v -race -cover $(PKGS)

test-coverage:
	CGO_ENABLED=0 go test -coverprofile=coverage.out $(PKGS)
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

generate:
	buf generate --timeout 0
	make format


format-generated:
	./formatting/run.sh -i "./api"

