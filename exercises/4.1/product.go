package product

import "golang.org/x/exp/constraints"

// Your constraint definition goes here!
type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Your Product function goes here!
func Product[T Number](x, y T) T {
	return x * y
}
