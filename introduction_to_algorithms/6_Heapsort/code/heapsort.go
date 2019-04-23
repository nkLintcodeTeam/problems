package main

import (
	"fmt"
)

type Heap struct {
	length int
	data   []int
}

func NewHeap(A []int) *Heap {
	return &Heap{
		data:   A,
		length: len(A),
	}
}

func (h *Heap) Elem(i int) int {
	return h.data[i]
}

func (h *Heap) Length() int {
	return h.length
}

func (h *Heap) IncrLength(d int) {
	h.length += d
}

func (h *Heap) Swap(a, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
}

func (h *Heap) Data() []int {
	return h.data
}

func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return i*2 + 1
}

func Right(i int) int {
	return i*2 + 2
}

func MaxHeapify(h *Heap, i int) {
	l := Left(i)
	r := Right(i)
	largest := i
	if l < h.Length() && h.Elem(l) > h.Elem(largest) {
		largest = l
	}
	if r < h.Length() && h.Elem(r) > h.Elem(largest) {
		largest = r
	}

	if i != largest {
		h.Swap(i, largest)
		MaxHeapify(h, largest)
	}
}

func BuildMaxHeap(h *Heap) {
	for i := h.Length() / 2; i >= 0; i = i - 1 {
		MaxHeapify(h, i)
	}
}

// Heapsort
func Heapsort(A []int) []int {
	h := NewHeap(A)
	BuildMaxHeap(h)
	for i := h.Length() - 1; i > 0; i = i - 1 {
		h.Swap(0, i)
		h.IncrLength(-1)
		MaxHeapify(h, 0)
	}
	return h.Data()
}

func localT() {
	a := []int{3, 9, 5, 8, 1, 0, 2, 7, 4, 6}
	b := Heapsort(a)
	for idx, v := range b {
		fmt.Printf("idx:%d, v:%d\n", idx, v)
	}
}

func main() {
	localT()
}
