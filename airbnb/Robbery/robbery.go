package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMaxInt(a int, s ...int) int {
	res := a
	for _, v := range s {
		if v > res {
			res = v
		}
	}
	return res
}

// rob1D 给定一维数组，获取抢劫最大值
// rst[i]表示数组s的前i个数的最大值，那么对于rst[i]有两种选择: 抢劫i/不抢劫i
// rst[i]=max(rst[i-1](不抢劫i+1), rst[i-2]+s[i](抢劫i+1))
func rob1D(s []int) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return s[0]
	}
	rst := make([]int, len(s), len(s))
	rst[0] = 0
	rst[1] = getMaxInt(s[0], s[1])
	for i := 2; i < len(s); i = i + 1 {
		rst[i] = getMaxInt(rst[i-1], rst[i-2]+s[i])
	}
	return rst[len(rst)-1]
}

// robCircle s是一个环
func robCircle(s []int) int {
	// 环形的，可以分为两类: 1.包含s[0]不包含s[-1]，
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return s[0]
	}

	return getMaxInt(rob1D(s[1:]), rob1D(s[:len(s)-1]))
}

// robTree 对一颗二叉树进行抢劫
// 对于每一个节点A，左右孩子节点B、C，有两种选择: 抢劫A/不抢劫A
// 如果抢劫A，那么不能抢劫节点B、C
// 如果不抢劫A，那么对于左子树，只要最大值；对于右子树，
func robTree(treeNode *TreeNode) (int, int) {
	if treeNode == nil {
		return 0, 0
	}
	robLeft, skipLeft := robTree(treeNode.Left)
	robRight, skipRight := robTree(treeNode.Right)
	rob := skipLeft + skipRight + treeNode.Val
	skip := getMaxInt(robLeft, skipLeft) + getMaxInt(robRight, skipRight)

	return rob, skip
}

func main() {
	s := []int{5, 1, 2, 3}
	fmt.Printf("maxLen:%d\n", rob1D(s))
	fmt.Printf("maxLen:%d\n", robCircle(s))
}
