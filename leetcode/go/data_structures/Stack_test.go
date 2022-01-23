package data_structures

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()
	fmt.Println("len ", s.Len())
	fmt.Println("isEmpty ", s.IsEmpty())
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	fmt.Println("len ", s.Len())
	fmt.Println("isEmpty ", s.IsEmpty())
	for !s.IsEmpty() {
		fmt.Printf("%d ", s.Pop())
	}
	fmt.Println("")
	fmt.Println("len ", s.Len())
	fmt.Println("isEmpty ", s.IsEmpty())
}
