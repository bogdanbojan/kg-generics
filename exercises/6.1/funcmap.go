package funcmap

// Your implementation of FuncMap goes here!
type FuncMap[T, U any] map[string]func(T) U

func (f FuncMap[T, U]) Apply(operation string, thing T) U {
	return f[operation](thing)
}
