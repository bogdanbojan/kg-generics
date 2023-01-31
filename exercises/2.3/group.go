package group

type Group[E any] []E

// Your Len function goes here!
func Len[E any](vv Group[E]) int {
	return len(vv)
}
