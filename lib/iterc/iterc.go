package iterc

import (
	"advent-of-code/lib/queue"
	"fmt"
)

type Iterator[T any] struct {
	C <-chan T
}

func EmptyIterator[T any]() Iterator[T] {
	elements := make(chan T)
	close(elements)
	return Iterator[T]{
		C: elements,
	}
}

// iterate over each character in a string
func StringIterator(s string) Iterator[string] {
	elements := make(chan string)
	go func() {
		for i := range s {
			elements <- s[i : i+1]
		}
		close(elements)
	}()
	return Iterator[string]{
		C: elements,
	}
}

func ListIterator[T any](list []T) Iterator[T] {
	elements := make(chan T)
	go func() {
		for _, v := range list {
			elements <- v
		}
		close(elements)
	}()
	return Iterator[T]{
		C: elements,
	}
}

type KV[K comparable, V any] struct {
	Key   K
	Value V
}

func MapIterator[K comparable, V any](m map[K]V) Iterator[KV[K, V]] {
	elements := make(chan KV[K, V])
	go func() {
		for k, v := range m {
			elements <- KV[K, V]{Key: k, Value: v}
		}
		close(elements)
	}()
	return Iterator[KV[K, V]]{
		C: elements,
	}
}

func Chain[T any](iters ...Iterator[T]) Iterator[T] {
	elements := make(chan T)
	go func() {
		for _, iter := range iters {
			for e := range iter.C {
				elements <- e
			}
		}
		close(elements)
	}()
	return Iterator[T]{
		C: elements,
	}
}

// sadly, Go generics can't quite handle Map() as a struct method
func Map[I any, O any](iter Iterator[I], mapper func(orig I) O) Iterator[O] {
	elements := make(chan O)
	go func() {
		for e := range iter.C {
			elements <- mapper(e)
		}
		close(elements)
	}()
	return Iterator[O]{
		C: elements,
	}
}

func Reduce[T any](iter Iterator[T], reducer func(T, T) T, basis T) T {
	val := basis
	for e := range iter.C {
		val = reducer(val, e)
	}
	return val
}

func (iter Iterator[T]) Reduce(reducer func(T, T) T, basis T) T {
	return Reduce(iter, reducer, basis)
}

func SlidingWindow[T any](iter Iterator[T], windowSize int) Iterator[[]T] {
	elements := make(chan []T)
	q := queue.New[T](windowSize)

	// pre-fill the window
	for i := 0; i < windowSize; i++ {
		e, more := <-iter.C
		if !more {
			return EmptyIterator[[]T]()
		}
		q.Push(e)
	}
	go func() {
		elements <- q.Items()
		for e := range iter.C {
			_, _ = q.Pop()
			q.Push(e)
			elements <- q.Items()
		}
		close(elements)
	}()
	return Iterator[[]T]{
		C: elements,
	}
}

func List[T any](iter Iterator[T]) []T {
	list := make([]T, 0)
	for e := range iter.C {
		list = append(list, e)
	}
	return list
}

func (iter Iterator[T]) List() []T {
	return List(iter)
}

func Repeat[T any](iter Iterator[T], times int) Iterator[T] {
	elements := make(chan T)
	go func() {
		all := iter.List()
		for round := 0; round < times; round++ {
			for _, e := range all {
				elements <- e
			}
		}
		close(elements)
	}()
	return Iterator[T]{
		C: elements,
	}
}

func (iter Iterator[T]) Repeat(times int) Iterator[T] {
	return Repeat(iter, times)
}

// TODO: oscillate()

type IndexedValue[T any] struct {
	Index int
	Value T
}

func Enumerate[T any](iter Iterator[T]) Iterator[IndexedValue[T]] {
	return EnumerateFrom(iter, 0)
}

func EnumerateFrom[T any](iter Iterator[T], start int) Iterator[IndexedValue[T]] {
	elements := make(chan IndexedValue[T])
	go func() {
		i := start
		for e := range iter.C {
			elements <- IndexedValue[T]{Index: i, Value: e}
			i++
		}
		close(elements)
	}()
	return Iterator[IndexedValue[T]]{
		C: elements,
	}
}

func Filter[T any](iter Iterator[T], include func(T) bool) Iterator[T] {
	elements := make(chan T)
	go func() {
		for e := range iter.C {
			if include(e) {
				elements <- e
			}
		}
		close(elements)
	}()
	return Iterator[T]{
		C: elements,
	}
}

func (iter Iterator[T]) Filter(include func(T) bool) Iterator[T] {
	return Filter(iter, include)
}

// TODO: Partition() is omitted because it's too risky for deadlocks or memory pressure
// use ForEach if you want to do your own partitioning
// or call Filter() twice if you can iterate twice

func Take[T any](iter Iterator[T], count int) Iterator[T] {
	elements := make(chan T)
	go func() {
		for i := 0; i < count; i++ {
			e, more := <-iter.C
			if more {
				elements <- e
			} else {
				break
			}
		}
		close(elements)
	}()
	return Iterator[T]{
		C: elements,
	}
}

func (iter Iterator[T]) Take(count int) Iterator[T] {
	return Take(iter, count)
}

func Skip[T any](iter Iterator[T], count int) Iterator[T] {
	_ = Take(iter, count)
	return iter
}

func (iter Iterator[T]) Skip(count int) Iterator[T] {
	return Skip(iter, count)
}

func TakeOne[T any](iter Iterator[T]) (T, error) {
	l := Take(iter, 1).List()
	if len(l) == 0 {
		var empty T
		return empty, fmt.Errorf("none remaining")
	}
	return l[0], nil
}

func (iter Iterator[T]) TakeOne() (T, error) {
	return TakeOne(iter)
}

func MustTakeOne[T any](iter Iterator[T]) T {
	if v, err := TakeOne(iter); err == nil {
		return v
	} else {
		panic(err)
	}
}

func (iter Iterator[T]) MustTakeOne() T {
	return MustTakeOne(iter)
}

// FIXME: this doesn't work right. we lose an element!
func TakeWhile[T any](iter Iterator[T], test func(T) bool) Iterator[T] {
	elements := make(chan T)
	go func() {
		for e := range iter.C {
			if test(e) {
				elements <- e
			} else {
				break
			}
		}
		close(elements)
	}()
	return Iterator[T]{
		C: elements,
	}
}

// FIXME: this doesn't work right. we lose an element!
func (iter Iterator[T]) TakeWhile(test func(T) bool) Iterator[T] {
	return TakeWhile(iter, test)
}

func Chunk[T any](iter Iterator[T], size int) Iterator[[]T] {
	elements := make(chan []T)
	go func() {
		keepLooping := true
		for keepLooping {
			// make() intentionally inside the loop
			// we are passing ownership so we don't want to reuse it
			chunk := make([]T, size)
			end := len(chunk)
			for i := range chunk {
				if e, more := <-iter.C; more {
					chunk[i] = e
				} else {
					end = i
					keepLooping = false
					break
				}
			}
			elements <- chunk[:end]
		}
		close(elements)
	}()
	return Iterator[[]T]{
		C: elements,
	}
}

func ForEach[T any](iter Iterator[T], do func(T)) {
	for e := range iter.C {
		do(e)
	}
}

func (iter Iterator[T]) ForEach(do func(T)) {
	ForEach(iter, do)
}

func Range(startInclusive int, stopExclusive int) Iterator[int] {
	return RangeStepped(startInclusive, stopExclusive, 1)
}

func RangeStepped(startInclusive int, stopExclusive int, step int) Iterator[int] {
	elements := make(chan int)
	go func() {
		for i := startInclusive; i < stopExclusive; i += step {
			elements <- i
		}
		close(elements)
	}()
	return Iterator[int]{
		C: elements,
	}

}

func Sum(iter Iterator[int]) int {
	return Reduce(iter, func(i int, j int) int {
		return i + j
	}, 0)
}

func Product[T any](slices ...[]T) Iterator[[]T] {
	elements := make(chan []T)
	indices := make([]int, len(slices))

	go func() {
		current := func() []T {
			res := make([]T, 0, len(indices))
			for d := 0; d < len(indices); d++ {
				res = append(res, slices[d][indices[d]])
			}
			return res
		}

		next := func() bool {
			indices[0]++
			for d := 0; d < len(slices); d++ {
				if indices[d] == len(slices[d]) {
					indices[d] = 0
					if d == len(slices)-1 {
						return false
					}
					indices[d+1]++
				}
			}
			return true

		}

		elements <- current()
		for next() {
			elements <- current()
		}
		close(elements)
	}()

	return Iterator[[]T]{
		C: elements,
	}
}

func Count[T any](iter Iterator[T]) int {
	c := 0
	for range iter.C {
		c++
	}
	return c
}

func (iter Iterator[T]) Count() int {
	return Count(iter)
}
