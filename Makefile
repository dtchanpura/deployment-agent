PKGS := $(shell go list ./... | grep -v /vendor)
GO_FILES := $(shell find . -iname '*.go' -type f | grep -v /vendor/) # All the .go files, excluding vendor/
BUILD_DATE := $(shell date --iso-8601=seconds)
VERSION ?= latest

LDFLAGS := -s -w -X github.com/dtchanpura/deployment-agent/constants.Version=$(VERSION) -X github.com/dtchanpura/deployment-agent/constants.BuildDateStr=$(BUILD_DATE)
BINARY := deployment-agent
PLATFORMS := windows linux darwin
os = $(word 1, $@)

bootstrap:
	go get -u golang.org/x/lint/golint           # Linter
	go get -u honnef.co/go/tools/cmd/staticcheck # Badass static analyzer/linter
	# go get honnef.co/go/tools/cmd/megacheck    # Badass static analyzer/linter
	go get -u github.com/fzipp/gocyclo           # Cyclomatic complexity check
	go mod download

test:
	go test -v -race $(PKGS)        # Normal Test
	go vet ./...                    # go vet is the official Go static analyzer
	staticcheck ./...               # "go vet on steroids" + linter
	gocyclo -over 19 $(GO_FILES)    # forbid code with huge functions
	golint -set_exit_status $(PKGS) # one last linter


$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64 -ldflags="$(LDFLAGS)"
	tar -czvf release/$(BINARY)-$(VERSION)-$(os)-amd64.tar.gz README.md -C release/ $(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: windows linux darwin

clean:
	rm -rf release/*
