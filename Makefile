BINARY_NAME=ki18n
ARCH="`uname -s`"

all: build
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -ldflags "-s -w" -o bin/Darwin-amd64/$(BINARY_NAME) -v

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/Linux-amd64/$(BINARY_NAME) -v

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/windows/$(BINARY_NAME) -v

build:
	go build  -ldflags "-s -w" -o bin/$(ARCH)-amd64/$(BINARY_NAME) -v
