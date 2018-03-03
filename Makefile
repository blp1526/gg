VERSION = $(shell git describe --always --dirty --tags)
LDFLAGS = -ldflags "-X github.com/blp1526/gg.Version="$(VERSION)

.PHONY: all
all: build

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: vet
	go test ./... -v --cover -race -covermode=atomic -coverprofile=coverage.txt

.PHONY: build
build: test
	go build $(LDFLAGS) -o tmp/gg cmd/gg/gg.go
	@echo

.PHONY: clean
clean:
	rm -rf dist

.PHONY: tagging
tagging:
	git tag -a ${TAG} -m "${TAG} release"

.PHONY: release
release:
	goreleaser --rm-dist
