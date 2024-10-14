package list

type LinkedList interface {
	PushBack(T any)
	PushFront(T any)
	PopBack()
	PopFront()
	Size() int
	IsEmpty() bool
	Print()
}
