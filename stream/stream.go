package stream

import (
	"github.com/avila-r/ego/function"
	"github.com/avila-r/ego/optional"
	"sort"
)

type Stream[T comparable] struct {
	elements []T
}

func Of[T comparable](elements ...T) Stream[T] {
	return Stream[T]{elements: elements}
}

func From[T comparable](collectable Collectable[T]) Stream[T] {
	return Stream[T]{elements: collectable.Elements()}
}

func (s Stream[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Stream[T]) ToList() []T {
	return s.elements
}

//func (s Stream[T]) ToSet() map[T]struct{} {
//	set := make(map[T]struct{}, len(s.elements))
//	for _, e := range s.elements {
//		set[e] = struct{}{}
//	}
//	return set
//}

func (s Stream[T]) Count() int {
	return len(s.elements)
}

func (s Stream[T]) Limit(number int) Stream[T] {
	if number >= len(s.elements) {
		return s
	}
	return Stream[T]{elements: s.elements[:number]}
}

func (s Stream[T]) Skip(number int) Stream[T] {
	if number >= len(s.elements) {
		return Stream[T]{elements: []T{}}
	}
	return Stream[T]{elements: s.elements[number:]}
}

func (s Stream[T]) Concat(other Stream[T]) Stream[T] {
	newElements := make([]T, 0, len(s.elements)+len(other.elements))
	newElements = append(newElements, s.elements...)
	newElements = append(newElements, other.elements...)
	return Stream[T]{elements: newElements}
}

func (s Stream[T]) TakeWhile(predicate function.Predicate[T]) Stream[T] {
	newElements := make([]T, 0)
	for _, element := range s.elements {
		if !predicate(element) {
			break
		}
		newElements = append(newElements, element)
	}
	return Stream[T]{elements: newElements}
}

func (s Stream[T]) DropWhile(predicate function.Predicate[T]) Stream[T] {
	newElements := make([]T, 0)
	dropping := true
	for _, element := range s.elements {
		if dropping && predicate(element) {
			continue
		}
		dropping = false
		newElements = append(newElements, element)
	}
	return Stream[T]{elements: newElements}
}

func (s Stream[T]) FindFirst() (optional.Optional[T], bool) {
	if s.IsEmpty() {
		var zero optional.Optional[T]
		return zero, false
	}

	return optional.Of(s.elements[0]), true
}

func (s Stream[T]) FindAny() (optional.Optional[T], bool) {
	if s.IsEmpty() {
		var empty optional.Optional[T]
		return empty, false
	}
	return optional.Of(s.elements[0]), true
}

func (s Stream[T]) Min(compare function.Comparator[T]) (optional.Optional[T], bool) {
	if s.IsEmpty() {
		var zero optional.Optional[T]
		return zero, false
	}

	m := s.elements[0]
	for _, element := range s.elements[1:] {
		if compare(element, m) < 0 {
			m = element
		}
	}
	return optional.Of(m), true
}

func (s Stream[T]) Max(compare function.Comparator[T]) (optional.Optional[T], bool) {
	if s.IsEmpty() {
		var zero optional.Optional[T]
		return zero, false
	}

	m := s.elements[0]
	for _, element := range s.elements[1:] {
		if compare(element, m) > 0 {
			m = element
		}
	}
	return optional.Of(m), true
}

func (s Stream[T]) Sorted(less function.BiPredicate[T]) Stream[T] {
	sort.Slice(s.elements, func(i, j int) bool {
		return less(s.elements[i], s.elements[j])
	})
	return Stream[T]{elements: s.elements}
}

func (s Stream[T]) Peek(consumer function.Consumer[T]) Stream[T] {
	for _, element := range s.elements {
		consumer(element)
	}
	return s
}

func (s Stream[T]) Reduce(reducer function.BinaryOperator[T]) (optional.Optional[T], bool) {
	if s.IsEmpty() {
		var zero optional.Optional[T]
		return zero, false
	}

	result := s.elements[0]
	for _, element := range s.elements[1:] {
		result = reducer(result, element)
	}
	return optional.Of(result), true
}

func (s Stream[T]) ReduceWithIdentity(identity T, reducer function.BinaryOperator[T]) T {
	result := identity
	for _, element := range s.elements {
		result = reducer(result, element)
	}
	return result
}

func (s Stream[T]) AllMatch(predicate function.Predicate[T]) bool {
	for _, element := range s.elements {
		if !predicate(element) {
			return false
		}
	}
	return true
}

func (s Stream[T]) AnyMatch(predicate function.Predicate[T]) bool {
	for _, element := range s.elements {
		if predicate(element) {
			return true
		}
	}
	return false
}

func (s Stream[T]) NoneMatch(predicate function.Predicate[T]) bool {
	for _, element := range s.elements {
		if predicate(element) {
			return false
		}
	}
	return true
}

func (s Stream[T]) ForEach(consumer function.Consumer[T]) {
	for _, element := range s.elements {
		consumer(element)
	}
}

func (s Stream[T]) Filter(predicate function.Predicate[T]) Stream[T] {
	newElements := make([]T, 0)
	for _, element := range s.elements {
		if predicate(element) {
			newElements = append(newElements, element)
		}
	}

	return Stream[T]{elements: newElements}
}

// TODO: I need help here
//func (s Stream[T]) Map[R any](mapper function.Function[T, R]) Stream[R] {
//	newElements := make([]R, 0, len(s.elements))
//	for _, element := range s.elements {
//		newElements = append(newElements, mapper(element))
//	}
//	return Stream[R]{elements: newElements}
//}
//
//func (s Stream[T]) FlatMap[R comparable](mapper function.Function[T, []R]) Stream[R] {
//	newElements := make([]R, 0)
//	for _, element := range s.elements {
//		mapped := mapper(element)
//		newElements = append(newElements, mapped...)
//	}
//	return Stream[R]{elements: newElements}
//}

func (s Stream[T]) Distinct() Stream[T] {
	seen := make(map[T]bool)
	newElements := make([]T, 0)

	for _, element := range s.elements {
		if !seen[element] {
			seen[element] = true
			newElements = append(newElements, element)
		}
	}
	return Stream[T]{elements: newElements}
}
