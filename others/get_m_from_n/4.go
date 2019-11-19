package main

import (
	"fmt"
)

func comb(n, m int) {
	s := make([]int, n, n)
	for i := 0; i < n; i = i + 1 {
		s[i] = 0
		f(s, n, m, n)
	}
}

func display(s []int, n int) {
	for i := 0; i < n; i = i + 1 {
		fmt.Printf("%d, ", s[i])
	}
	fmt.Println("--")
}

func f(s []int, n, m, N int) {
	if m == 0 {
		display(s, N)
		return
	}
	for i := n - 1; i >= m-1; i = i - 1 {
		s[i] = 1
		f(s, i, m-1, N)
		s[i] = 0
	}
}

func main() {
	comb(3, 2)
}
