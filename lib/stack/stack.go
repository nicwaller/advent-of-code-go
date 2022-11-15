package stack

import (
	"advent-of-code/lib/iter"
	"errors"
)

type Stack[T any] struct {
	items []T
	ptr   int // points to top item
}

//goland:noinspection GoUnusedExportedFunction
func NewStack[T any]() Stack[T] {
	s := Stack[T]{}
	s.Init()
	return s
}

func (stack *Stack[T]) Init() {
	stack.items = make([]T, 0)
	stack.ptr = -1
}

func (stack *Stack[T]) Push(v T) {
	stack.ptr++
	if stack.ptr < len(stack.items) {
		stack.items[stack.ptr] = v
	} else {
		stack.items = append(stack.items, v)
	}
}

func (stack *Stack[T]) MustPop() T {
	if stack.ptr == -1 {
		panic(errors.New("cannot Pop() empty stack"))
	}
	stack.ptr--
	return stack.items[stack.ptr+1]
}

func (stack *Stack[T]) Pop() (T, error) {
	if stack.ptr == -1 {
		var none T
		return none, errors.New("cannot Pop() empty stack")
	}
	stack.ptr--
	return stack.items[stack.ptr+1], nil
}

func (stack *Stack[T]) Peek() T {
	return stack.items[stack.ptr]
}

// Iterator starts at the top of the stack
func (stack *Stack[T]) Iterator() iter.Iterator[T] {
	var cur T
	return iter.Iterator[T]{
		Next: func() bool {
			var err error
			cur, err = stack.Pop()
			return err == nil
		},
		Value: func() T {
			return cur
		},
	}
}
