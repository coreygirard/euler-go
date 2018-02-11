package main

import "fmt"

func main() {
	target := 10001

	m := make(map[int]int)
	ptr := make(map[int]int)
	ptr[0] = 2

	for n := 2; n < 110000; n++ {
		m[n] = 1
		ptr[n-1] = n
	}

	numFound := 0
	var factor int
	for numFound < target {
		factor = ptr[0]
		numFound++

		trailPtr := 0
		for i := ptr[0]; i != 0; {
			if i%factor == 0 {
				m[i] = 0
				ptr[trailPtr] = ptr[i]
			}
			trailPtr = i
			i = ptr[i]
		}
	}

	fmt.Println(factor)
}
