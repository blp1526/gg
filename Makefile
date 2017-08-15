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
	go build -o tmp/gg cmd/gg/gg.go
	@echo

.PHONY: install
install: build
	mv tmp/gg ${GOPATH}/bin
	@echo

.PHONY: clean
clean:
	rm -rf tmp
	@echo
