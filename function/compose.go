package function

func Identity[T any]() UnaryOperator[T] {
	return func(value T) T {
		return value
	}
}

func AndThen[T, R any](fistFunction Function[T, R], secondFunction Function[R, R]) Function[T, R] {
	return func(value T) R {
		return secondFunction(fistFunction(value))
	}
}

func AndThenMap[T, R, Z any](firstFunction Function[T, R], secondFunction Function[R, Z]) Function[T, Z] {
	return func(value T) Z {
		return secondFunction(firstFunction(value))
	}
}

func Compose[T, R any](firstFunction Function[R, R], secondFunction Function[T, R]) Function[T, R] {
	return func(value T) R {
		return firstFunction(secondFunction(value))
	}
}
