BIN = ow

GO ?= GO111MODULE=on go

deps:
	$(GO) mod download

build: deps
	$(GO) build -o $(BIN) 

.PHONY: deps build
