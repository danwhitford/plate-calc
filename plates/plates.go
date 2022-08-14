package plates

import (
	"container/heap"
	"fmt"
	"syscall/js"
)

type Path []float64
type PlateHeap []Path

var PLATES []float64 = []float64{25, 20, 15, 10, 5, 2.5, 2, 1.5, 1, 0.5}

func NewHeap(target float64) *PlateHeap {
	var paths PlateHeap = make([]Path, len(PLATES))
	for i, v := range PLATES {
		if v <= target {
			paths[i] = []float64{v}
		}
	}
	hp := &paths
	heap.Init(hp)
	return hp
}

func total(h Path) float64 {
	sum := float64(0)
	for _, v := range h {
		sum += v
	}
	return sum
}

func (h PlateHeap) Len() int      { return len(h) }
func (h PlateHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h PlateHeap) Less(i, j int) bool {
	return total(h[i]) > total(h[j])
}

func (h *PlateHeap) Push(x any) {
	*h = append(*h, x.(Path))
}

func (h *PlateHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func find_neighbours(target, tot float64) []float64 {
	ret := make([]float64, 0)
	for _, v := range PLATES {
		if tot+v <= target {
			ret = append(ret, v)
		}
	}
	return ret
}

func GetPlatesForWeight(target float64) []float64 {
	pq := NewHeap(target)

	for {
		head := heap.Pop(pq).(Path)
		fmt.Printf("Trying %v\n", head)
		tot := total(head)
		if tot == target {
			return head
		}

		neighbours := find_neighbours(target, tot)
		for _, v := range neighbours {
			sl := make(Path, len(head))
			copy(sl, head)
			sl = append(sl, v)
			heap.Push(pq, sl)
		}
	}
}

func GetPlatesForBar(target float64, bar int) []float64 {
	fbar := float64(bar)
	if target < fbar {
		panic("Weight is less than the bar!")
	}
	target -= fbar
	target /= 2
	return GetPlatesForWeight(target)
}

func GetPlatesForBarJS() js.Func {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		target := args[0].Float()
		bar := args[1].Int()
		sl := GetPlatesForBar(target, bar)
		ret := make([]interface{}, len(sl))
		for i, v := range sl {
			ret[i] = v
		}
		return ret
	})
	return jsFunc
}
