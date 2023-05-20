// **********************************************************************************
// Problem Number: 80
// Problem Name: Remove Duplicates from Sorted Array II
// Problem Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i < len(nums)-2 && nums[i] == nums[i+1] && nums[i+1] == nums[i+2] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
