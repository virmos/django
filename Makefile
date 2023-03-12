## test: runs all tests
test:
	@go test -v ./...

## cover: opens coverage in browser
cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

## coverage: displays test coverage
coverage:
	@go test -cover ./...

## build_cli: builds the command line tool Django and copies it to myapp
build_cli:
	@go build -o ../myapp/Django ./cmd/cli

## build: builds the command line tool dist directory
build:
# @go build -o ./dist/Django ./cmd/cli
# windows users should delete the line above this one, and use the line below instead (uncommented)
	@go build -o dist/Django.exe ./cmd/cli
