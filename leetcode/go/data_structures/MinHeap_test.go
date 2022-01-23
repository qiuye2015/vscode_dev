package data_structures

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMinHeap_Push(t *testing.T) {
	ih := new(MinHeap)
	heap.Init(ih)
	heap.Push(ih, 1)
	heap.Pop(ih)
	for i := 0; i < 10; i++ {
		heap.Push(ih, i)
		fmt.Println("push", i, "Min", (*ih)[0], "heap", (*ih))
	}

	for i := 0; i < 10; i++ {
		x := heap.Pop(ih)
		fmt.Println("pop", i, "Min", x, "heap", *ih)

	}
}
