package list_test

import (
	"github.com/karineayrs/go-mph/internal/list"
	"testing"
)

func TestList(t *testing.T) {
	l := list.NewList()
	l.PushBack(1)
	l.Print()

	l.PopFront()
	l.Print()

	l.PushFront([]int{1, 2, 3})
	l.PushBack(2)
	l.Print()

	l.PopBack()
	l.Print()
}

func BenchmarkList_PushBack(b *testing.B) {
	l := list.NewList()
	for n := 0; n < b.N; n++ {
		l.PushBack(n)
	}
}

func BenchmarkList_PushFront(b *testing.B) {
	l := list.NewList()
	for n := 0; n < b.N; n++ {
		l.PushFront(n)
	}
}
