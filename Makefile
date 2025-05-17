.PHONY: install
install:
	go mod tidy
	go mod download
	go mod verify

.PHONY: build
build:
	gofmt -w cmd/
	gofmt -w internal/
	GOARCH=arm64 go build -o ./tmp/moneyminder cmd/*.go

.PHONY: run
run: build
	./tmp/moneyminder