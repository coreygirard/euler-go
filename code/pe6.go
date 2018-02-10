package main

import "fmt"

func main() {
	number := 100

	sumOfSquares := 0
	squareOfSums := 0
	for n := 1; n <= number; n++ {
		sumOfSquares += n * n
		squareOfSums += n
	}
	squareOfSums *= squareOfSums

	fmt.Println(squareOfSums - sumOfSquares)
}
