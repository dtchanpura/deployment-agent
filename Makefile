BINARY := cd-go

PACKAGES := $(shell go list ./... | grep -v /vendor)

DEPENDENCIES := \
    gopkg.in/src-d/go-git.v4


all: build silent-test

build: deps
	go build -o bin/$(BINARY) main.go

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)
