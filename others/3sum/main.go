package main

import (
	"fmt"
	"sort"
)

func ts(s []int) [][]int {
	res := make([][]int, 0)
	if len(s) < 3 {
		return res
	}

	sort.Ints(s)
	sLen := len(s)
	for i := 0; i < sLen; i = i + 1 {
		target := -1 * s[i]
		front := i + 1
		back := sLen - 1
		for front < back {
			curSum := s[front] + s[back]
			if curSum < target {
				front += 1
			} else if curSum > target {
				back -= 1
			} else {
				one := []int{s[i], s[front], s[back]}
				res = append(res, one)
				// is one[1]==s[front], not one[1]==s[front+1]
				for front < back && one[1] == s[front] {
					front += 1
				}
				for front < back && one[2] == s[back] {
					back = back - 1
				}

			}
		}

		for i+1 < sLen && s[i] == s[i+1] {
			i = i + 1
		}
	}
	return res
}

func main() {
	s := []int{1, 2, 3, -1, -2, -3}
	fmt.Printf("rst: %+v\n", ts(s))
}
