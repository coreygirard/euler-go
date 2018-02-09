package main

import "fmt"

func main() {
	number := 20

	m := make(map[int]int)

	for n := 2; n < number; n++ {
		temp := n
		i := 2

		q := 0
		for temp > 1 {
			if temp%i == 0 {
				q++
				temp /= i

				value, ex := m[i]
				if ex == true {
					m[i] = q
				} else if q > value {
					m[i] = q
				}
			} else {
				i++
			}
		}
	}

	result := 1
	for key, value := range m {
		for i := 0; i < value; i++ {
			result *= key
		}
	}
	fmt.Println(result)
}

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}
