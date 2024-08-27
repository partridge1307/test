package main

import "fmt"

type DrawAnyShape struct{}

func (d *DrawAnyShape) drawShape(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			if k == 1 || k == 2*i-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	rhombus := new(DrawAnyShape)

	var n int
	fmt.Print("Input height of rhombus: ")
	fmt.Scan(&n)

	rhombus.drawShape(n)
}
