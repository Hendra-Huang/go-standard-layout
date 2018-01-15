# go-standard-layout

Reading some articles for designing go standard layout. These are the guidelines that I use:
1. Root package is for domain types
2. Group subpackages by dependency
3. Make dependencies explicit!
4. Main package ties together dependencies
5. Loggers are dependencies!
6. Use the “underscore test” package
7. Use a shared mock subpackage
8. Using test cases style for testing
9. Provide database integration test
10. `testdata` folder name for containing test fixtures
11. Define interface on the implementation side
12. Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values
13. The implementing package should return concrete (usually pointer or struct) types

References:
1. https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
2. https://peter.bourgon.org/go-best-practices-2016
3. https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c
4. https://www.youtube.com/watch?v=yszygk1cpEc
5. https://dave.cheney.net/2016/05/10/test-fixtures-in-go
6. https://github.com/golang/go/wiki/CodeReviewComments#interfaces
