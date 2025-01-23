build:
	@go build -o bin/inv cmd/*.go

run: build
	@./bin/inv