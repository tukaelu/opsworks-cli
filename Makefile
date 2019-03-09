BIN = ow

VERSION = 0.1.0
CURRENT_REVISION = $(shell git rev-parse --short HEAD)

REPOSITORY = github.com/tukaelu/opsworks-cli
CMD_PACKAGE = $(REPOSITORY)/cmd

GO ?= GO111MODULE=on go

deps:
	$(GO) mod download

build: deps
	$(GO) build -ldflags "-w -s -X $(CMD_PACKAGE).version=$(VERSION) -X $(CMD_PACKAGE).revision=$(CURRENT_REVISION)" -o $(BIN)

.PHONY: deps build
