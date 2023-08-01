package slice

func Shrink[T any](src []T) []T {
	c := cap(src)
	l := len(src)

	if l*2 <= c {
		s := make([]T, l)
		copy(src, s)
		return s
	}
	return src
}
