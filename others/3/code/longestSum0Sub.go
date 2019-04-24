package main

import "fmt"

func getLongestSum0Sub(arr []int) []int {
	helpArr := make([]int, len(arr), len(arr))
	helpArr[0] = arr[0]
	for i := 1; i < len(arr); i = i + 1 {
		helpArr[i] += helpArr[i-1] + arr[i]
	}

	lookup := make(map[int][]int)

	for idx, v := range helpArr {
		if _, ok := lookup[v]; !ok {
			lookup[v] = make([]int, 0)
		}
		lookup[v] = append(lookup[v], idx)
	}

	for k, v := range lookup {
		fmt.Printf("line22 k:%#v, v:%#v\n", k, v)
	}

	resLeft := -1
	resRight := -1
	resLen := 0

	for _, v := range lookup {
		if len(v) > 1 {
			curLen := v[len(v)-1] - v[0]
			if curLen > resLen {
				resLeft = v[0]
				resRight = v[len(v)-1]
				resLen = curLen
			}
		}
	}
	return []int{
		resLen,
		resLeft + 1,
		resRight,
	}
}

func main() {
	//            0   1   2   3  4  5   6   7   8   9  10  11
	//           -1  -2  -3  -4 -3 -2  -3  -4  -5  -6 -7  -6
	arr := []int{-1, -1, -1, -1, 1, 1, -1, -1, -1, -1, -1, 1}

	rst := getLongestSum0Sub(arr)

	fmt.Printf("rst: %#v\n", rst)
}
