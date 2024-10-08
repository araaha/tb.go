BINARY_NAME := tb
BIN_DIR := /usr/local/bin
SRCS := $(shell git ls-files '*.go')
LDFLAGS := "-X github.com/araaha/tb.go/cmd.Version=$(shell git describe --tags --abbrev=0 --always) -X github.com/araaha/tb.go/cmd.Revision=$(shell git rev-parse --verify --short HEAD)"

all: build

test: $(SRCS)
	go test ./...

deps:
	go mod tidy

build: deps $(BINARY_NAME)

$(BINARY_NAME): $(SRCS)
	go build -o $(BINARY_NAME) -ldflags $(LDFLAGS)

install:
	go install -ldflags $(LDFLAGS)

sys-install: build
	sudo install $(BINARY_NAME) /usr/local/bin

clean:
	rm -f $(BINARY_NAME)

.PHONY: all test deps build install clean
