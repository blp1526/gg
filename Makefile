VERSION = $(shell git describe --always --dirty --tags)
LDFLAGS = -ldflags "-X github.com/blp1526/gg.Version="$(VERSION)

.PHONY: all
all: build

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: vet
	go test -v --cover ./...
	@echo

.PHONY: tmp
tmp: test
	mkdir -p tmp
	@echo

.PHONY: build
build: tmp
	go build $(LDFLAGS) -o tmp/gg cmd/gg/gg.go
	@echo

.PHONY: install
install: build
	mv tmp/gg ${GOPATH}/bin
	@echo

.PHONY: clean
clean:
	rm -rf tmp
	@echo

.PHONY: tagging
tagging:
	git tag -a ${TAG} -m "${TAG} release"
	git push origin ${TAG}
	@echo

.PHONY: release
release:
	goreleaser --rm-dist
	@echo
