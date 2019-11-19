package main

import (
	"fmt"
)

func main() {
	m := make([]int, 0, Fact(5))
	nums := []int{1, 2, 7, 8}
	Perm(nums, 4, func(v []int) {
		x := 0
		for _, vv := range v {
			x = x*10 + vv
		}
		m = append(m, x)
	})
	k := make(map[int]bool)
	for _, v := range m {
		k[v] = true
	}

	fmt.Println("在以下各数中有5个数的和仍为以下各数中的数")
	for i, v := range m {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println("\n\n解:")
	Comb(m, 5, func(v []int) {
		c := 0
		for _, vv := range v {
			c += vv
		}
		// fmt.Println(c)
		if _, e := k[c]; e {
			for i, vv := range v {
				if i > 0 {
					fmt.Print(" + ")
				}
				fmt.Print(vv)
			}
			fmt.Printf(" = %v\n", c)
		}
	})
}

// 阶乘
func Fact(n int) int {
	v := 1
	for ; n > 1; n-- {
		v *= n
	}
	return v
}

// 全排列 m: 目标slice f: 回调
func FullPerm(m []int, f func([]int)) {
	N := len(m)
	o := make([]int, N) // 输出
	d := make([]int, N) // 索引
	for i := 0; i < N; i++ {
		d[i] = i
	}
	pEnd := N - 1 // 替换点
	p := pEnd     // 替换终点
	q := 0        // 替换点的下一个数
	pMax := 0     // 替换点后比替换点大的最小数
	for i, v := range d {
		o[i] = m[v]
	}
	f(o)
	// O(N!)
	for p != 0 {
		q = p
		p--
		if d[p] < d[q] {
			// 找与替换点交换的点
			pMax = pEnd
			for d[p] >= d[pMax] {
				pMax--
			}
			d[p], d[pMax] = d[pMax], d[p] // 交换
			// 将替换点后所有数进行反转
			for i, j := q, pEnd+q; i < (q+pEnd+1)/2; i++ {
				d[i], d[j-i] = d[j-i], d[i]
			}
			p = pEnd //将替换点置最后一个点，开始下一轮循环
			// 输出
			for i, v := range d {
				o[i] = m[v]
			}
			f(o)
		}
	}
}

// 排列 m: 目标slice k: 排列数 f: 回调
func Perm(m []int, k int, f func([]int)) {
	if k == len(m) {
		FullPerm(m, f)
	} else {
		Comb(m, k, func(v []int) {
			FullPerm(v, f)
		})
	}
}

// 组合 m: 目标slice k: 组合数 f: 回调
func Comb(m []int, k int, f func([]int)) {
	N := len(m)
	d := make([]int, k)
	v := make([]int, k)
	for i := 0; i < k; i++ {
		d[i] = i
	}

	for {
		for i, g := range d {
			v[i] = m[g]
		}
		f(v)
		// 终止 d1 >= n-k
		if d[0] >= N-k {
			return
		}
		// p: 自增维度
		for p := 1; p <= k; p++ {
			d[k-p] += 1 // d(n)++
			for i := p; i > 1; i-- {
				d[k-i+1] = d[k-i] + 1 // d(n) = d(n-1) + 1
			}
			if d[k-p] < N-p { // if d(n) < N
				break
			}
		}
	}
}
