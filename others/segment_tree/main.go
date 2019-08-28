package main

import (
	"fmt"
)

type Node struct {
	LeftIdx   int64
	RightIdx  int64
	Value     int64
	LeftNode  *Node
	RightNode *Node
}

var keyFunc func(a, b int64) int64

func buildTree(leftIdx, rightIdx int64, arr []int64) *Node {
	if leftIdx == rightIdx {
		return &Node{
			LeftIdx:  leftIdx,
			RightIdx: rightIdx,
			Value:    arr[leftIdx],
		}
	}

	mid := (leftIdx + rightIdx) / 2
	leftNode := buildTree(leftIdx, mid, arr)
	rightNode := buildTree(mid+1, rightIdx, arr)
	return &Node{
		LeftIdx:   leftIdx,
		RightIdx:  rightIdx,
		Value:     keyFunc(leftNode.Value, rightNode.Value),
		LeftNode:  leftNode,
		RightNode: rightNode,
	}
}

func updateValue(root *Node, idx, value int64) {
	if root == nil {
		return
	}
	if root.LeftIdx == root.RightIdx && root.LeftIdx == idx {
		root.Value = value
		return
	}
	mid := (root.LeftIdx + root.RightIdx) / 2
	if idx <= mid {
		updateValue(root.LeftNode, idx, value)
	} else {
		updateValue(root.RightNode, idx, value)
	}
	root.Value = keyFunc(root.LeftNode.Value, root.RightNode.Value)
	return
}

func query(root *Node, leftIdx, rightIdx int64) int64 {
	if root.LeftIdx == leftIdx && root.RightIdx == rightIdx {
		return root.Value
	}
	mid := (root.LeftIdx + root.RightIdx) / 2
	if rightIdx <= mid {
		return query(root.LeftNode, leftIdx, rightIdx)
	} else if leftIdx > mid {
		return query(root.RightNode, leftIdx, rightIdx)
	}
	leftValue := query(root, leftIdx, mid)
	rightValue := query(root, mid+1, rightIdx)
	return keyFunc(leftValue, rightValue)
}

func logFunc(left, right, value int64) {
	fmt.Printf("left:%d, right:%d, value:%d\n", left, right, value)
}

func main() {
	// keyFunc = func(a, b int64) int64 { return a + b }
	keyFunc = func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	arr := []int64{3, 1, 7, 0, 5, 9, 2, 4, 6, 8}
	fmt.Printf("arr: %#v\n", arr)
	root := buildTree(0, int64(len(arr)-1), arr)
	var queryLeftIdx int64
	var queryRightIdx int64
	queryLeftIdx = 0
	queryRightIdx = 9
	logFunc(queryLeftIdx, queryRightIdx, query(root, queryLeftIdx, queryRightIdx))
	queryLeftIdx = 1
	queryRightIdx = 4
	logFunc(queryLeftIdx, queryRightIdx, query(root, queryLeftIdx, queryRightIdx))
	queryLeftIdx = 5
	queryRightIdx = 8
	logFunc(queryLeftIdx, queryRightIdx, query(root, queryLeftIdx, queryRightIdx))
}
