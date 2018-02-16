package main

import "fmt"

func numFactors(n int) int {
	factors := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors++
		}
	}

	return factors
}

func main() {
	n := 1
	t := 2
	for numFactors(n) <= 500 {
		n += t
		t++
	}
	fmt.Println(n)
}
