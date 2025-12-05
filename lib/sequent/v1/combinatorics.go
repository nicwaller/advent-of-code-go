package v1

// Permute generates all possible permutations (orderings) of input objects
// This function implements Heap's algorithm:
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
func Permute[T any](original []T) Iterator[[]T] {
	A := make([]T, len(original))
	copy(A, original)
	swap := func(i int, j int) {
		tmp := A[i]
		A[i] = A[j]
		A[j] = tmp
	}

	// TODO: send the first, non-permuted one
	c := make([]int, len(A))
	i := 1

	sentOriginal := false

	done := false

	return Iterator[[]T]{iterator: func(yield func([]T) bool) {
		// TODO: make sure to yield a copy of the slice

		next := func() bool {
			if !sentOriginal {
				sentOriginal = true
				return true
			}
			for {
				if c[i] < i {
					if i%2 == 0 {
						swap(0, i)
					} else {
						swap(c[i], i)
					}
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
		}

		for !done && next() {
			result := make([]T, len(original))
			copy(result, A)
			if !yield(result) {
				return
			}
		}
		done = true
	}}
}
