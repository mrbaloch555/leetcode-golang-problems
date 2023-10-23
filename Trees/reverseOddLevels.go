package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	odd := 1
	for len(queue) > 0 {
		nextLevel := []*TreeNode{}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				nextLevel = append(nextLevel, current.Left)
			}
			if current.Right != nil {
				nextLevel = append(nextLevel, current.Right)
			}

		}
		if odd%2 == 1 {
			sort.Slice(nextLevel, func(i, j int) bool {
				return nextLevel[i].Val < nextLevel[j].Val
			})
		}
		odd++
		queue = append(queue, nextLevel...)
	}
	return root
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 3,
		Left: &TreeNode{Val: 8}, Right: &TreeNode{Val: 13},
	}, Right: &TreeNode{Val: 5, Left: &TreeNode{Val: 21}, Right: &TreeNode{Val: 34}}}

	fmt.Println(reverseOddLevels(root))
	fmt.Println(root.Right)
}
