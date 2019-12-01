package main

import "fmt"

func main() {
	//s := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	//s := [][]int{{1}, {2}, {3}, {4}}
	//s := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	// s := [][]int{{1, 2}, {5, 6}, {9, 10}, {13, 14}}

	s := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	printMatrix(s)
}

func printMatrix(s [][]int) {
	beginX, beginY := 0, 0
	rowNum := len(s)
	columnNum := len(s[0])
	if rowNum == 1 {
		for _, v := range s[0] {
			fmt.Printf("%d, ", v)
		}
		return
	}
	if columnNum == 1 {
		for _, v := range s {
			fmt.Printf("%d, ", v[0])
		}
		return
	}
	for beginX < rowNum && beginY < columnNum {
		curX, curY := beginX, beginY
		for curY < columnNum-1 {
			fmt.Printf("%d, ", s[curX][curY])
			curY++
		}
		for curX < rowNum-1 {
			fmt.Printf("%d, ", s[curX][curY])
			curX++
		}
		for curY > beginY {
			fmt.Printf("%d, ", s[curX][curY])
			curY--
		}
		for curX > beginX {
			fmt.Printf("%d, ", s[curX][curY])
			curX--
		}
		beginX++
		beginY++
		rowNum--
		columnNum--
	}
}
