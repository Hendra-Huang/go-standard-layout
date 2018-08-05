appenv := $(shell cat ./env_${MYAPPENV})

run-myappserver:
	${appenv} go run ./cmd/myappserver/main.go

build:
	CGO_ENABLED=0 go build -o bin/myappserver ./cmd/myappserver/main.go

unit-test:
	${appenv} go test -v ./...

partial-test:
	${appenv} go test -v -tags=integration ${ARGS}

test:
	${appenv} go test -v -tags=integration ./...

.PHONY: run-myappserver build unit-test partial-test test
