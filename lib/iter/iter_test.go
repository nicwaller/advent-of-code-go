package iter

import (
	"fmt"
	"testing"
)

func TestProductV(t *testing.T) {
	l := []int{1, 2}
	x := ProductV(l, l, l).List()
	fmt.Println(x)
}
