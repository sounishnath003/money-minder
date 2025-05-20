.PHONY: install
install:
	go mod tidy
	go mod download
	go mod verify
	cd web && npm install

.PHONY: build-backend
build-backend:
	gofmt -w cmd/
	gofmt -w internal/
	GOARCH=arm64 go build -o ./tmp/moneyminder cmd/*.go

.PHONY: run-backend
run-backend: build-backend
	./tmp/moneyminder

.PHONY: frontend-run-dev
frontend-run-dev:
	cd web && npm run dev