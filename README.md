# go-utils

A library of quite basic functionality that is missing from Go's std lib

## Packages

* [math](https://github.com/michaelcampbellgithub/go-utils/tree/main/math): Extension of the std lib [math](https://cs.opensource.google/go/go/+/master:src/math/) package

## ToDo

This library remains under active development, and so there are plenty of issues unresolved (and I'm too lazy to create actual issues on Github):
* math package - Add badge in [math/README.md](https://github.com/michaelcampbellgithub/go-utils/blob/main/math/README.md) to indicate successful tests
* math package - Look at missing benchmark tests (commented out). Need to substitute the lack of support for these data types in the std lib [math/rand](https://cs.opensource.google/go/go/+/master:src/math/rand/) package
* math package - What to do with benchmark output? Seems pretty pointless not having a quality gate set
* Think up other functionality for this library beyond Min & Max functions
