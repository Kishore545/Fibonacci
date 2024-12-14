package main

import (
	"fmt"
)

func FibonacciArray(x []int) []int {

	if len(x) < 2 {
		return []int{}
	}

	fibseries := []int{x[0], x[1]}

	for i := 2; i < len(x); i++ {
		next := fibseries[i-1] + fibseries[i-2]
		fibseries = append(fibseries, next)
	}
	return fibseries
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fibseries := FibonacciArray(x)

	fmt.Println("Generated Fibonacci Series:", fibseries)
}
