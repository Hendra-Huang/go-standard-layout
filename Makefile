include ./env_${MYAPPENV}

run-myappserver:
	DB_HOST=${DB_HOST} DB_PORT=${DB_PORT} go run ./cmd/myappserver/main.go

build:
	GOOS=linux GOARCH=amd64 go build ./cmd/myappserver/main.go

unit-test:
	go test -v ./...

partial-test:
	TEST_DB_HOST=${TEST_DB_HOST} TEST_DB_PORT=${TEST_DB_PORT} go test -v -tags=integration ${ARGS}

test:
	TEST_DB_HOST=${TEST_DB_HOST} TEST_DB_PORT=${TEST_DB_PORT} go test -v -tags=integration ./...

.PHONY: build unit-test partial-test test
