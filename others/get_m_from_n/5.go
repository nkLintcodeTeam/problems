package main

import "fmt"

func f(N, M int) [][]int {
	resList := make([][]int, 0)
	all := make([]int, N, N)
	for i := 0; i < N; i = i + 1 {
		all[i] = i
	}
	res := make([]int, 0, M)
	var comb func(k, m int)

	comb = func(k, m int) {
		if len(res) == M {
			newRes := make([]int, M, M)
			copy(newRes, res)
			resList = append(resList, newRes)
		} else {
			if k < N {
				res = append(res, all[k])
				comb(k+1, m-1)
				res = res[:len(res)-1]
				comb(k+1, m)
			}
		}
	}

	comb(0, M)
	return resList
}

func PrintArr(s []int) {
	for _, v := range s {
		fmt.Printf("%d, ", v)
	}
	fmt.Println("")
}

func main() {
	res := f(5, 3)
	for _, v := range res {
		fmt.Println(v)
	}
}
