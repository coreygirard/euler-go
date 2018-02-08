package main

import (
	"fmt"
	"strconv"
)

func main() {
	soFar := -1
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			prod := i * j
			prodStr := strconv.Itoa(prod)
			if prodStr == Reverse(prodStr) {
				if prod > soFar {
					soFar = prod
				}
			}
		}
	}
	fmt.Println(soFar)
}

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
