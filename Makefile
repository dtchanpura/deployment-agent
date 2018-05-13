PKGS := $(shell go list ./... | grep -v /vendor)
GO_FILES := $(shell find . -iname '*.go' -type f | grep -v /vendor/) # All the .go files, excluding vendor/

BINARY := deployment-agent
PLATFORMS := windows linux darwin
VERSION ?= latest
os = $(word 1, $@)

bootstrap:
	go get github.com/golang/dep/cmd/dep
	go get github.com/golang/lint/golint                        # Linter
	go get honnef.co/go/tools/cmd/megacheck                     # Badass static analyzer/linter
	go get github.com/fzipp/gocyclo
	dep ensure

test:
	go test -v -race $(PKGS)
	go vet ./...                    # go vet is the official Go static analyzer
	megacheck ./...                 # "go vet on steroids" + linter
	gocyclo -over 19 $(GO_FILES)      # forbid code with huge functions
	golint -set_exit_status $(PKGS) # one last linter


$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64
	tar -czvf release/$(BINARY)-$(VERSION)-$(os)-amd64.tar.gz README.md -C release/ $(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: windows linux darwin

clean:
	rm -rf release/*
