package genericclient

func Default[T any]() T {
	return *new(T)
}
