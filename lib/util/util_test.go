package util

import (
	"fmt"
	"testing"
)

func TestNumberFields(t *testing.T) {
	var nf []int
	nf = NumberFields("")
	if len(nf) != 0 {
		t.Error()
	}

	nf = NumberFields("941")
	if len(nf) != 1 || nf[0] != 941 {
		fmt.Println(nf)
		t.Error("Failed with simple integer")
	}

	nf = NumberFields("24 -48")
	if len(nf) != 2 || nf[0] != 24 || nf[1] != -48 {
		fmt.Println(nf)
		t.Error("Failed to detect negative number")
	}

	nf = NumberFields("941,230 -> 322,849")
	if len(nf) != 4 || nf[1] != 230 || nf[2] != 322 {
		t.Error("Failed when hyphen included")
	}
}
