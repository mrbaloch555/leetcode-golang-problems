// **********************************************************************************
// Problem Number: 100
// Problem Name: Same Tree
// Problem Link: https://leetcode.com/problems/same-tree/
// Status: Solved
// **********************************************************************************

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := TreeNode{}

	// a.Left = &b
	// a.Right = &c
	fmt.Println(isSameTree(&a, &a))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack_p := []*TreeNode{p}
	stack_q := []*TreeNode{q}

	for len(stack_p) > 0 && len(stack_q) > 0 {
		current_p := stack_p[len(stack_p)-1]
		current_q := stack_q[len(stack_q)-1]
		stack_p = stack_p[:len(stack_p)-1]
		stack_q = stack_q[:len(stack_q)-1]

		if current_p != nil && current_q != nil {
			if current_p.Val != current_q.Val {
				return false
			}

			if current_p.Left != nil && current_q.Left != nil {
				stack_p = append(stack_p, current_p.Left)
				stack_q = append(stack_q, current_q.Left)
			} else if current_p.Left == nil && current_q.Left != nil || current_p.Left != nil && current_q.Left == nil {
				return false
			}

			if current_p.Right != nil && current_q.Right != nil {
				stack_p = append(stack_p, current_p.Right)
				stack_q = append(stack_q, current_q.Right)
			} else if current_p.Right == nil && current_q.Right != nil || current_p.Right != nil && current_q.Right == nil {
				return false
			}
		}
	}
	return true
}
