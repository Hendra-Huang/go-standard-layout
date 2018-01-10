# go-standard-layout

Reading some articles for designing go standard layout. These are the points that I follow:
1. Root package is for domain types
2. Group subpackages by dependency
3. Make dependencies explicit!
4. Main package ties together dependencies
5. Loggers are dependencies!

References:
1. https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
2. https://peter.bourgon.org/go-best-practices-2016
