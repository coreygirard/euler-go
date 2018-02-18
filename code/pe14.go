package main

import "fmt"

func collatzLength(n int) int {
	i := 1
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		i++
	}

	return i
}

func main() {
	loc := -1
	leng := -1

	temp := 0
	for n := 1; n < 1000000; n++ {
		temp = collatzLength(n)
		if temp > leng {
			loc = n
			leng = temp
		}
	}
	fmt.Println(loc)
}
