package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceRemover(t *testing.T) {
	t.Run("Should remove int32", func(t *testing.T) {
		s := []int32{10, 23, 45, 68, 2}
		x := int32(45)
		expectedSlice := []int32{10, 23, 68, 2}

		ret := SliceRemover(s, x)
		assert.EqualValues(t, expectedSlice, ret)
	})
	t.Run("Should remove string", func(t *testing.T) {
		s := []string{"10", "23", "45", "68", "2"}
		x := "45"
		expectedSlice := []string{"10", "23", "68", "2"}

		ret := SliceRemover(s, x)
		assert.EqualValues(t, expectedSlice, ret)
	})
	t.Run("Should remove float64", func(t *testing.T) {
		s := []float64{10.2, 23.5, 45.2, 68.2, 2.5}
		x := 45.2
		expectedSlice := []float64{10.2, 23.5, 68.2, 2.5}

		ret := SliceRemover(s, x)
		assert.EqualValues(t, expectedSlice, ret)
	})
	t.Run("Should remove many values", func(t *testing.T) {
		s := []int{10, 23, 45, 68, 2}
		x := []int{45, 68}
		expectedSlice := []int{10, 23, 2}

		ret := SliceRemover(s, x...)
		assert.EqualValues(t, expectedSlice, ret)
	})
	t.Run("Should not remove because not found", func(t *testing.T) {
		s := []int32{10, 23, 45, 68, 2}
		x := int32(42)
		expectedSlice := []int32{10, 23, 45, 68, 2}

		ret := SliceRemover(s, x)
		assert.EqualValues(t, expectedSlice, ret)
	})
}

func TestSliceFinder(t *testing.T) {
	t.Run("Should return false because list is empty", func(t *testing.T) {
		s := []any{}
		x := any(42)
		expectedIDX := -1
		expectedRet := false

		idx, ret := SliceFinder(s, x)
		if assert.EqualValues(t, expectedRet, ret) {
			assert.EqualValues(t, expectedIDX, idx)
		}
	})
	t.Run("Should return false because not found", func(t *testing.T) {
		s1 := []int32{15, 22, 8, 74, 23}
		x1 := int32(42)
		s2 := []string{"15", "22", "8", "74", "23"}
		x2 := "42"
		expectedIDX := -1

		idx, ret := SliceFinder(s1, x1)
		if assert.False(t, ret) {
			assert.EqualValues(t, expectedIDX, idx)
		}
		idx, ret = SliceFinder(s2, x2)
		if assert.False(t, ret) {
			assert.EqualValues(t, expectedIDX, idx)
		}
	})
	t.Run("Should return true because x finded", func(t *testing.T) {
		s1 := []int32{15, 22, 8, 42, 74, 23}
		x1 := int32(42)
		s2 := []string{"15", "22", "8", "42", "74", "23"}
		x2 := "42"
		expectedIDX := 3

		idx, ret := SliceFinder(s1, x1)
		if assert.True(t, ret) {
			assert.EqualValues(t, expectedIDX, idx)
		}
		idx, ret = SliceFinder(s2, x2)
		if assert.True(t, ret) {
			assert.EqualValues(t, expectedIDX, idx)
		}
	})
}

func TestSliceDuplicateRemover(t *testing.T) {
	t.Run("Should remove duplicated int32", func(t *testing.T) {
		s := []int32{10, 23, 10, 45, 68, 2, 10, 2, 23}
		expectedSlice := []int32{10, 23, 45, 68, 2}

		ret := SliceDuplicateRemover(s)
		assert.EqualValues(t, expectedSlice, ret)
	})
	t.Run("Should remove duplicated string", func(t *testing.T) {
		s := []string{"10", "23", "10", "45", "68", "2", "10", "2", "23"}
		expectedSlice := []string{"10", "23", "45", "68", "2"}

		ret := SliceDuplicateRemover(s)
		assert.EqualValues(t, expectedSlice, ret)
	})
	t.Run("Should remove duplicated float64", func(t *testing.T) {
		s := []float64{10.2, 23.5, 10.2, 45.2, 68.2, 2.5, 10.2, 2.5, 23.5}
		expectedSlice := []float64{10.2, 23.5, 45.2, 68.2, 2.5}

		ret := SliceDuplicateRemover(s)
		assert.EqualValues(t, expectedSlice, ret)
	})
	t.Run("Should not find duplicated", func(t *testing.T) {
		s := []float64{10.2, 23.5, 45.2, 68.2, 2.5}

		ret := SliceDuplicateRemover(s)
		assert.EqualValues(t, s, ret)
	})
}
