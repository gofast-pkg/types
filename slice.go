// Package types provides some useful functions for working with native types in Go
package types

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
// Return true if found, false otherwise
func SliceFinder[T comparable](s []T, x T) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}

	return false
}
