package v1

// Each uses a callback to handle each item produced by the iterator.
// This way is generally preferred because it works on infinite streams.
func (t Iterator[T]) Each(do func(v T)) {
	if do == nil {
		do = func(v T) {}
	}
	for v := range t.iterator {
		do(v)
	}
}

// List consumes the entire iterator, growing a list in memory.
// This only works for small finite iterators.
func (t Iterator[T]) List() []T {
	list := make([]T, 0)
	for v := range t.iterator {
		list = append(list, v)
	}
	return list
}

// Count consumes the entire iterator, discarding each item.
// This only works for finite iterators.
// As an alternative you can use Counting instead.
func (t Iterator[T]) Count() int {
	counter := 0
	for _ = range t.iterator {
		counter++
	}
	return counter
}

func Sum[T int](ite Iterator[T]) (sum T) {
	for v := range ite.iterator {
		sum += v
	}
	return
}
