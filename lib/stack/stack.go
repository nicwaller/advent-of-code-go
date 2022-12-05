package stack

import (
	"advent-of-code/lib/iter"
	"errors"
)

type Stack[T any] struct {
	items  []T
	ptr    int // points to top item
	inited bool
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
	stack.inited = true
}

func (stack *Stack[T]) Push(v T) {
	if !stack.inited {
		panic("must use NewStack()")
	}
	stack.ptr++
	if stack.ptr < len(stack.items) {
		stack.items[stack.ptr] = v
	} else {
		stack.items = append(stack.items, v)
	}
}

func (stack *Stack[T]) MustPop() T {
	if !stack.inited {
		panic("must use NewStack()")
	}
	if stack.ptr == -1 {
		panic(errors.New("cannot Pop() empty stack"))
	}
	stack.ptr--
	return stack.items[stack.ptr+1]
}

func (stack *Stack[T]) MustPopN(n int) []T {
	ret, err := stack.PopN(n)
	if err != nil {
		panic(err)
	}
	return ret
}

func (stack *Stack[T]) Pop() (T, error) {
	if !stack.inited {
		panic("must use NewStack()")
	}
	if stack.ptr == -1 {
		var none T
		return none, errors.New("cannot Pop() empty stack")
	}
	stack.ptr--
	return stack.items[stack.ptr+1], nil
}

func (stack *Stack[T]) PopN(n int) ([]T, error) {
	ret := make([]T, n)
	var err error
	for i := 0; i < n; i++ {
		ret[i], err = stack.Pop()
		if err != nil {
			var none []T
			return none, err
		}
	}
	return ret, nil
}

func (stack *Stack[T]) Peek() T {
	if !stack.inited {
		panic("must use NewStack()")
	}
	return stack.items[stack.ptr]
}

func (stack *Stack[T]) Empty() bool {
	if !stack.inited {
		panic("must use NewStack()")
	}
	return stack.ptr == -1
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

func (stack *Stack[T]) List() []T {
	c := make([]T, len(stack.items))
	copy(c, stack.items)
	return c
}
