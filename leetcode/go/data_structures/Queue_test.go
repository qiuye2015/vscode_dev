package data_structures

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	fmt.Println("len ", q.Len())
	fmt.Println("isEmpty ", q.IsEmpty())
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	fmt.Println("len ", q.Len())
	fmt.Println("isEmpty ", q.IsEmpty())
	for !q.IsEmpty() {
		fmt.Printf("%d ", q.Pop())
	}
	fmt.Println("")
	fmt.Println("len ", q.Len())
	fmt.Println("isEmpty ", q.IsEmpty())
}
