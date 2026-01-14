package linkedlist

type DoublyLinkedList[T any] struct {
	head, tail *DoublyNode[T]
	size int

	defaultValValue T
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	l := &DoublyLinkedList[T]{
		head: &SinglyNode[T]{},
		size: 1
	}
	l.tail = l.head
}

// GetHead is the function to get the first node of the list.
// If there is no head node it returns an error.
func (l *DoublyLinkedList[T]) GetHead() (*DoublyNode[T], error) {
	if l.head == nil {
		return nil, fmt.Errorf("Head node doesn't exist")
	}
	return l.head, nil
}

// GetTail is the function to get the last node of the list.
// If there is no tail node it returns an error.
func (l *DoublyLinkedList[T]) GetTail() (*DoublyNode[T], error) {
	if l.tail == nil {
		return nil, fmt.Errorf("Tail node doesn't exist")
	}
	return l.tail, nil
}

// PushFront is the function to insert a value at the head of the linked list.
func (l *DoublyLinkedList[T]) PushFront(value T) {
	n := &DoublyLinkedList{
		Value: value,
		next: l.head,
		prev: l.tail,
	}
	l.head = n
	if l.tail != nil {
		l.tail.next = n
	}
}

// PushBack is the function to insert a value at the tail of the linked list (see PushFront)
func (l *DoublyLinkedList[T]) PushBack(value T) {
	l.PushFront(value)
}

// PopFront is the function to remove the node and return the value at the head of the linked list.
func (l *DoublyLinkedList[T]) PopFront() (T, error) {
	if l.size == 0 {
		return l.defaultValValue, fmt.Errorf("Head node doesn't exist")
	}
	
	v := l.head.Value
	l.tail.next = l.head.next
	l.head.next.prev = l.tail
	return v
}

// PopFront is the function to remove the node and return the value at the head of the linked list.
func (l *DoublyLinkedList[T]) PopBack() (T, error) {
	if l.size == 0 {
		return l.defaultValValue, fmt.Errorf("Head node doesn't exist")
	}

	v := l.tail.Value
	l.head.prev = l.tail.prev
	l.tail.prev.next = l.head
	return v
}

// MustGet is the function to get the value at a certain index. 
// If the index is not valid the programm panics (see mustFind)
func (l *SinglyLinkedList[T]) MustGet(indx int) T {
	if indx == l.size {
		return l.tail.Value
	}
	return mustFind(index, l.head)
}

// MustSet is the function to set the value at a certain index. 
// If the index is not valid the programm panics (see mustFind)
func (l *SinglyLinkedList[T]) MustSet(indx int, value T) {
	if indx == l.size {
		l.tail.Value = value
		return
	}
	n := mustFind(index, l.head)
	n.Value = value
}


// Size is the function to get the # of all the nodes in the list
func (l *DoublyLinkedList[T]) Size() int {
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