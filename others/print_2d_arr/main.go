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

/*
func print2d(s [][]int) {
	curX := 0
	curY := 0
	delta := [][]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}
	curWidth := len(s[0])
	curHeight := len(s)

	curPos := 0
	for curWidth > 0 || curHeight > 0 {
		dX := delta[curPos][0]
		dY := delta[curPos][1]

		if

	}
}
*/

func printMatrix(s [][]int) {
	beginX, beginY := 0, 0
	rowNum := len(s) - 1
	columnNum := len(s[0]) - 1
	if rowNum == 0 {
		for _, v := range s[0] {
			fmt.Printf("%d, ", v)
		}
		return
	}
	if columnNum == 0 {
		for _, v := range s {
			fmt.Printf("%d, ", v[0])
		}
		return
	}
	for beginX <= rowNum && beginY <= columnNum {
		curX, curY := beginX, beginY
		for curY < columnNum {
			fmt.Printf("%d, ", s[curX][curY])
			curY++
		}
		for curX < rowNum {
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
