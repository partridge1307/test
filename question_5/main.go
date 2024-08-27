package main

import "fmt"

type ManagerInt struct{}

func (mI *ManagerInt) sumK(m []int, n int) {
	arr := make([][]int, len(m))

	for i := range arr {
		arr[i] = make([]int, 0)
	}

	i := 0
	for _, num := range m {
		s := 0
		if len(arr[i]) > 0 {
			for _, j := range arr[i] {
				s = s + j
			}
		}

		if s < n {
			arr[i] = append(arr[i], num)
		} else {
			i++
		}
	}

	for i := range arr {
		for j := range arr[i] {
			fmt.Printf("%d, ", arr[i][j])
		}
		fmt.Println()
	}
}

func main() {
	mI := new(ManagerInt)
	m := []int{1, 4, 6, 3, 2, 2, 8}
	n := 10

	mI.sumK(m, n)
}
