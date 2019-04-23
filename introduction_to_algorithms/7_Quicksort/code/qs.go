package main

func qs(s []int, l, r int) {
	if l < r {
		m := partition(s, l, r)
		qs(s, l, m-1)
		qs(s, m+1, r)
	}
}

func partition(s []int, l, r int) int {
	i := l - 1
	pivot := s[r]
	for j := l; j < r; j++ {
		if s[j] <= pivot {
			i = i + 1
			s[i], s[j] = s[j], s[i]
		}
	}

	i = i + 1
	s[i], s[r] = s[r], s[i]
	return i
}

func main() {
	a := []int{2, 3, 1, 6, 9, 8, 4, 7, 5, 0}
	qs(a, 0, len(a)-1)

	for i := 0; i < len(a); i++ {
		println(a[i])
	}
}
