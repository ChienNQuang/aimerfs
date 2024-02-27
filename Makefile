build:
	@go build -o ./bin/aimerfs

run: build
	@./bin/aimerfs
