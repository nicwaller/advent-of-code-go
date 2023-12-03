package iterc

import (
	"testing"
)

func TestEmptyIterator(t *testing.T) {
	r := EmptyIterator[interface{}]()
	e, more := <-r.C
	if more {
		t.Error("expected no more")
	}
	if e != nil {
		t.Error("what in the heck")
	}
}

func TestStringIterator(t *testing.T) {
	r := StringIterator("Hi")
	e := <-r.C
	if e != "H" {
		t.Error()
	}
	e = <-r.C
	if e != "i" {
		t.Error()
	}
	e, more := <-r.C
	if more {
		t.Error()
	}
}

func TestListIterator(t *testing.T) {
	r := ListIterator[string]([]string{"red", "green"})
	e := <-r.C
	if e != "red" {
		t.Error()
	}
	e = <-r.C
	if e != "green" {
		t.Error()
	}
	e, more := <-r.C
	if more {
		t.Error()
	}
}

func TestChain(t *testing.T) {
	r1 := ListIterator([]int{2, 3, 4})
	r2 := ListIterator([]int{5, 6, 7})
	rC := Chain(r1, r2)
	rL := rC.List()
	if len(rL) != 6 {
		t.Error()
	}
	if rL[0] != 2 {
		t.Error()
	}
	if rL[5] != 7 {
		t.Error()
	}
}

func TestMap(t *testing.T) {
	r1 := ListIterator([]int{2, 4, 6})
	rM := Map(r1, func(orig int) int {
		return orig + 1
	})
	rL := rM.List()
	if len(rL) != 3 {
		t.Error()
	}
	if rL[0] != 3 {
		t.Error()
	}
}

func TestReduce(t *testing.T) {
	r1 := ListIterator([]int{1, 2, 3, 4})
	rR := r1.Reduce(func(a, b int) int {
		return a * b
	}, 1)
	if rR != 24 {
		t.Error(rR)
	}
}

func TestSlidingWindow(t *testing.T) {
	r1 := ListIterator([]int{1, 2, 3, 4, 5})
	rSL := SlidingWindow(r1, 2).List()
	for _, c := range rSL {
		if len(c) != 2 {
			t.Error()
		}
	}
	if rSL[0][0] != 1 {
		t.Error()
	}
	if rSL[3][1] != 5 {
		t.Error()
	}
}

func TestRepeatCtx(t *testing.T) {
	r1 := ListIterator([]int{1, 2})
	rInf := r1.Repeat()
	sum := Sum(rInf.Take(5))
	if 7 != sum {
		t.Error(sum)
	}
}

func TestEnumerate(t *testing.T) {
	r1 := ListIterator([]string{"orange"})
	if 0 != Enumerate(r1).MustTakeOne().Index {
		t.Error()
	}
}

func TestFilter(t *testing.T) {
	r := ListIterator([]int{1, 2, 3, 4})
	rEven := r.Filter(func(i int) bool {
		return i%2 == 0
	})
	evens := rEven.List()
	if len(evens) != 2 {
		t.Error()
	}
	if evens[0] != 2 {
		t.Error()
	}
}

func TestChunk(t *testing.T) {
	r1 := ListIterator([]int{1, 2, 3})
	chunks := Chunk(r1, 2).List()
	if len(chunks) != 2 {
		t.Error()
	}
	if len(chunks[0]) != 2 {
		t.Error()
	}
	if len(chunks[1]) != 1 {
		t.Error()
	}
}

func TestForEach(t *testing.T) {
	r1 := ListIterator([]string{"apple", "orange"})
	receives := make([]string, 0, 2)
	r1.ForEach(func(s string) {
		receives = append(receives, s)
	})
	if len(receives) != 2 {
		t.Error()
	}
	if "apple" != receives[0] {
		t.Error()
	}
}

func TestRange(t *testing.T) {
	l := Range(0, 10).List()
	if len(l) != 10 {
		t.Error()
	}
	if l[0] != 0 {
		t.Error()
	}
	if l[len(l)-1] != 9 {
		t.Error()
	}
}

func TestSum(t *testing.T) {
	r := ListIterator([]int{1, 2, 3})
	if Sum(r) != 6 {
		t.Error()
	}
}
