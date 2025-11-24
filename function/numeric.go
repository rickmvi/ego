package function

import "github.com/avila-r/ego/constraint"

func GreaterThen[T constraint.Integer | constraint.Float](value T) Predicate[T] {
	return func(t T) bool {
		return t > value
	}
}

func GreaterThanOrEqual[T constraint.Integer | constraint.Float](t T) Predicate[T] {
	return func(v T) bool { return v >= t }
}

func LessThan[T constraint.Integer | constraint.Float](t T) Predicate[T] {
	return func(v T) bool { return v < t }
}

func LessThanOrEqual[T constraint.Integer | constraint.Float](t T) Predicate[T] {
	return func(v T) bool { return v <= t }
}

func IsNegative[T constraint.Integer | constraint.Float]() Predicate[T] {
	return func(v T) bool { return v < 0 }
}
func IsPositive[T constraint.Integer | constraint.Float]() Predicate[T] {
	return func(v T) bool { return v > 0 }
}
func IsZero[T constraint.Integer | constraint.Float]() Predicate[T] {
	return func(v T) bool { return v == 0 }
}
func IsNotZero[T constraint.Integer | constraint.Float]() Predicate[T] {
	return func(v T) bool { return v != 0 }
}

func Sum[T constraint.Integer | constraint.Float]() BinaryOperator[T] {
	return func(x, y T) T { return x + y }
}
func Subtract[T constraint.Integer | constraint.Float]() BinaryOperator[T] {
	return func(x, y T) T { return x - y }
}
func Multiply[T constraint.Integer | constraint.Float]() BinaryOperator[T] {
	return func(x, y T) T { return x * y }
}
func Divide[T constraint.Integer | constraint.Float]() BinaryOperator[T] {
	return func(x, y T) T { return x / y }
}

func Modulo[T constraint.Integer]() BinaryOperator[T] {
	return func(x, y T) T { return x % y }
}

func Min[T constraint.Integer | constraint.Float]() BinaryOperator[T] {
	return func(x, y T) T {
		if x < y {
			return x
		}
		return y
	}
}

func Max[T constraint.Integer | constraint.Float]() BinaryOperator[T] {
	return func(x, y T) T {
		if x > y {
			return x
		}
		return y
	}
}

func Abs[T constraint.Integer | constraint.Float]() UnaryOperator[T] {
	return func(v T) T {
		if v < 0 {
			return -v
		}
		return v
	}
}
