package queue

import "fmt"

// Queue represents a generic FIFO (First In First Out) queue.
// V can be any type. Internally, it uses a slice to store elements.
type Queue[V any] struct {
	// underlying slice holding the queue elements
	values []V

	// default zero value for type V, returned on errors
	defaultValValue V
}

func NewQueue[V any]() *Queue[V] {
	return &Queue[V]{}
}

// Enqueue adds a new element to the end of the queue.
func (q *Queue[V]) Enqueue(value V) {
	q.values = append(q.values, value)
}

// Dequeue removes and returns the front element of the queue.
// Returns an error if the queue is empty. This function have to
// shift all the values in the underlying slice so it may be slow!
func (q *Queue[V]) Dequeue() (V, error) {
	if q.IsEmpty() {
		return q.defaultValValue, fmt.Errorf("Can't dequeue on empty Queue")
	}
	val := q.values[0]
	q.values = q.values[1:]
	return val, nil
}

// Peek returns the front element without removing it.
// Returns an error if the queue is empty.
func (q *Queue[V]) Peek() (V, error) {
	if q.IsEmpty() {
		return q.defaultValValue, fmt.Errorf("Can't peek on empty Queue")
	}
	return q.values[0], nil
}

// IsEmpty returns true if the queue contains no elements.
func (q *Queue[V]) IsEmpty() bool {
	return len(q.values) == 0
}

// Len returns the number of elements currently in the queue.
func (q *Queue[V]) Len() int {
	return len(q.values)
}
