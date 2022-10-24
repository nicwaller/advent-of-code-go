package iter

// I really like itertools -NW
// https://docs.python.org/3/library/itertools.html

type Iterator[T any] struct {
	Next  func() bool
	Value func() T
}

func ListIterator[T any](list *[]T) Iterator[T] {
	index := 0
	var val T
	return Iterator[T]{
		Next: func() bool {
			if index < len(*list) {
				val = (*list)[index]
				index++
				return true
			}
			return false
		},
		Value: func() T {
			return val
		},
	}
}

func Chain[T any](iters ...Iterator[T]) Iterator[T] {
	remainingIters := iters
	var val T
	return Iterator[T]{
		Next: func() bool {
			if remainingIters[0].Next() {
				val = remainingIters[0].Value()
				return true
			} else {
				remainingIters = remainingIters[1:]
				if len(remainingIters) == 0 {
					return false
				}
				if remainingIters[0].Next() {
					val = remainingIters[0].Value()
					return true
				} else {
					return false
				}
			}
		},
		Value: func() T {
			return val
		},
	}
}

func (iter Iterator[T]) Map(maps ...func(orig T) T) Iterator[T] {
	var val T
	return Iterator[T]{
		Next: func() bool {
			if !iter.Next() {
				return false
			}
			val = iter.Value()
			for _, mapFn := range maps {
				val = mapFn(val)
			}
			return true
		},
		Value: func() T {
			return val
		},
	}
}

func (iter Iterator[T]) Reduce(reducer func(a T, b T) T, basis T) T {
	val := basis
	for iter.Next() {
		val = reducer(val, iter.Value())
	}
	return val
}

func (iter Iterator[T]) MapReduce(mapper func(orig T) T, reducer func(a T, b T) T, basis T) T {
	val := basis
	for iter.Next() {
		val = reducer(val, mapper(iter.Value()))
	}
	return val
}

func EmptyIterator[T any]() Iterator[T] {
	return Iterator[T]{
		Next: func() bool {
			return false
		},
		Value: func() T {
			var result T // get the zero value (can't use nil)
			return result
		},
	}
}

// SlidingWindow This cannot be implemented as a reciever method
// iter/iter.go:6:15: instantiation cycle:
//	iter/iter.go:136:65: T instantiated as []T
// It's also not possible to implement Pairwise() with multiple-return :(
func SlidingWindow[T any](windowSize int, iter Iterator[T]) Iterator[[]T] {
	window := make([]T, windowSize)
	index := 0
	// pre-fill the window (abort if there are too few items)
	for i := 0; i < windowSize-1; i++ {
		if !iter.Next() {
			return EmptyIterator[[]T]()
		}
		window[index] = iter.Value()
		index = (index + 1) % windowSize
	}
	// return windowed results
	// TODO: clever circular queue bullshit
	return Iterator[[]T]{
		Next: func() bool {
			if !iter.Next() {
				return false
			}
			window[index] = iter.Value()
			index = (index + 1) % windowSize
			return true
		},
		Value: func() []T {
			// I'm not convinced the memory performance of append() is good here
			return append(window[index:], window[:index]...)
		},
	}

}

type IndexedValue[T any] struct {
	index int
	value T
}

func Enumerate[T any](iter Iterator[T]) Iterator[IndexedValue[T]] {
	index := -1
	var value T
	return Iterator[IndexedValue[T]]{
		Next: func() bool {
			if !iter.Next() {
				return false
			}
			value = iter.Value()
			index++
			return true
		},
		Value: func() IndexedValue[T] {
			return IndexedValue[T]{
				index: index,
				value: value,
			}
		},
	}
}

func (iter Iterator[T]) Filter(include func(T) bool) Iterator[T] {
	var val T
	return Iterator[T]{
		Next: func() bool {
			for iter.Next() {
				if include(iter.Value()) {
					val = iter.Value()
					return true
				}
			}
			return false
		},
		Value: func() T {
			return val
		},
	}
}

func (iter Iterator[T]) Take(count int) Iterator[T] {
	taken := 0
	var val T
	return Iterator[T]{
		Next: func() bool {
			if taken >= count {
				return false
			}
			if !iter.Next() {
				return false
			}
			val = iter.Value()
			taken++
			return true
		},
		Value: func() T {
			return val
		},
	}
}

func (iter Iterator[T]) List() []T {
	list := make([]T, 0)
	for iter.Next() {
		list = append(list, iter.Value())
	}
	return list

}
