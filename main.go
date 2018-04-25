package main

import (
	"fmt"
	"log"
	"plugin"
)

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

func main() {
	fmt.Println("hello")
	plug, err := plugin.Open("./plugins/calculator/calc.so")
	if err != nil {
		log.Fatalln(err)
	}

	symCalc, err := plug.Lookup("Calc")
	if err != nil {
		log.Fatalln(err)
	}

	var calc Calculator
	calc, ok := symCalc.(Calculator)
	if !ok {
		log.Fatalln("Unexpected type from module symbol")
	}
	var exC struct{}
	calc.Init(exC)
	fmt.Printf("SUM: %d\n", calc.Sum(2, 4))
	fmt.Printf("SUM: %d\n", calc.Subtract(2, 4))
	fmt.Printf("SUM: %d\n", calc.SolveEquation(4, 2))
}
