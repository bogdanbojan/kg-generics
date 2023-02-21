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

// Alternative proposed solution:
// func Merge[M ~map[K]V, K comparable, V any](ms ...M) M {
// 	result := M{}
// 	for _, m := range ms {
// 		maps.Copy[map[K]V](result, m)
// 	}
// 	return result
// }