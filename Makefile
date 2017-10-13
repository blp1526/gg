.PHONY: all
all: snapshot

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test: vet
	go test -v --cover ./...

.PHONY: snapshot
snapshot: test
	goreleaser --rm-dist --snapshot

.PHONY: clean
clean:
	rm -rf dist

.PHONY: tagging
tagging:
	git tag -a ${TAG} -m "${TAG} release"

.PHONY: release
release:
	goreleaser --rm-dist
