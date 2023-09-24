package main

import (
	"fmt"
	"math/big"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	fmt.Println(addTwoNumbers(&l1, &l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1_val := new(big.Int)
	l1_val, _ = l1_val.SetString(travers(l1, ""), 10)
	l2_val := new(big.Int)
	l2_val, _ = l1_val.SetString(travers(l2, ""), 10)
	total := l1_val.Add(l1_val, l2_val)
	fmt.Println(total)
	// list := strconv.Itoa(total)
	// if len(list) > 0 {
	// 	fir_val, _ := strconv.ParseInt(list[len(list)-1], 10, 64)
	// 	head := &ListNode{
	// 		Val: int(fir_val),
	// 	}
	// 	current := head
	// 	for i := len(list) - 2; i > 0; i-- {
	// 		val, _ := strconv.ParseInt(list[i], 10, 64)
	// 		node := &ListNode{
	// 			Val: int(val),
	// 		}
	// 		current.Next = node
	// 		current = current.Next
	// 	}
	// 	return head
	// } else {
	// 	return &ListNode{}
	// }
	return &ListNode{}
}

func travers(node *ListNode, s string) string {
	if node == nil {
		return s
	} else {
		s += travers(node.Next, s) + strconv.Itoa(node.Val)
	}
	return s
}
