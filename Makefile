PKG = github.com/k1LoW/repin
COMMIT = $$(git describe --tags --always)
OSNAME=${shell uname -s}
ifeq ($(OSNAME),Darwin)
	DATE = $$(gdate --utc '+%Y-%m-%d_%H:%M:%S')
else
	DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
endif

export GO111MODULE=on
export CGO_ENABLED=0

BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)

default: test

ci: depsdev test

test:
	go test ./... -coverprofile=coverage.out -covermode=count

fuzz:
	go test -fuzz ./... -fuzztime=60s

lint:
	golangci-lint run ./...

build:
	go build -ldflags="$(BUILD_LDFLAGS)" -o repin cmd/repin/main.go

depsdev:
	go install github.com/Songmu/gocredits/cmd/gocredits@latest

prerelease_for_tagpr:
	go mod tidy
	gocredits -w .
	git add CHANGELOG.md CREDITS go.mod go.sum

.PHONY: default test
