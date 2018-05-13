PKGS := $(shell go list ./... | grep -v /vendor)
BINARY := deployment-agent
PLATFORMS := windows linux darwin
os = $(word 1, $@)

.PHONY:	test
test:
	go test $(PKGS)

.PHONY:	lint
lint:
	gometalinter ./... --vendor

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	mkdir -p release
	GOOS=$(os) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(os)-amd64
	@echo "Creating archive..."
	@tar -czvf release/$(BINARY)-$(VERSION)-$(os)-amd64.tar.gz README.md -C release/ $(BINARY)-$(VERSION)-$(os)-amd64

.PHONY: release
release: windows linux darwin

clean:
	rm -r release/*
