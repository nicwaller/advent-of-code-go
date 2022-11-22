package analyze

import "golang.org/x/exp/constraints"

type Box[T constraints.Ordered] struct {
	seenFirst bool
	Min       T
	Max       T
	Total     T
}

func (box *Box[T]) Put(v T) *Box[T] {
	if !box.seenFirst {
		box.Min = v
		box.Max = v
		box.seenFirst = true
		return box
	}

	box.Min = AnyMin(box.Min, v)
	box.Max = AnyMax(box.Max, v)
	box.Total += v
	return box
}

func AnyMin[T constraints.Ordered](a T, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func AnyMax[T constraints.Ordered](a T, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}
