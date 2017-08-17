VERSION = $(shell git describe --dirty --tags)
LDFLAGS = -ldflags "-X github.com/blp1526/gg.Version="$(VERSION)

.PHONY: all
all: build

.PHONY: test
test:
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
