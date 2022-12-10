package iter

import (
	"advent-of-code/lib/f8l"
	"advent-of-code/lib/queue"
	"errors"
	"fmt"
	"github.com/ernestosuarez/itertools"
)

// I really like itertools -NW
// https://docs.python.org/3/library/itertools.html

type Iterator[T any] struct {
	Next  func() bool
	Value func() T
}

func StringIterator(s string, step int) Iterator[string] {
	if step <= 0 {
		panic(step)
	}
	offset := -step
	return Iterator[string]{
		Next: func() bool {
			offset += step
			return offset+step <= len(s)
		},
		Value: func() string {
			return s[offset : offset+step]
		},
	}
}

func ListIterator[T any](list []T) Iterator[T] {
	index := 0
	var val T
	return Iterator[T]{
		Next: func() bool {
			if index < len(list) {
				val = (list)[index]
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

func Map[I any, O any](iter Iterator[I], mapper func(orig I) O) Iterator[O] {
	var val O
	return Iterator[O]{
		Next: func() bool {
			if !iter.Next() {
				return false
			}
			val = mapper(iter.Value())
			return true
		},
		Value: func() O {
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
//
//	iter/iter.go:136:65: T instantiated as []T
//
// It's also not possible to implement Pairwise() with multiple-return :(
func SlidingWindow[T any](windowSize int, iter Iterator[T]) Iterator[[]T] {
	buf := queue.New[T](windowSize + 2)
	var none T
	buf.Push(none)
	// pre-fill the window (abort if there are too few items)
	for i := 0; i < windowSize-1; i++ {
		if !iter.Next() {
			return EmptyIterator[[]T]()
		}
		_ = buf.Push(iter.Value())
	}
	// return windowed results
	// TODO: clever circular queue bullshit
	return Iterator[[]T]{
		Next: func() bool {
			if !iter.Next() {
				return false
			}
			_ = buf.Push(iter.Value())
			_, _ = buf.Pop()
			return true
		},
		Value: func() []T {
			return buf.Items()
		},
	}

}

type IndexedValue[T any] struct {
	index int
	value T
}

func Repeat[T any](iter Iterator[T]) Iterator[T] {
	all := iter.List()
	index := -1
	return Iterator[T]{
		Next: func() bool {
			index = (index + 1) % len(all)
			return true
		},
		Value: func() T {
			return all[index]
		},
	}
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

func (iter Iterator[T]) Skip(count int) error {
	for i := 0; i < count; i++ {
		if !iter.Next() {
			return errors.New("nothing left to skip")
		}
	}
	return nil
}

func (iter Iterator[T]) Go() Iterator[T] {
	for iter.Next() {
	}
	return iter
}

func (iter Iterator[T]) Echo() Iterator[T] {
	return Iterator[T]{
		Next: func() bool {
			if iter.Next() {
				fmt.Println(iter.Value())
				return true
			}
			return false
		},
		Value: func() T {
			return iter.Value()
		},
	}
}

func (iter Iterator[T]) MustTakeArray(count int) []T {
	arr, err := iter.TakeArray(count)
	if err != nil {
		panic(err)
	}
	return arr
}

func (iter Iterator[T]) TakeArray(count int) ([]T, error) {
	ret := make([]T, count)
	for i := 0; i < count; i++ {
		if !iter.Next() {
			return []T{}, errors.New("nothing left to take")
		}
		ret[i] = iter.Value()
	}
	return ret, nil
}

func (iter Iterator[T]) TakeFirst() T {
	tk, _ := iter.TakeArray(1)
	return tk[0]
}

func (iter Iterator[T]) TakeWhile(condition func(v T) bool) Iterator[T] {
	return Iterator[T]{
		Next: func() bool {
			if !iter.Next() {
				return false
			}
			v := iter.Value()
			return condition(v)
		},
		Value: func() T {
			return iter.Value()
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

func (iter Iterator[T]) Count() int {
	c := 0
	for iter.Next() {
		c++
	}
	return c
}

func (iter Iterator[T]) Counting(c *int) Iterator[T] {
	return Iterator[T]{
		Next: func() bool {
			if iter.Next() {
				*c++
				return true
			}
			return false
		},
		Value: func() T {
			return iter.Value()
		},
	}
}

// Cannot be implemented as a reciever method because of limits in generic typing system
func Chunk[T any](size int, iter Iterator[T]) Iterator[[]T] {
	chunk := make([]T, size)
	return Iterator[[]T]{
		Next: func() bool {
			for i := 0; i < size; i++ {
				if !iter.Next() {
					return false
				}
				chunk[i] = iter.Value()
			}
			return true
		},
		Value: func() []T {
			return chunk
		},
	}
}

func (iter Iterator[T]) Each(do func(T)) {
	for iter.Next() {
		do(iter.Value())
	}
}

func Transform[I any, O any](iter Iterator[I], fn func(i I) O) Iterator[O] {
	return Iterator[O]{
		Next: func() bool {
			return iter.Next()
		},
		Value: func() O {
			return fn(iter.Value())
		},
	}
}

func Range(start int, stop int) Iterator[int] {
	return RangeStepped(start, stop, 1)
}

func RangeStepped(start int, stop int, step int) Iterator[int] {
	cur := start - step
	return Iterator[int]{
		Next: func() bool {
			cur += step
			return cur < stop
		},
		Value: func() int {
			return cur
		},
	}
}

func Product[T any](a []T, b []T) Iterator[[2]T] {
	offset := 0
	terminus := len(a) * len(b)
	return Iterator[[2]T]{
		Next: func() bool {
			offset++
			return offset < terminus
		},
		Value: func() [2]T {
			return [2]T{
				a[offset/len(a)],
				b[offset%len(a)],
			}
		},
	}
}

func ProductV[T any](a ...[]T) Iterator[[]T] {
	index := make([]int, len(a))
	index[0] = -1

	return Iterator[[]T]{
		Next: func() bool {
			index[0]++
			for d := 0; d < len(a); d++ {
				if index[d] == len(a[d]) {
					index[d] = 0
					if d == len(a)-1 {
						return false
					}
					index[d+1]++
				}
			}
			return true
		},
		Value: func() []T {
			// Yes, I know all these memory copies are inefficient
			// But at least it is safe
			// consumers don't expect delivered results to change after calling .Next()
			res := make([]T, len(a))
			for d := 0; d < len(a); d++ {
				res[d] = a[d][index[d]]
			}
			return res
		},
	}
}

func CombinationsN[T any](original []T, m int) Iterator[[]T] {
	ch := itertools.GenCombinations(len(original), m)
	var cur []T
	return Iterator[[]T]{
		Next: func() bool {
			P, more := <-ch
			cur = f8l.Map[int, T](P, func(i int) T { return original[i] })
			return more
		},
		Value: func() []T {
			return cur
		},
	}
}

func Permute[T any](original []T) Iterator[[]T] {
	A := make([]T, len(original))
	copy(A, original)
	// This implements Heap's Algorithm
	// https://en.wikipedia.org/wiki/Heap%27s_algorithm
	swap := func(i int, j int) {
		tmp := A[i]
		A[i] = A[j]
		A[j] = tmp
	}

	// TODO: send the first, non-permuted one
	c := make([]int, len(A))
	i := 1

	sentOriginal := false

	return Iterator[[]T]{
		Next: func() bool {
			if !sentOriginal {
				sentOriginal = true
				return true
			}
			//fmt.Printf("i=%d c=%v A=%v\n", i, c, A)
			//if i >= len(A) {
			//	return false
			//}
			for {
				if c[i] < i {
					if i%2 == 0 {
						swap(0, i)
					} else {
						swap(c[i], i)
					}
					//fmt.Printf("c[%d]++\n", i)
					c[i]++
					i = 1
					break
				} else {
					c[i] = 0
					i++
				}
				if i >= len(A) {
					return false
				}
			}
			return true
		},
		Value: func() []T {
			ret := make([]T, len(original))
			copy(ret, A)
			return ret
		},
	}

	//    // i acts similarly to a stack pointer
	//    i := 1;
	//    while i < n do
	//        if  c[i] < i then
	//            if i is even then
	//                swap(A[0], A[i])
	//            else
	//                swap(A[c[i]], A[i])
	//            end if
	//            output(A)
	//            // Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
	//            c[i] += 1
	//            // Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
	//            i := 1
	//        else
	//            // Calling generate(i+1, A) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
	//            c[i] := 0
	//            i += 1
	//        end if
	//    end while

}

// TODO: iter.Chunk(chunkSize int)
// iter.Flatten() would be useful in a dynamically typed language, but I don't think it makes sense for Go.
