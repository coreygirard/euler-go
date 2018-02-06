package main

import "fmt"

func main() {
	a := 1
	b := 1
	c := a + b

	total := 0

	for {
		if a > 4000000 {
			break
		}
		if a%2 == 0 {
			total += a
		}

		a = b
		b = c
		c = a + b
	}
	fmt.Println(total)
}
