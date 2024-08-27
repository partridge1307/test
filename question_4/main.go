package main

import "fmt"

type ArrIntManager struct {
	value []int
}

func (a *ArrIntManager) sum() int {
	sum := 0

	for num := range a.value {
		sum = sum + num
	}

	return sum
}

func isPrimeNum(num int) bool {
	i := 1
	count := 0
	for i <= num {
		if num%i == 0 {
			count++
		}
	}

	if count == 2 {
		return true
	} else {
		return false
	}
}

func (a *ArrIntManager) sumOfPrimeNums() int {
	sum := 0

	for num := range a.value {
		if isPrimeNum(num) {
			sum = sum + num
		}
	}

	return sum
}

func (a *ArrIntManager) printSeqValue() {
	for i := 0; i < len(a.value)-2; i++ {
		if a.value[i]+a.value[i+1] == a.value[i+2] {
			fmt.Printf("%d, %d, %d\n", a.value[i], a.value[i+1], a.value[i+2])
		}
	}
}

func main() {
	aM := ArrIntManager{
		value: []int{1, 2, 3, 5, 4, 1, 3, 4, 5, 4, 5, 9, 7, 0, 11, 13, 10, 23},
	}

	a := aM.sum()
	fmt.Printf("Sum of elements in array: %d\n", a)

	b := aM.sumOfPrimeNums()
	fmt.Printf("Sum of prime number: %d\n", b)

	aM.printSeqValue()
}
