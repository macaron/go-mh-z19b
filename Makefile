NAME 		= mhz19b
GOOS		= linux
GOARCH		= arm
GOARM 		= 7
REVISION 	= $(shell git rev-parse --short HEAD)
LDFLAGS 	= -ldflags="-s -w -X main.revision=$(REVISION)"

export GO111MODULES=ON

## Setup
.PHONY: deps
devel-deps: deps
	go get \
		github.com/Songmu/make2help/cmd/make2help

## Lint
.PHONY: lint
lint: devel-deps
	go vet ./...

## Build binary
.PHONY: build
build:
	go mod tidy && \
	GOOS=$(GOOS) \
	GOARCH=$(GOARCH) \
	GOARM=$(GOARM) \
	go build $(LDFLAGS) -o $(NAME) ./cmd/mhz19b/main.go

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
