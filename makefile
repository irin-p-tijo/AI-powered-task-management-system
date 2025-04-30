SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build


build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR)/app -v ./cmd


${BINARY_DIR}:
	mkdir -p ${BINARY_DIR}

run: ## Start application
	$(GOCMD) run cmd/main.go

