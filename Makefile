GO=go
MOCK=mockery
BINARY=govdev
MAIN=cmd/govdev/main.go

VERSION=`git describe --abbrev=0 --tags`
GIT_HASH=`git rev-parse HEAD`
BUILD_TIME=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.GitHash=${GIT_HASH} -X main.BuildTime=${BUILD_TIME}"

.PHONY: clean test docs

default: install

build: test
	$(GO) build $(LDFLAGS) -o $(BINARY) $(MAIN)

build-darwin-amd64: test
	GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o bin/$(BINARY)-darwin $(MAIN)

build-linux-amd64: test
	GOOS=linux GOARCH=amd64 $(GO) build $(LDFLAGS) -o bin/$(BINARY)-linux $(MAIN)

build-windows-amd64: test
	GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o bin/$(BINARY)-windows $(MAIN)

all: clean test build-darwin-amd64 build-linux-amd64 build-windows-amd64

test:
	$(GO) test -v ./...

run: test
	$(GO) run $(MAIN)

install: test
	$(GO) install $(LDFLAGS) ./...

docker:
	docker build -t $(BINARY):$(VERSION) .

docs:
	godoc -http=:6061

fmt:
	$(GO) fmt ./...

vet: 
	$(GO) vet -v ./...

coverage:
	$(GO) test -cover -coverprofile=c.out ./...
	$(GO) tool cover -html=c.out -o coverage.html

clean:
	$(GO) clean
	rm -rf bin/$(BINARY)*
	rm -f $(BINARY)

# Project Specific
