package main

import (
	"fmt"
	"math/rand"
)

func selectKElem(data []int, k int) []int {
	res := make([]int, 0, k)
	if len(data) <= k {
		for _, v := range data {
			res = append(res, v)
		}
		return res
	}
	var idx int
	for idx = 0; idx < k; idx = idx + 1 {
		res = append(res, data[idx])
	}

	for ; idx < len(data); idx = idx + 1 {
		r := rand.Intn(idx + 1)
		if r < k {
			res[r] = data[idx]
		}
	}
	return res
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	num := make([]int, 10, 10)
	for i := 0; i < 10000000; i = i + 1 {
		rst := selectKElem(a, 3)
		for _, v := range rst {
			num[v] += 1
		}
	}
	fmt.Printf("num: %+v\n", num)
}
