.PHONY: build install clean

build:
	@mkdir -p tmp
	@go build -o tmp/gg cmd/gg/gg.go

install:
	@mv tmp/gg ${GOPATH}/bin

clean:
	@rm -rf tmp
