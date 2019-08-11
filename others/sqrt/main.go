package main

import "fmt"

func equal(a, b float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -1 * diff
	}
	if diff < 0.0000001 {
		return true
	}
	return false
}

func sqrt(a float64) float64 {
	low := float64(0)
	high := a
	mid := (low + high) / 2
	cur := mid * mid

	for !equal(a, cur) {
		fmt.Printf("a:%+v, cur:%+v, mid:%+v, low:%+v, high:%+v\n", a, cur, mid, low, high)
		if cur < a {
			low = mid
		} else if cur > a {
			high = mid
		}
		mid = (low + high) / 2
		cur = mid * mid
	}
	return mid
}

func main() {
	// s := []float64{0.5}
	s := []float64{1}
	for _, v := range s {
		curRst := sqrt(v)
		fmt.Printf("v: %+v, curRst: %+v\n", v, curRst)
	}
	fmt.Println("vim-go")
}
