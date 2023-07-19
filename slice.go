// Package types provides some useful functions for working with native types in Go
package types

// SliceRemover removes many element(s) from a slice
// Return a new slice without the element(s)
func SliceRemover[T comparable](s []T, x ...T) []T {
	for _, v := range x {
		idx, found := SliceFinder(s, v)
		if !found {
			continue
		}

		s = append(s[:idx], s[idx+1:]...)
	}

	return s
}

// SliceDuplicateRemover removes duplicated elements from a slice
// Return a new slice without duplicated elements
func SliceDuplicateRemover[T comparable](s []T) []T {
	counts := make(map[T]bool)
	ret := []T{}

	for _, v := range s {
		if !counts[v] {
			counts[v] = true
			ret = append(ret, v)
		}
	}

	return ret
}

// SliceFinder search for an element in a slice
// If found, return the index and true.
// Return false otherwise
func SliceFinder[T comparable](s []T, x T) (int, bool) {
	for index, v := range s {
		if v == x {
			return index, true
		}
	}

	return -1, false
}
