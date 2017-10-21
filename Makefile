BINARY := cd-go

PACKAGES := $(shell go list ./... | grep -v /vendor)

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

DEPENDENCIES := \
    gopkg.in/src-d/go-git.v4 \
    github.com/gin-gonic/gin


all: build silent-test

build: deps
	go build -o bin/$(BINARY)-$(GOOS)-$(GOARCH) main.go

test:
	go test -v $(PACKAGES)

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

deps:
	go get $(DEPENDENCIES)
