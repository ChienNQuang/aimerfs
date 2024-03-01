build:
	@go build -o ./bin/aimerfs

run: build
	@./bin/aimerfs

test:
	@go test -v ./...