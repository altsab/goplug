package main

type calc struct{}
type excalc struct{}

// Initialize complex calc
func (c calc) Init(someVar struct{}) {
	exCalc = someVar
}

func (c calc) Sum(a, b int) int {
	return a + b
}

func (c calc) Subtract(a, b int) int {
	return a - b
}

func (ex excalc) Multiply(a, b int) int {
	return a * b
}

func (ex excalc) Divide(a, b int) int {
	return a / b
}

func (c calc) SolveEquation(a, b int) int {
	return exCalc.Multiply(a, a) + exCalc.Divide(a, b) + exCalc.Multiply(2, a)
}

var Calc calc
var exCalc excalc
