package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Length(t *testing.T) {
	var err error

	q := New[int](2)
	if q.Length() != 0 {
		fmt.Println(q.Length())
		t.Errorf("Empty queue should have length==0\n")
	}
	if len(q.Items()) != 0 {
		t.Errorf("Items() has wrong length")
	}

	err = q.Push(123)
	if err != nil {
		t.Errorf("Queue must not return error when pushing first element")
	}
	if q.Length() != 1 {
		fmt.Println(q.Length())
		t.Errorf("Queue length should be 1 after adding item\n")
	}
	if len(q.Items()) != 1 {
		t.Errorf("Items() has wrong length")
	}
	if q.Items()[0] != 123 {
		t.Errorf("Items() has wrong value")
	}

	err = q.Push(456)
	if err != nil {
		t.Errorf("Queue must not return error when pushing second element")
	}
	if q.Length() != 2 {
		fmt.Println(q.Length())
		t.Errorf("Queue length should be 2 after adding second item\n")
	}
	if len(q.Items()) != 2 {
		t.Errorf("Items() has wrong length")
	}
	if q.Items()[0] != 123 {
		t.Errorf("Items() has wrong value")
	}
	if q.Items()[1] != 456 {
		t.Errorf("Items() has wrong value")
	}

	v, err := q.Peek()
	if v != 123 {
		t.Errorf("Got wrong value when peeking queue\n")
	}
	if err != nil {
		t.Errorf("Queue must not return error when pushing first element\n")
	}
	if q.Length() != 2 {
		fmt.Println(q.Length())
		t.Errorf("Queue length should not change after peeking\n")
	}

	v, err = q.Pop()
	if v != 123 {
		t.Errorf("Got wrong value when peeking queue\n")
	}
	if err != nil {
		t.Errorf("Queue must not return error when popping first element\n")
	}
	if q.Length() != 1 {
		fmt.Println(q.Length())
		t.Errorf("Queue length should decrease after pop\n")
	}
	if len(q.Items()) != 1 {
		t.Errorf("Items() has wrong length")
	}
	if q.Items()[0] != 456 {
		t.Errorf("Items() has wrong value")
	}

	v, err = q.Pop()
	if v != 456 {
		t.Errorf("Got wrong value when popping queue\n")
	}
	if err != nil {
		t.Errorf("Queue must not return error when popping second element\n")
	}
	if q.Length() != 0 {
		fmt.Println(q.Length())
		t.Errorf("Queue must pop() last item\n")
	}
	if len(q.Items()) != 0 {
		t.Errorf("Items() has wrong length")
	}
}
