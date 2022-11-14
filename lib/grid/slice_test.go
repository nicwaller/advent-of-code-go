package grid

import (
	"fmt"
	"testing"
	"time"
)

func TestSlice1D(t *testing.T) {
	var slice Slice
	slice = SliceEnclosing([]int{0})
	if len(slice) != 1 {
		t.Error("Slice has wrong number of dimensions")
	}
	if slice[0].Origin != 0 {
		t.Error("Slice Origin should be inclusive")
	}
	if slice[0].Terminus != 1 {
		t.Error("Slice Terminus should be inclusive")
	}
}

func TestSlice2D(t *testing.T) {
	var slice Slice
	slice = SliceEnclosing([]int{0, 0})
	if len(slice) != 2 {
		t.Error("Slice has wrong number of dimensions")
	}
	if slice[0].Origin != 0 {
		t.Error("Slice Origin should be inclusive")
	}
	if slice[1].Origin != 0 {
		t.Error("Slice Origin should be inclusive")
	}
	if slice[0].Terminus != 1 {
		t.Error("Slice Terminus should be inclusive")
	}
	if slice[1].Terminus != 1 {
		t.Error("Slice Terminus should be inclusive")
	}
}

func TestSliceNonzero(t *testing.T) {
	var slice Slice
	slice = SliceEnclosing([]int{1, 1})
	if len(slice) != 2 {
		t.Error("Slice has wrong number of dimensions")
	}
	if slice[0].Origin != 1 {
		t.Error("Slice Origin should be inclusive")
	}
	if slice[1].Origin != 1 {
		t.Error("Slice Origin should be inclusive")
	}
	if slice[0].Terminus != 2 {
		t.Error("Slice Terminus should be inclusive")
	}
	if slice[1].Terminus != 2 {
		t.Error("Slice Terminus should be inclusive")
	}
}

func TestSlice_Cells(t *testing.T) {
	slice := SliceEnclosing([]int{1, 1}, []int{1, 1})
	for iter := slice.Cells(); iter.Next(); {
		time.Sleep(200 * time.Millisecond)
	}
}

//func TestLine(t *testing.T) {
//	point := Line([]int{0}, []int{5}).List()
//}

func TestSlice_Intersect(t *testing.T) {
	// 1D
	a := Slice{Range{Origin: 0, Terminus: 2}}
	b := Slice{Range{Origin: 3, Terminus: 2}}
	v := a.Intersect(b)
	fmt.Println(v)
	// TODO: what if there is no intersection?
}
