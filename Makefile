.PHONY: install
install:
	go mod tidy
	go mod download
	go mod verify

build:
	gofmt -w cmd/
	gofmt -w internal/
	GOARCH=arm64 go build -o ./tmp/moneyminder cmd/*.go

run: build
	./tmp/moneyminder