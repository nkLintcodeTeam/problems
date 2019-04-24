package main

import (
	"fmt"
)

const (
	ColorWhite = iota
	ColorGray
	ColorBlack
)

type FIFO struct {
	data []*Node
}

func NewFIFO(s int) *FIFO {
	return &FIFO{
		data: make([]*Node, 0, s),
	}
}

func (f *FIFO) Write(n *Node) {
	f.data = append(f.data, n)
}

func (f *FIFO) Read() *Node {
	if f.Empty() {
		return nil
	}
	res := f.data[0]
	f.data = f.data[1:]
	return res
}

func (f *FIFO) Empty() bool {
	return len(f.data) == 0
}

func (f *FIFO) Print() {
	fmt.Printf("fifo: ")
	for _, v := range f.data {
		fmt.Printf("%d,", v.Value)
	}
	fmt.Printf("\n")
}

type Node struct {
	Value    int
	Color    int
	Parent   *Node
	Distance int
}

func NewNode(v int) *Node {
	return &Node{
		Value: v,
	}
}

func convertIntToNodeGraph(g map[int][]int) (map[*Node][]*Node, map[int]*Node) {
	int2NodeLookup := make(map[int]*Node)
	ng := make(map[*Node][]*Node, len(g))
	for src, children := range g {
		if _, ok := int2NodeLookup[src]; !ok {
			int2NodeLookup[src] = NewNode(src)
		}
		for _, child := range children {
			if _, ok := int2NodeLookup[child]; !ok {
				int2NodeLookup[child] = NewNode(child)
			}
		}
	}

	for src, children := range g {
		srcNode := int2NodeLookup[src]
		if _, ok := ng[srcNode]; !ok {
			ng[srcNode] = make([]*Node, 0, len(children))
		}
		for _, child := range children {
			ng[srcNode] = append(ng[srcNode], int2NodeLookup[child])
		}
	}

	return ng, int2NodeLookup
}

func BFS(g map[int][]int, src int) []*Node {
	fifo := NewFIFO(len(g))
	ng, int2NodeLookup := convertIntToNodeGraph(g)
	startNode := int2NodeLookup[src]
	startNode.Color = ColorGray
	fifo.Write(startNode)

	res := make([]*Node, 0)

	for !fifo.Empty() {
		fifo.Print()
		u := fifo.Read()
		if children, ok := ng[u]; ok {
			for _, v := range children {
				if v.Color == ColorWhite {
					v.Color = ColorGray
					v.Distance++
					v.Parent = u
					fifo.Write(v)
				}
			}
			u.Color = ColorBlack
			res = append(res, u)
		}
	}
	return res
}

func main() {
	/*
			   0
			 1    2
		   3     4 5
		  6 7
	*/
	g := map[int][]int{
		0: []int{1, 2},
		1: []int{0, 3},
		2: []int{0, 4, 5},
		3: []int{1, 6, 7},
		4: []int{2},
		5: []int{2},
		6: []int{3},
		7: []int{3},
	}

	nodes := BFS(g, 0)
	for _, n := range nodes {
		parent := -1
		if n.Parent != nil {
			parent = n.Parent.Value
		}
		fmt.Printf("node: %#v, parent:%d\n", n, parent)
	}
	fmt.Printf("Running ...")
}
