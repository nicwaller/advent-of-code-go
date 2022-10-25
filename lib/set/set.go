package set

import "fmt"

type void interface{}

type Set[T comparable] struct {
	items map[T]void
}

func New[T comparable](items ...T) Set[T] {
	s := Set[T]{}
	s.items = make(map[T]void)
	for _, item := range items {
		s.Insert(item)
	}
	return s
}

func Union[T comparable](sets ...Set[T]) Set[T] {
	union := New[T]()
	for _, input := range sets {
		union.Extend(input.Items()...)
	}
	return union
}

func Intersection[T comparable](sets ...Set[T]) Set[T] {
	intersection := Union[T](sets...)
	for k := range intersection.items {
		for _, s := range sets {
			if !s.Has(k) {
				intersection.Remove(k)
			}
		}
	}
	return intersection
}

// Compare is like the `comm` program that finds common lines between files
func Compare[T comparable](left Set[T], right Set[T]) (Set[T], Set[T], Set[T]) {
	onlyLeft := New[T]()
	inBoth := New[T]()
	onlyRight := New[T]()
	for k := range left.items {
		if right.Contains(k) {
			inBoth.Add(k)
		} else {
			onlyLeft.Add(k)
		}
	}
	for k := range right.items {
		if inBoth.Contains(k) {
			// well... it should already be in there
		} else {
			onlyRight.Add(k)
		}
	}
	return onlyLeft, inBoth, onlyRight
}

func (set Set[T]) Add(val T) {
	set.Insert(val)
}

func (set Set[T]) Insert(val T) {
	// maybe alias this as Add()?
	var empty struct{}
	set.items[val] = empty
}

func (set Set[T]) Extend(items ...T) {
	for _, item := range items {
		set.Insert(item)
	}
}

func (set Set[T]) Remove(val T) {
	delete(set.items, val)
}

func (set Set[T]) Filter(keep func(item T) bool) {
	for k := range set.items {
		if !keep(k) {
			set.Remove(k)
		}
	}
}

func (set Set[T]) Contains(val T) bool {
	return set.Has(val)
}

func (set Set[T]) Has(val T) bool {
	// maybe alias this as Contains() ?
	_, ok := set.items[val]
	return ok
}

func (set Set[T]) Items() []T {
	keys := make([]T, len(set.items))
	i := 0
	for k := range set.items {
		keys[i] = k
		i++
	}
	return keys
}

func (set Set[T]) Size() int {
	return len(set.items)
}

//goland:noinspection GoMixedReceiverTypes
func (set Set[T]) String() string {
	return fmt.Sprint(set.Items())
}
