package function

import (
	"github.com/avila-r/ego/constraint"
	"reflect"
)

func Equals[T constraint.Comparable](value T) Predicate[T] {
	return func(v T) bool {
		return v == value
	}
}

func DeepEquals[T any](comparing, comparator T) bool {
	return reflect.DeepEqual(comparing, comparator)
}
