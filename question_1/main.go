package main

import (
	"fmt"
	"math"
)

type Expression struct {
	n int
	x float32
}

func (e *Expression) input(n int, x float32) {
	e.n = n
	e.x = x
}

func (e *Expression) calc() float64 {
	sum := 0.0
	for i := 0; i <= e.n; i++ {
		sum = sum + math.Pow(float64(e.x), float64(i))
	}

	return sum
}

func main() {
	var n int
	var x float32

	fmt.Print("Enter n: ")
	fmt.Scan(&n)
	fmt.Print("Enter x: ")
	fmt.Scan(&x)

	exp := new(Expression)
	exp.input(n, x)
	result := exp.calc()
	fmt.Printf("Result: %g", result)
}
