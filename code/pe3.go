package main

import "fmt"

func main() {
	n := 600851475143
	for i := n; i > 1; i-- {
		if n%i == 0 {
			if IsPrime(i) {
				fmt.Println(i)
				break
			}
		}
	}
}

func IsPrime(n int) bool {
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
