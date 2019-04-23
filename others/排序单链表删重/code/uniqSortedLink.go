package main

import (
	// "errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func NewNode(v int) *Node {
	return &Node{
		data: v,
		next: nil,
	}
}

func PrintList(n *Node) {
	fmt.Printf("print list:\n")
	for n != nil {
		fmt.Printf("%d,", n.data)
		n = n.next
	}
	fmt.Printf("\n")
}

func deleteDuplicate(n *Node) *Node {
	if n == nil || n.next == nil {
		return nil
	}

	start := NewNode(0)
	start.next = n

	curVal := n.data
	curNum := 1
	p1 := start
	p2 := n.next
	for p2 != nil {
		if p2.data == curVal {
			curNum += 1
		} else {
			if curNum == 1 {
				p1 = p1.next
			} else {
				p1.next = p2
			}
			curVal = p2.data
			curNum = 1
		}
		p2 = p2.next
	}
	if curNum > 1 {
		p1.next = p2
	}

	return start.next
}

func main() {
	n1 := NewNode(1)
	n11 := NewNode(1)
	n2 := NewNode(2)
	n22 := NewNode(2)
	n23 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)
	n44 := NewNode(4)

	n1.next = n11
	n11.next = n2
	n2.next = n22
	n22.next = n23
	n23.next = n3
	n3.next = n4
	n4.next = n44

	PrintList(n1)

	res := deleteDuplicate(n1)

	PrintList(res)

}
