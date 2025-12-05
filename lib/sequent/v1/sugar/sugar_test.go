package sugar

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSugar_Args(t *testing.T) {
	list := []int{1, 2, 3}
	sIter := S[int]{}.Args(list...)
	assert.Equal(t, []int{1, 2, 3}, sIter.List())
}

func TestSugar_List(t *testing.T) {
	list := []int{1, 2, 3}
	sIter := S[int]{}.List(list)
	assert.Equal(t, []int{1, 2, 3}, sIter.List())
}

func TestSugar_Seq(t *testing.T) {
	list := []int{1, 2, 3}
	goIter := slices.Values(list)
	sIter := S[int]{}.Seq(goIter)
	assert.Equal(t, []int{1, 2, 3}, sIter.List())
}
