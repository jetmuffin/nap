.PHONY: all build test clean

VERSION=$(shell git describe --always --tags --abbre=0)
BUILD_DATE=$(shell date -u +%Y-%m-%d)
gitCommit=$(shell git describe --tags --long)
PKG="github.com/jetmuffin/nap/pkg"

GO_LDFLAGS=-X $(PKG)/version.gitCommit=$(GIT_COMMIT) -X $(PKG)/version.buildDate=$(BUILD_DATE) -w -s

UNAME=$(shell uname -s)
CGO_ENABLED=0
ifeq ($(UNAME),Darwin)
CGO_ENABLED=1
endif

default: build

build: clean
	CGO_ENABLED=${CGO_ENABLED} go build -v -a -ldflags "${GO_LDFLAGS}" -o bin/nap main.go

docker:
	docker build --tag nap:$(shell git rev-parse --short HEAD) .

clean:
	rm -rfv ./bin/*

fmt:
	@echo $@
	@test -z "$$(gofmt -s -l . 2>/dev/null | grep -Fv 'vendor/' | grep -v ".pb.go$$"| tee /dev/stderr)" || \
		(echo "please format Go code with 'gofmt -s -w .'" && false)

lint:
	@echo $@
	@test -z "$$(golint ./... | grep -Fv 'vendor/' | grep -v ".pb.go:" | tee /dev/stderr)"

test:
	@echo $@
	@go test -v -cover ./...
