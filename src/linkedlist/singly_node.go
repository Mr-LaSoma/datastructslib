package linkedlist

import "fmt"

type SinglyNode[T any] struct {
	Value T
	next  *SinglyNode[T]
}

// GetNext is the function to get the next node of the list.
// If there is no next node it returns an error.
func (n *SinglyNode[T]) GetNext() (*SinglyNode[T], error) {
	if n.next == nil {
		return nil, fmt.Errorf("Next node doesn't exist")
	}
	return n.next, nil
}
