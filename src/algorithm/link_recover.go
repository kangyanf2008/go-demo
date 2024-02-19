package main

import "fmt"

func main() {
	header := &LinkNode{value: 100}
	next := header
	for i := 200; i < 1000; i = i + 100 {
		tem := &LinkNode{
			value: i,
		}
		next = next.add(tem)
	}
	next = header
	for next != nil {
		fmt.Print(next.value, "\t")
		next = next.next
	}
	fmt.Print("\n")
	next = header.RLinkNode()
	for next != nil {
		fmt.Print(next.value, "\t")
		next = next.next
	}
}

type LinkNode struct {
	value int
	next  *LinkNode
}

func (header *LinkNode) add(n *LinkNode) *LinkNode {
	header.next = n
	return n
}

func (header *LinkNode) RLinkNode() *LinkNode {
	cur := header
	var pre *LinkNode
	for cur != nil {
		next := cur.next
		cur.next, pre = pre, cur
		cur = next
	}
	return pre
}
