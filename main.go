package main

import (
	"container/heap"
	"fmt"
)

type Path []float32
type PlateHeap []Path

var PLATES []float32 = []float32{25, 20, 15, 10, 5, 2.5, 1, 0.5}

func NewHeap(target float32) *PlateHeap {
	var paths []Path = make([]Path, len(PLATES))
	for i, v := range PLATES {
		if v <= target {
			paths[i] = []float32{v}
		}
	}
	hp := &paths
	heap.Init(hp)
	return hp
}

func total(h Path) float32 {
	sum := float32(0)
	for _, v := range h {
		sum += v
	}
	return sum
}

func (h PlateHeap) Len() int      { return len(h) }
func (h PlateHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h PlateHeap) Less(i, j int) bool {
	a := (h.target - total(h.paths[i]))
	b := (h.target - total(h.paths[j]))

	return a < b
}

func (h *PlateHeap) Push(x any) {
	h.paths = append(h.paths, x.(Path))
}

func (h *PlateHeap) Pop() any {
	old := h.paths
	n := len(old)
	x := old[n-1]
	h.paths = old[0 : n-1]
	return x
}

func find_neighbours(target, tot float32) []float32 {
	ret := make([]float32, 0)
	for _, v := range PLATES {
		if tot+v <= target {
			ret = append(ret, v)
		}
	}
	return ret
}

func GetPlatesForWeight(weight float32) []float32 {
	pq := NewHeap(weight)

	for {
		head := heap.Pop(pq).(Path)
		fmt.Printf("Trying %v\n", head)
		tot := total(head)
		if tot == pq.target {
			return head			
		}

		neighbours := find_neighbours(pq.target, tot)
		for _, v := range neighbours {
			sl := make(Path, len(head))
			copy(sl, head)
			sl = append(sl, v)
			heap.Push(pq, sl)
		}
	}
}

func main() {
	answer := GetPlatesForWeight(42)
	fmt.Printf("Found %v\n", answer)
}
