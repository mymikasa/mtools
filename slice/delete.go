package slice

import "errors"

func Delete[T any](src []T, index int) ([]T, T, error) {

	length := len(src)
	if index < 0 || index >= length {
		var nilT T
		return nil, nilT, errors.New("下标超出范围")
	}

	res := src[index]

	for i := index; i+1 < length; i++ {
		src[i] = src[i+1]
	}
	src = src[:length-1]
	return src, res, nil
}
