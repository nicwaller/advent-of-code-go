package queue

import "errors"

type Queue[T any] struct {
	items []T
	head  int
	count int
}

func New[T any](size int) Queue[T] {
	if size < 1 {
		panic("Queue must have >= 1 length")
	}
	return Queue[T]{
		items: make([]T, size),
	}
}

func (q *Queue[T]) Grow(margin int) {
	olditems := q.Items()
	q.items = make([]T, len(q.items)+margin)
	q.head = 0
	q.count = 0
	for i := 0; i < len(olditems); i++ {
		q.Push(olditems[i])
	}
}

func FromSlice[T any](slice []T) Queue[T] {
	q := New[T](len(slice))
	for _, v := range slice {
		q.Push(v)
	}
	return q
}

func (q *Queue[T]) Push(v T) {
	if q.count >= len(q.items) {
		q.Grow(q.count)
	}
	q.items[(q.head+q.count)%len(q.items)] = v
	q.count++
}

func (q *Queue[T]) MustPop() T {
	r, err := q.Pop()
	if err != nil {
		panic(err)
	}
	return r
}

func (q *Queue[T]) Pop() (T, error) {
	if q.count < 1 {
		var none T
		return none, errors.New("cannot pop empty queue")
	}
	r := q.items[q.head]
	q.head++
	q.head %= len(q.items)
	q.count--
	return r, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.count < 1 {
		var none T
		return none, errors.New("cannot peek empty queue")
	}
	return q.items[q.head], nil
}

func (q *Queue[T]) Length() int {
	return q.count
}

func (q *Queue[T]) Items() []T {
	ret := make([]T, q.count)
	for i := 0; i < q.count; i++ {
		ret[i] = q.items[(q.head+i)%len(q.items)]
	}
	return ret
}

func (q *Queue[T]) Reset() {
	q.head = 0
	q.count = 0
}
