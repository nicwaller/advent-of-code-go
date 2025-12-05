package v1

import (
	"iter"
	"slices"
)

func SeqIterator[T any](seq iter.Seq[T]) Iterator[T] {
	return Iterator[T]{iterator: seq}
}

func ListIterator[T any](list []T) Iterator[T] {
	return SeqIterator[T](slices.Values(list))
}

func StringIterator(s string) Iterator[string] {
	i := 0
	return SeqIterator[string](func(yield func(string) bool) {
		for i < len(s) {
			yield(s[i : i+1])
			i++
		}
	})
}

func RuneIterator(s string) Iterator[rune] {
	return SeqIterator[rune](func(yield func(rune) bool) {
		for i := 0; i < len(s); i++ {
			yield(rune(s[i]))
		}
	})
}

// start (inclusive) stop (exclusive)
func Range(start int, stop int) Iterator[int] {
	return RangeStepped(start, stop, 1)
}

// start (inclusive) stop (exclusive)
func RangeStepped(start int, stop int, step int) Iterator[int] {
	if step == 0 {
		panic("step cannot be zero")
	}
	i := start
	return Iterator[int]{func(yield func(int) bool) {
		for i < stop {
			ok := yield(i)
			i += step
			if !ok {
				break
			}
		}
	}}
}
