build:
	@go build -o bin/SeaOfKeysApi

run: build
	@./bin/SeaOfKeysApi
