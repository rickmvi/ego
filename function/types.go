package function

type Predicate[T any] func(T) bool

type BiPredicate[T any] func(T, T) bool

type Consumer[T any] func(T)

type BiConsumer[T any] func(T, T)

type Function[T any, R any] func(T) R

type BiFunction[T any, U any, R any] func(T, U) R

type Supplier[T any] func() T

type BooleanSupplier func() bool

type UnaryOperator[T any] func(T) T

type BinaryOperator[T any] func(T, T) T

type Runnable func()

type Comparator[T any] func(T, T) int
