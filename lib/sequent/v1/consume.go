package v1

import "fmt"

// Non-terminal consumer functions
// We don't return an iterator for chaining
// But we also don't consume everything from the iterator
// So it's likely the iterator will still have data in it after calling these

func (t Iterator[T]) TakeOne() T {
	var first *T
	t.iterator(func(v T) bool {
		first = &v
		return false
	})
	if first == nil {
		panic("none to take first")
	}
	return *first
}

// TakeN consumes exactly N elements from the iterator
// If fewer than N elements are available, it panics
func (t Iterator[T]) TakeN(n int) []T {
	if n == 0 {
		return []T{}
	}
	collected := make([]T, 0, n)
	t.iterator(func(v T) bool {
		collected = append(collected, v)
		return len(collected) < n
	})
	if len(collected) < n {
		panic(fmt.Errorf("only collected %d/%d elements", len(collected), n))
	}
	return collected
}

// TODO: TakeWhile() (maybe needs buffered reader?)
