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
	if slice[0].origin != 0 {
		t.Error("Slice origin should be inclusive")
	}
	if slice[0].terminus != 1 {
		t.Error("Slice terminus should be inclusive")
	}
}

func TestSlice2D(t *testing.T) {
	var slice Slice
	slice = SliceEnclosing([]int{0, 0})
	if len(slice) != 2 {
		t.Error("Slice has wrong number of dimensions")
	}
	if slice[0].origin != 0 {
		t.Error("Slice origin should be inclusive")
	}
	if slice[1].origin != 0 {
		t.Error("Slice origin should be inclusive")
	}
	if slice[0].terminus != 1 {
		t.Error("Slice terminus should be inclusive")
	}
	if slice[1].terminus != 1 {
		t.Error("Slice terminus should be inclusive")
	}
}

func TestSliceNonzero(t *testing.T) {
	var slice Slice
	slice = SliceEnclosing([]int{1, 1})
	if len(slice) != 2 {
		t.Error("Slice has wrong number of dimensions")
	}
	if slice[0].origin != 1 {
		t.Error("Slice origin should be inclusive")
	}
	if slice[1].origin != 1 {
		t.Error("Slice origin should be inclusive")
	}
	if slice[0].terminus != 2 {
		t.Error("Slice terminus should be inclusive")
	}
	if slice[1].terminus != 2 {
		t.Error("Slice terminus should be inclusive")
	}
}

func TestSlice_Cells(t *testing.T) {
	slice := SliceEnclosing([]int{1, 1}, []int{1, 1})
	for iter := slice.Cells(); iter.Next(); {
		time.Sleep(200 * time.Millisecond)
	}
}

func TestLine(t *testing.T) {
	point := Line([]int{0}, []int{5}).List()
	fmt.Println(point)
}
