package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	data     []int
	writeIdx int
	readIdx  int
	num      int
}

func NewQueue(l int) *Queue {
	return &Queue{
		data:     make([]int, l, l),
		writeIdx: 0,
		readIdx:  0,
		num:      0,
	}
}

func (q *Queue) Read() (int, error) {
	if q.num <= 0 {
		return 0, errors.New("empty queue")
	}
	if q.readIdx >= len(q.data) {
		q.readIdx = 0
	}
	res := q.data[q.readIdx]
	q.readIdx += 1
	q.num -= 1
	return res, nil
}

func (q *Queue) Append(v int) error {
	if q.writeIdx >= len(q.data) {
		q.writeIdx = 0
	}
	if q.readIdx == q.writeIdx && q.num > 0 {
		q.readIdx += 1
		q.num -= 1
	}
	q.data[q.writeIdx] = v
	q.writeIdx += 1
	q.num += 1
	return nil
}

func main() {
	q := NewQueue(3)
	fmt.Println("vim-go")
	// s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var rst int
	var err error
	q.Append(1)
	q.Append(2)
	q.Append(3)
	q.Append(4)
	q.Append(5)
	rst, err = q.Read()
	fmt.Printf("rst:%d, err:%+v\n", rst, err)
	rst, err = q.Read()
	fmt.Printf("rst:%d, err:%+v\n", rst, err)
	rst, err = q.Read()
	fmt.Printf("rst:%d, err:%+v\n", rst, err)
	rst, err = q.Read()
	fmt.Printf("rst:%d, err:%+v\n", rst, err)
}
