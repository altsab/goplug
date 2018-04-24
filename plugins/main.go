package main

// Additional variable for more complex computations
var Additional EnhancedCalc

// Calculator implements simple calculations
type Calculator interface {
	Sum(a, b int) int
	Subtract(a, b int) int
	Init(enhanced EnhancedCalc)
}

// EnhancedCalc implements more complex calculations
type EnhancedCalc interface {
	Multiply(a, b int) int
	Divide(a, b int) int
}

// Initialize complex calc
func Init(someVar EnhancedCalc) {
	Additional = someVar
}

func Sum(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}
