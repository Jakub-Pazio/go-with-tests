package functional

func foldl[T any, R any](arr []T, f func(T, R) R, acc R) R {
	if len(arr) < 1 {
		return acc
	}
	return foldl[T, R](arr[1:], f, f(arr[0], acc))
}
