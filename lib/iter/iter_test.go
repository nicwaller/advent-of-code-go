package iter

import (
	"fmt"
	"testing"
)

//func TestProductV(t *testing.T) {
//	l := []int{1, 2}
//	x := ProductV(l, l, l).List()
//	fmt.Println(x)
//}

func TestStringIterator(t *testing.T) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	chunks := StringIterator(alphabet, 7).List()
	expected := []string{
		"abcdefg",
		"hijklmn",
		"opqrstu",
		"vwxyz",
	}
	for i, _ := range chunks {
		if chunks[i] != expected[i] {
			t.Errorf("expected %q but got %q", expected[i], chunks[i])
		}
	}
}

func TestPermute(t *testing.T) {
	p := Permute[int]([]int{0, 1, 2}).List()
	if len(p) != 6 {
		t.Errorf("Expected 6 but got %d\n", len(p))
	}
}

func TestSlidingWindow(t *testing.T) {
	alpha := "abcdef"
	expectedPairs := [][]string{
		[]string{"a", "b"},
		[]string{"b", "c"},
		[]string{"c", "d"},
		[]string{"d", "e"},
		[]string{"e", "f"},
	}
	expectedTriples := [][]string{
		[]string{"a", "b", "c"},
		[]string{"b", "c", "d"},
		[]string{"c", "d", "e"},
		[]string{"d", "e", "f"},
	}
	actualPairs := SlidingWindow[string](2, StringIterator(alpha, 1)).List()
	if !slicesAreEqual(expectedPairs, actualPairs) {
		t.Errorf("Incorrect result for pairs in sliding window")
	}

	actualTriples := SlidingWindow[string](3, StringIterator(alpha, 1)).List()
	if !slicesAreEqual(expectedTriples, actualTriples) {
		t.Errorf("Incorrect result for pairs in sliding window")
	}
}

func slicesAreEqual[T comparable](a [][]T, b [][]T) bool {
	if len(a) != len(b) {
		fmt.Println("outer lengths do not match")
		return false
	}
	for i := 0; i < len(a); i++ {
		aa := a[i]
		bb := b[i]
		if len(aa) != len(bb) {
			fmt.Println(aa, bb)
			fmt.Println("inner lengths do not match")
			return false
		}
		for j := 0; j < len(aa); j++ {
			if aa[j] != bb[j] {
				fmt.Println(aa, bb)
				return false
			}
		}
	}
	return true
}
