run: build
	@./bin/filecrypt

build:
	@go build -o ./bin/filecrypt

test:
	@go test -v ./...