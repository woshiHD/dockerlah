package main

import (
	"fmt"
)

func main() {
	var l1 = []int{1, 2, 4}
	var l2 = []int{1, 3, 4}

	fmt.Print(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newNode := new(ListNode)
	newNode := append(l1, l2...)
	return newNode
}

// func mergeTwoLists(x []int, y []int) []int {
// 	s := append(x, y...)
// 	sort.Ints(s)
// 	return s
// }
