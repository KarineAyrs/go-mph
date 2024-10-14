package list

import "fmt"

type node struct {
	val  any
	next *node
}

type List struct {
	first *node
	last  *node
	size  int
}

func NewList() *List {
	return &List{
		first: nil,
		last:  nil,
		size:  0,
	}
}

func (l *List) PushBack(el any) {
	p := &node{
		val: el,
	}
	if l.IsEmpty() {
		l.first = p
		l.last = p
		return
	}
	l.last.next = p
	l.last = p
	l.size++
}
func (l *List) PushFront(el any) {
	p := &node{
		val: el,
	}
	if l.IsEmpty() {
		l.first = p
		l.last = p
		return
	}
	p.next = l.first
	l.first = p
	l.size++
}
func (l *List) PopBack() {
	if l.IsEmpty() {
		return
	}
	if l.first == l.last {
		l.PopFront()
		return
	}
	p := l.first
	for p.next != l.last {
		p = p.next
	}
	p.next = nil
	l.last = p
	l.size--
}
func (l *List) PopFront() {
	if l.IsEmpty() {
		return
	}
	p := l.first
	l.first = p.next
	p = nil
	l.size--
}
func (l *List) Size() int {
	return l.size
}
func (l *List) IsEmpty() bool {
	return l.first == nil
}

func (l *List) Print() {
	if l.IsEmpty() {
		return
	}
	p := l.first
	for p != nil {
		fmt.Println(p.val)
		p = p.next
	}
}
