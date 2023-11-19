package goZip

import (
	"fmt"
)

type Pair[A, B any] struct {
	First  A
	Second B
}

func Zip[A, B any](a []A, b []B) ([]Pair[A, B], error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("Length of slices are not equal")
	}
	pairs := make([]Pair[A, B], len(a))
	for i := range a {
		pairs[i] = Pair[A, B]{a[i], b[i]}
	}
	return pairs, nil

}
