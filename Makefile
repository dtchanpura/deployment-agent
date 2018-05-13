PKGS := $(shell go list ./... | grep -v /vendor)

BINARY := deployment-agent
PLATFORMS := windows linux darwin
VERSION ?= latest
os = $(word 1, $@)

.PHONY:	test
test: config
	go test $(PKGS)

.PHONY: config
config:
	go get github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64
	@echo "Creating archive..."
	@tar -czvf release/$(BINARY)-$(VERSION)-$(os)-amd64.tar.gz README.md -C release/ $(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: windows linux darwin

.PHONY: clean
clean:
	rm -r release/*
