include ./env_${MYAPPENV}

build:
	GOOS=linux GOARCH=amd64 go build ./cmd/myappserver/main.go

unit-test:
	go test -v ./...

test:
	TEST_DB_HOST=${TEST_DB_HOST} TEST_DB_PORT=${TEST_DB_PORT} go test -v -tags=integration ./...

.PHONY: build unit-test test
