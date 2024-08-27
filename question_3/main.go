package main

import (
	"fmt"
)

type MyString struct {
	value string
}

func (s *MyString) sumOfString() int {
	sum := 0
	for _, char := range s.value {
		sum = sum + int(char)
	}

	return sum
}

func (s *MyString) sumOfCapitalCharacters() int {
	sum := 0
	for _, char := range s.value {
		if char-'0' >= 65 && char-'0' <= 90 {
			sum = sum + int(char)
		}
	}

	return sum
}

func main() {
	s := MyString{value: "Hello programmers. Im Developers"}

	a := s.sumOfString()
	fmt.Printf("Sum of string in ascii: %d\n", a)

	b := s.sumOfCapitalCharacters()
	fmt.Printf("Sum of capital characters in string: %d\n", b)

	var n int
	fmt.Print("Iput n: ")
	fmt.Scan(&n)
}
