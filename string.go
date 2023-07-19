package types

import (
	"strconv"
	"strings"
)

// StringToSliceNumber converts a string to a slice of int.
// The string must be a comma separated list of int.
func StringToSliceNumber(s string, sep string) ([]int, error) {
	var slice []int
	var splitted []string
	var err error

	if s == "" {
		return []int{}, nil
	}
	splitted = strings.Split(s, sep)

	for _, v := range splitted {
		var n int

		if n, err = strconv.Atoi(v); err != nil {
			return []int{}, err
		}
		slice = append(slice, n)
	}

	return slice, nil
}
