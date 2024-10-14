package graph

import "github.com/karineayrs/go-mph/internal/list"

type node struct {
	first list.LinkedList
	edges list.LinkedList
}
