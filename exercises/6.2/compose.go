package compose

// Your implementation of Compose goes here!

func Compose[T, V, U any](outer func(T) U, inner func(V) T, value V) U {
	return outer(inner(value))
}
