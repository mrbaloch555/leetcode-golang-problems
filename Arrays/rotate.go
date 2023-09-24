package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reverse(nums[:n-k])
	reverse(nums[n-k:])
	reverse(nums)
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
