package linkedlist

import "fmt"

type SinglyLinkedList[T any] struct {
	head *SinglyNode[T]
	size int

	defaultValValue T
}

// NewSinglyLinkedList is the function to get a default singly linked list.
// This function returns a linked list with a empty head node.
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: &SinglyNode[T]{},
		size: 1,
	}
}

// GetHead is the function to get the first node of the list.
// If there is no head node it returns an error.
func (l *SinglyLinkedList[T]) GetHead() (*SinglyNode[T], error) {
	if l.head == nil {
		return nil, fmt.Errorf("Head node doesn't exist")
	}
	return l.head, nil
}

func (l *SinglyLinkedList[T]) GetTail() (*SinglyNode[T], error) {
	if l.head == nil {
		return nil, fmt.Errorf("Tail node doesn't exist")
	}
	return mustFind(l.size-1, l.head), nil
}

// PushFront is the function to insert a value at the head of the linked list.
func (l *SinglyLinkedList[T]) PushFront(value T) {
	n := &SinglyNode[T]{
		Value: value,
		next:  l.head,
	}
	l.head = n
	l.size++
}

// PushBack is the function to insert a value at the tail of the linked list
// This function checks every node (recursively) so it may be slow!
func (l *SinglyLinkedList[T]) PushBack(value T) {
	if l.head == nil {
		l.head = &SinglyNode[T]{
			Value: value,
		}
		l.size++
		return
	}

	n := mustFind(l.size-1, l.head)

	n.next = &SinglyNode[T]{
		Value: value,
	}
	l.size++
}

// PopFront is the function to remove the node and return the value at the head of the linked list.
func (l *SinglyLinkedList[T]) PopFront() (T, error) {
	if l.head == nil {
		return l.defaultValValue, fmt.Errorf("Head node doesn't exist")
	}
	v := l.head.Value
	l.head = l.head.next
	l.size--
	return v, nil
}

// PopFront is the function to remove the node and return the value at the head of the linked list.
func (l *SinglyLinkedList[T]) PopBack() (T, error) {
	if l.head == nil {
		return l.defaultValValue, fmt.Errorf("Head node doesn't exist")
	}

	n := mustFind(l.size-2, l.head)

	val := n.next.Value

	n.next = nil
	l.size--
	return val, nil
}

// MustGet is the function to get the value at a certain index.
// If the index is not valid the programm panics (see mustFind)
func (l *SinglyLinkedList[T]) MustGet(indx int) T {
	return mustFind(indx, l.head).Value
}

// MustSet is the function to set the value at a certain index.
// If the index is not valid the programm panics (see mustFind)
func (l *SinglyLinkedList[T]) MustSet(indx int, value T) {
	n := mustFind(indx, l.head)
	n.Value = value
}

// Size is the function to get the # of all the nodes in the list
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// mustFind is the helper function to search in the list recursively
func mustFind[T any](cindx int, n *SinglyNode[T]) *SinglyNode[T] {
	if cindx < 0 {
		panic("Invalid index")
	}

	if cindx == 0 {
		return n
	}

	nn, err := n.GetNext()
	if err != nil {
		panic("Invalid index")
	}

	cindx--
	return mustFind(cindx, nn)
}
