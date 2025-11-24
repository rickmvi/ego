package function

func And[T any](firstPredicate, secondPredicate Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return firstPredicate(t) && secondPredicate(t)
	}
}

func Or[T any](firstPredicate, secondPredicate Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return firstPredicate(t) || secondPredicate(t)
	}
}

func Not[T any](predicate Predicate[T]) Predicate[T] {
	return func(t T) bool {
		return !predicate(t)
	}
}
