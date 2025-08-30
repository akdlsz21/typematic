BINARY := typematic
PKG := ./cmd/typematic
DIST := dist

.PHONY: build test vet fmt clean dist build-all package checksums all

build:
	go build -o $(BINARY) $(PKG)

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -s -w .

clean:
	rm -rf $(DIST) $(BINARY) $(BINARY).exe

dist:
	mkdir -p $(DIST)

# Build statically (pure Go) for three OS targets
build-all: dist
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o $(DIST)/$(BINARY)-linux-amd64   $(PKG)
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o $(DIST)/$(BINARY)-darwin-arm64  $(PKG)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(DIST)/$(BINARY)-windows-amd64.exe $(PKG)

# Package artifacts into archives suitable for release
package: build-all
	tar -C $(DIST) -czf $(DIST)/$(BINARY)-linux-amd64.tar.gz $(BINARY)-linux-amd64
	tar -C $(DIST) -czf $(DIST)/$(BINARY)-darwin-arm64.tar.gz $(BINARY)-darwin-arm64
	zip -j $(DIST)/$(BINARY)-windows-amd64.zip $(DIST)/$(BINARY)-windows-amd64.exe

# Generate SHA256 checksums (Ubuntu runner has sha256sum)
checksums: package
	(cd $(DIST) && sha256sum *.tar.gz *.zip > $(BINARY)-SHA256SUMS.txt)

all: clean fmt vet test build-all package checksums

