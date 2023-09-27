build:
	@go build -o bin/SeaOfKeysApi

run: build
	@HOST=localhost USER=mikkel PASSWORD=password PORT=5432 DATABASE=skole ./bin/SeaOfKeysApi
