package main

import "fmt"

func main() {

	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	max := ""
	accumalator := ""

	for i := 0; i < len(s); i++ {
		accumalator += string(s[i])

		reverse := func(str string) string {
			res := ""
			for i := len(str) - 1; i > 0; i-- {
				res += string(str[i])
			}
			return res
		}(accumalator)

		if reverse != accumalator {
			if len(max) < len(reverse) {
				max = reverse
			}
			accumalator = ""
		}
	}

	fmt.Println("MAX => ", max)
	return s
}
