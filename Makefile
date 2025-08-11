test:
	go mod tidy
	go test -v

install:
	go install

lint:
	golangci-lint run cmd/hexlet-path-size

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size
