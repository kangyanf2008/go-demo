package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func recover(header *ListNode) *ListNode {
	var prev *ListNode
	cur := header
	for cur != nil {
		next := cur.Next
		cur.Next, prev = prev, cur
		prev, cur = cur, next
	}
	return prev
}

func recover2(node *ListNode) {
	var pre *ListNode
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

func main() {
	head := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	// 输出反转前的链表
	fmt.Print("Original List: ")
	p := head
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()

	// 反转链表
	head = recover(head)

	// 输出反转后的链表
	fmt.Print("Reversed List: ")
	p = head
	for p != nil {
		fmt.Printf("%d ", p.Val)
		p = p.Next
	}
	fmt.Println()
}
