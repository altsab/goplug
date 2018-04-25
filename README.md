# Golang plugins example

* Example made to understand how plugins work.

* To build a plugin run `go build -buildmode=plugin -o ./plugins/calculator/calc.so ./plugins/calculator/calculator.go`
* Main package loads "calculator" plugin and makes some calculations