package main

import "fmt"

func main() {
	// 40! / 20! / (40-20)!

	numer := 1
	denom := 1
	for i := 1; i <= 20; i++ {
		denom *= i
	}

	for i := 21; i <= 40; i++ {
		numer *= i
		for j := 2; j <= 5; j++ {
			for (numer%j == 0) && (denom%j == 0) {
				numer /= j
				denom /= j
			}
		}
	}

	fmt.Println(numer / denom)
}
