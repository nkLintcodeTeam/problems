package main

import "fmt"

func f() {
	s := [][]string{
		{"A", "A", "B", "A"},
		{"A", "B", "B", "A"},
		{"A", "B", "B", "A"},
		{"A", "A", "B", "A"},
	}

	maxX := len(s)
	maxY := len(s[0])

	var search func(int, int)
	search = func(startX, startY int) {
		direction := [][]int{
			{0, 1},
			{1, 0},
			{0, -1},
			{-1, 0},
		}
		for startX < maxX && startY < maxY && startX >= 0 && startY >= 0 {
			fmt.Println("line25")
			if s[startX][startY] == "A" {
				s[startX][startY] = "C"
				for _, dir := range direction {
					search(startX+dir[0], startY+dir[1])
				}
			} else {
				return
			}
		}
	}
	search(0, 0)
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	f()
	fmt.Println("vim-go")
}
