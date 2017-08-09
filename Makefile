.PHONY: build install clean

build:
	@mkdir -p tmp
	@go build -o tmp/gg

install:
	@mv tmp/gg ${GOPATH}/bin

clean:
	@rm -rf tmp
