package main

import "fmt"

func integerBreak(n int) int {

	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	product := 1
	for n > 4 {
		product *= 3
		n -= 3
	}
	product *= n
	return product
}

func main() {
	n := 2
	fmt.Println(integerBreak(n))
}
