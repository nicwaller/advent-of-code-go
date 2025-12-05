package v1

import (
	"iter"
)

type Iterator[T any] struct {
	iterator iter.Seq[T]
}

// Filter produces a new iterator that may skip over some elements
func (t Iterator[T]) Filter(include func(T) bool) Iterator[T] {
	return Iterator[T]{func(yield func(T) bool) {
		for v := range t.iterator {
			if include(v) {
				yield(v)
			}
		}
	}}
}

func (t Iterator[T]) Map(f func(orig T) T) Iterator[T] {
	return Map[T, T](t, f)
}

// Map transforms each element returned by an iterator, potentially to a different type
// This must be implemented as a first class function due to limitations of the Go type system
func Map[I any, O any](iter Iterator[I], f func(I) O) Iterator[O] {
	return Iterator[O]{func(yield func(O) bool) {
		iter.Each(func(v I) {
			yield(f(v))
		})
	}}
}

// Reduce consumes all elements from the iterator and combines them into a single result
func (t Iterator[T]) Reduce(basis T, reduce func(T, T) T) T {
	a := basis
	for v := range t.iterator {
		a = reduce(a, v)
	}
	return a
}

// Counting is a pass-through iterator that also increments a counter variable
func (t Iterator[T]) Counting(counter *int) Iterator[T] {
	if counter == nil {
		panic("nil counter")
	}
	return Iterator[T]{func(yield func(T) bool) {
		for v := range t.iterator {
			*counter++
			if !yield(v) {
				break
			}
		}
	}}
}

// Chain is a meta-iterator that yields elements from multiple iterators in order
func Chain[T any](iters ...Iterator[T]) Iterator[T] {
	return Iterator[T]{func(yield func(T) bool) {
		ListIterator(iters).Each(func(i Iterator[T]) {
			i.Each(func(v T) {
				yield(v)
			})
		})
	}}
}

func Chunk[T any](ite Iterator[T], chunkSize int) Iterator[[]T] {
	return Iterator[[]T]{func(yield func([]T) bool) {
		for {
			chunk := make([]T, 0, chunkSize)
			ite.iterator(func(v T) bool {
				chunk = append(chunk, v)
				return len(chunk) < chunkSize
			})
			if len(chunk) == 0 {
				// empty chunk and the source iterator is empty
				return
			}
			keepGoing := yield(chunk)
			if !keepGoing {
				// the consumer wants us to stop
				return
			}
			if len(chunk) < chunkSize {
				// partial chunk and the source iterator is empty
				return
			}
		}
	}}
}

func SlidingWindow[T any](ite Iterator[T], windowSize int) Iterator[[]T] {
	window := make([]T, 0, windowSize)

	// pre-fill the window
	ite.iterator(func(v T) bool {
		window = append(window, v)
		return len(window) < windowSize
	})

	return Iterator[[]T]{func(yield func([]T) bool) {
		// ... I'm not sure if this is quite right.
		if !yield(window) {
			return
		}

		ite.iterator(func(next T) bool {
			// PERF: using append potentially wastes a lot of memory
			// TODO: use a fancy ring queue instead
			window = append(window[1:], next)
			return yield(window)
		})
	}}
}

type IndexedValue[T any] struct {
	Index int
	Value T
}

func Enumerate[T any](ite Iterator[T]) Iterator[IndexedValue[T]] {
	index := 0
	return Iterator[IndexedValue[T]]{func(yield func(IndexedValue[T]) bool) {
		for v := range ite.iterator {
			iv := IndexedValue[T]{index, v}
			index++
			if !yield(iv) {
				break
			}
		}
	}}
}
