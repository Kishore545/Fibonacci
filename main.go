package main

import (
	"fmt"

	"github.com/Kishore545/pop"
)

func FibonacciArray(x []int) []int {

	fibSeries := []int{x[0], x[1]}

	for i := 2; i < len(x); i++ {
		next := fibSeries[i-1] + fibSeries[i-2]
		fibSeries = append(fibSeries, next)
	}
	return fibSeries
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fibseries := FibonacciArray(x)

	fmt.Println("Generated Fibonacci Series:", fibseries)
	y := pop.Bark()
	z := pop.Barks()

	fmt.Println(y)
	fmt.Println(z)
}
