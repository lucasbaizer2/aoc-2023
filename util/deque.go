package util

type Deque[T any] struct {
	elements []T
}

func NewDeque[T any](capacity int) *Deque[T] {
	return &Deque[T]{
		elements: make([]T, 0, capacity),
	}
}

func (deque *Deque[T]) Len() int {
	return len(deque.elements)
}

func (deque *Deque[T]) Peek() T {
	if len(deque.elements) == 0 {
		panic("no elements remaining")
	}
	return deque.elements[0]
}

func (deque *Deque[T]) Pop() T {
	if len(deque.elements) == 0 {
		panic("no elements remaining")
	}
	value := deque.elements[0]
	deque.elements = deque.elements[1:]
	return value
}

func (deque *Deque[T]) Push(element T) {
	deque.elements = append(deque.elements, element)
}
