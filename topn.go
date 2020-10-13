// Package topn is an experimental pkg for computing the top N items
// from a list.
package topn

import (
	"container/heap"
	"sort"
)

type elt struct {
	n        int // value
	maxIndex int // index in the max heap for fix operations
}

// MaxInts returns the max n ints from x.
func MaxInts(x []int, n int) []int {
	if len(x) <= n {
		sort.Ints(x) // changes x
		return x
	}

	elts := make([]elt, n)
	min := make(minEltHeap, n)
	max := make(maxEltHeap, n)
	for i := range elts {
		p := &elts[i]

		p.n = x[i] // init with the first n values
		p.maxIndex = i
		min[i] = p
		max[i] = p
	}

	heap.Init(&min)
	heap.Init(&max)

	// Go through the rest of the items and check if they should be
	// added.
	e := min[0]
	l := e.n
	for _, z := range x[n:] {
		if l < z {
			// swap out the value
			e.n = z
			heap.Fix(&min, 0)
			heap.Fix(&max, e.maxIndex)

			e = min[0]
			l = e.n
		}
	}

	out := make([]int, n)
	for i := 0; i < n; i++ {
		e := heap.Pop(&max).(*elt)
		out[i] = e.n
	}
	return out
}

// An minEltHeap is a min-heap of ints.
type minEltHeap []*elt

func (h minEltHeap) Len() int           { return len(h) }
func (h minEltHeap) Less(i, j int) bool { return h[i].n < h[j].n }
func (h minEltHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minEltHeap) Push(x interface{}) {
	panic("should not use")
}

func (h *minEltHeap) Pop() interface{} {
	panic("should hot use")
}

// An maxEltHeap is a min-heap of ints.
type maxEltHeap []*elt

func (h maxEltHeap) Len() int           { return len(h) }
func (h maxEltHeap) Less(i, j int) bool { return h[i].n > h[j].n }
func (h maxEltHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]

	h[j].maxIndex = j
	h[i].maxIndex = i
}

func (h *maxEltHeap) Push(x interface{}) {
	panic("should not use")
}

func (h *maxEltHeap) Pop() interface{} {
	n := len(*h)
	elt := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return elt
}
