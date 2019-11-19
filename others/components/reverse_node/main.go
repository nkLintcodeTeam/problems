package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

func makeList(s []int) *Node {
	if len(s) == 0 {
		return nil
	}
	head := &Node{}
	cur := head
	for _, v := range s {
		cur.Next = &Node{Val: v}
		cur = cur.Next
	}
	return head.Next
}

func printList(l *Node) {
	fmt.Printf("PrintList: ")
	cur := l
	for cur != nil {
		fmt.Printf("%+v,", cur.Val)
		cur = cur.Next
	}
	fmt.Printf("\n")
}

func reverseLink(l *Node) *Node {
	if l == nil {
		return nil
	}
	var newHead *Node
	oldCur := l
	for oldCur != nil {
		oldNext := oldCur.Next

		oldCur.Next = newHead
		newHead = oldCur

		oldCur = oldNext
	}
	return newHead
}

func tryCase(arr []int) {
	l := makeList(arr)
	printList(l)
	printList(reverseLink(l))
}

func main() {
	tryCase([]int{1, 2, 3, 4})
	tryCase([]int{1})
	tryCase([]int{})
}
