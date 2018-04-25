package main

type calc struct{}
type excalc struct{}

// Calculator implements basic arithmetic calculations
type Calculator interface {
	Init(enhanced struct{})
	Sum(a, b int) int
	Subtract(a, b int) int
	SolveEquation(a, b int) int
}

// EnhancedCalc implements more complex calculations
type EnhancedCalc interface {
	Multiply(a, b int) int
	Divide(a, b int) int
}

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
	return exCalc.Multiply(a, a) + exCalc.Divide(a, b) + exCalc.Multiply(4, a)
}

var Calc calc
var exCalc excalc
