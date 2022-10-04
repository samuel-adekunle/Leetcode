/* How to maintain median invariant with two heaps
 */
package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []float32

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(float32))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []float32

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(float32))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianCalculator struct {
	Count   int
	MaxHeap *MaxHeap // keep track of the smallest numbers seen
	MinHeap *MinHeap // keep track of the largest numbers seen
	Median  float32
}

func NewMedianCalculator() *MedianCalculator {
	calc := &MedianCalculator{}
	calc.MaxHeap = &MaxHeap{}
	heap.Init(calc.MaxHeap)
	calc.MinHeap = &MinHeap{}
	heap.Init(calc.MinHeap)
	return calc
}

func (calc *MedianCalculator) insert(x float32) {
	// update count
	calc.Count += 1

	// update heaps
	if calc.MaxHeap.Len() == 0 || x < (*calc.MaxHeap)[0] {
		heap.Push(calc.MaxHeap, x)
	} else {
		heap.Push(calc.MinHeap, x)
	}

	// balance heaps
	for calc.MaxHeap.Len() < calc.MinHeap.Len()-1 {
		heap.Push(calc.MaxHeap, heap.Pop(calc.MinHeap))
	}

	for calc.MinHeap.Len() < calc.MaxHeap.Len()-1 {
		heap.Push(calc.MinHeap, heap.Pop(calc.MaxHeap))
	}

	// update median
	if calc.MaxHeap.Len() > calc.MinHeap.Len() {
		calc.Median = (*calc.MaxHeap)[0]
	} else if calc.MinHeap.Len() > calc.MaxHeap.Len() {
		calc.Median = (*calc.MinHeap)[0]
	} else {
		calc.Median = ((*calc.MaxHeap)[0] + (*calc.MinHeap)[0]) / 2
	}
}

func main() {
	calc := NewMedianCalculator()
	stream := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range stream {
		calc.insert(v)
		fmt.Printf("%v\n", calc)
	}
}
