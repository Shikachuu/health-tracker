export CGO_ENABLED=0

all: test build

build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/health
	GOOS=windows GOARCH=amd64 go build -o ./bin/health.exe

test:
	go test -v ./pkg/* -cover