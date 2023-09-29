build:
	@go build -o bin/SeaOfKeysApi
run: build
	@DATABASETYPE="mysql" HOST=localhost USER=kronborg PASSWORD=password PORT=3306 DATABASE=seaofkeys ./bin/SeaOfKeysApi
	# @DATABASETYPE="postgres" HOST=localhost USER=mikkel PASSWORD=password PORT=5432 DATABASE=skole ./bin/SeaOfKeysApi

test:
	@go test -v ./...

