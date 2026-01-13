package linkedlist

import "fmt"

type DoublyNode[T any] struct {
	Value      T
	next, prev *DoublyNode[T]
}

// GetNext is the function to get the next node of the list.
// If there is no next node it returns an error.
func (n *DoublyNode[T]) GetNext() (*DoublyNode[T], error) {
	if n.next == nil {
		return nil, fmt.Errorf("Next node doesn't exist")
	}
	return n.next, nil
}

// GetNext is the function to get the previous node of the list.
// If there is no previous node it returns an error.
func (n *DoublyNode[T]) GetPrev() (*DoublyNode[T], error) {
	if n.prev == nil {
		return nil, fmt.Errorf("Previous node doesn't exist")
	}
	return n.prev, nil
}
