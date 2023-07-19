package types

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToSliceNumber(t *testing.T) {
	t.Run("Should return empty slice because empty string", func(t *testing.T) {
		s := ""
		sep := ","
		expectedSlice := []int{}

		ret, err := StringToSliceNumber(s, sep)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedSlice, ret)
		}
	})
	t.Run("Should return an error because string is not a splitable", func(t *testing.T) {
		s := "10,23,45,68,2"
		sep := "a"

		ret, err := StringToSliceNumber(s, sep)
		if assert.Error(t, err) {
			assert.Empty(t, ret)
			assert.ErrorIs(t, err, strconv.ErrSyntax)
		}
	})
	t.Run("Should return a slice of int", func(t *testing.T) {
		s := "10,23,45,68,2"
		sep := ","
		expectedSlice := []int{10, 23, 45, 68, 2}

		ret, err := StringToSliceNumber(s, sep)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedSlice, ret)
		}
	})
}
