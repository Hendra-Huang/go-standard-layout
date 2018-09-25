[![Go Report Card](https://goreportcard.com/badge/github.com/Hendra-Huang/go-standard-layout?style=flat-square)](https://goreportcard.com/report/github.com/Hendra-Huang/go-standard-layout)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/Hendra-Huang/go-standard-layout)
[![Release](https://img.shields.io/github/release/Hendra-Huang/go-standard-layout.svg?style=flat-square)](https://github.com/Hendra-Huang/go-standard-layout/releases/latest)

# go-standard-layout

Reading some articles for designing go standard layout. These are the guidelines that I use:
1. Root package is for domain types
2. Group subpackages by dependency
3. Make dependencies explicit!
4. Main package ties together dependencies
5. Loggers are dependencies!
6. Use the “underscore test” package
7. Use a shared mock subpackage
8. Using table-driven-test style for testing
9. Provide database integration test
10. `testdata` folder name for containing test fixtures
11. Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values
12. The implementing package should return concrete (usually pointer or struct) types

## About the project

This is a small app that consist of 2 entities, User and Article. 1 user can have more than 1 articles. For both entities, there are some sample endpoints provided. I wrap the router and provides basic monitoring using prometheus. I use [dep](https://github.com/golang/dep) for depedency management. There is simple opentracing implementation in this app. This app is also provided with unit test and parallel database integration test.

## About the structure

The files in the root package contains domain logic (business logic) of the app. `cmd` contains entrypoint of the app / command. `script` contains shell script that can help to automate the build process, etc. Currently, `script` contains script for running test with code coverage preview. The rest of folders are the dependencies of the app. The data layer is at `mysql` package. You can find list of endpoints in `router` package. HTTP handler is located in `handler` package inside `server` folder.

## References
1. https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
2. https://peter.bourgon.org/go-best-practices-2016
3. https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c
4. https://www.youtube.com/watch?v=yszygk1cpEc
5. https://dave.cheney.net/2016/05/10/test-fixtures-in-go
6. https://github.com/golang/go/wiki/CodeReviewComments#interfaces
