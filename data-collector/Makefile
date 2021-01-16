.PHONY: build test clean

build:
	CGO_ENABLED=0 GO111MODULE=on go build -o app

run:
	./app

test:
	GO111MODULE=on go test ./...

clean:
	rm app