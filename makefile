GOCMD:=$(shell which go)
GOINSTALL:=$(GOCMD) install
GOBUILD:=$(GOCMD) build

BINARY_NAME:=goflux

all: build install
.PHONY: all

build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/goflux

install:
	$(GOINSTALL) ./cmd/goflux