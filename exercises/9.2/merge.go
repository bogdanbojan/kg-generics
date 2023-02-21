package merge

import "golang.org/x/exp/maps"

// Your implementation of Merge goes here!
func Merge[K comparable, V any](mm ...map[K]V) map[K]V {
	res := map[K]V{}

	for _, m := range mm {
		maps.Copy(res, m)
	}
	return res
}
