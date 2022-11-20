package iter

import (
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
