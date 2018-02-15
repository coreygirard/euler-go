package main

import (
	"fmt"
	"math"
)

func main() {
	upTo := 2000000
	total := 0

	for a := 0; a < upTo; a++ {
		if IsPrime(a) {
			total += a
		}
	}
	fmt.Println(total)
}

func IsPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
