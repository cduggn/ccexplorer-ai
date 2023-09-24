# This file contains convenience targets for the project.
# It is not intended to be used as a build system.

.DEFAULT_GOAL := build

GOPATH ?= $(shell go env GOPATH)
GOBIN ?= $(GOPATH)/bin
GORELEASER ?= $(GOBIN)/goreleaser
GOLANGCI_LINT ?= $(GOBIN)/golangci-lint

$(GORELEASER):
	go install github.com/goreleaser/goreleaser@latest

$(GOLANGCI_LINT):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: setup
setup:
	git config --local core.hooksPath .githooks/

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run --color=always --sort-results ./...

.PHONY: lint-exp
lint-exp:
	$(GOLANGCI_LINT) run --fix --config .golangci-exp.yaml ./...

.PHONY: lint-fix
lint-fix:
	$(GOLANGCI_LINT) run --fix --skip-dirs=./exp ./...

.PHONY: lint-all
lint-all:
	$(GOLANGCI_LINT) run --color=always --sort-results ./...


.PHONY: test-race
test-race:
	go run test -race ./...

.PHONY: test-cover
test-cover:
	go run test -cover ./...

.PHONY: clean
clean: clean-lint-cache

.PHONY: clean-lint-cache
clean-lint-cache:
	$(GOLANGCI_LINT) cache clean

#.PHONY: build
#build:
#	go mod tidy; go build -o .