package iter

// I really like itertools -NW
// https://docs.python.org/3/library/itertools.html

type Iterator[T any] func(**T) bool

func ListIterator[T any](list *[]T) Iterator[T] {
	index := 0
	return func(ref **T) bool {
		if index < len(*list) {
			*ref = &(*list)[index]
			index++
			return true
		} else {
			return false
		}
	}
}

func Chain[T any](iters ...Iterator[T]) Iterator[T] {
	remainingIters := iters
	var valPtr *T
	return func(ref **T) bool {
		if remainingIters[0](&valPtr) {
			*ref = valPtr
			return true
		} else {
			remainingIters = remainingIters[1:]
			if len(remainingIters) == 0 {
				return false
			}
			if remainingIters[0](&valPtr) {
				*ref = valPtr
				return true
			} else {
				return false
			}
		}
	}
}
