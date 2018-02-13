package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 0; a < 1000; a++ {
		for b := a + 1; b < 1000; b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if c >= float64(b+1) {
				if c == math.Trunc(c) {
					c := int(c)
					if a+b+c == 1000 {
						fmt.Println(a * b * c)
						break
					}
				}
			}
		}
	}
}
