package contains

import "golang.org/x/exp/slices"

// Your implementation of ContainsFunc goes here!
func ContainsFunc[T any](s []T, f func(T) bool) bool {
	return slices.IndexFunc(s, f) >= 0
}
