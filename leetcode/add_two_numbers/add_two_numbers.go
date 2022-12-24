package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummyHead := &ListNode{}
	curr := dummyHead
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		newNode := &ListNode{sum % 10, nil}
		curr.Next = newNode
		curr = newNode

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return dummyHead.Next
}

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	n1 := buildSingleLinkedListFromArray([]int{3, 4, 2})
	n2 := buildSingleLinkedListFromArray([]int{4, 6, 5})
	printResult(addTwoNumbers(n1, n2))
}

func case2() {
	n1 := buildSingleLinkedListFromArray([]int{0})
	n2 := buildSingleLinkedListFromArray([]int{0})
	printResult(addTwoNumbers(n1, n2))
}

func case3() {
	n1 := buildSingleLinkedListFromArray([]int{9, 9, 9, 9, 9, 9, 9})
	n2 := buildSingleLinkedListFromArray([]int{9, 9, 9, 9})
	printResult(addTwoNumbers(n1, n2))
}

func buildSingleLinkedListFromArray(arr []int) *ListNode {
	var head *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		if head == nil {
			head = &ListNode{arr[i], nil}
		} else {
			temp := head
			head = &ListNode{arr[i], nil}
			head.Next = temp
		}
	}
	return head
}

func printResult(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	println()
}

func stringResult(node *ListNode) string {
	var str strings.Builder
	for node != nil {
		str.WriteString(strconv.Itoa(node.Val))
		node = node.Next
	}
	return str.String()
}
