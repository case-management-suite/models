package models

type Optional[T any] struct {
	value     *T
	isPresent bool
	isEmpty   bool
}

func OptionalEmpty[T any]() Optional[T] {
	return Optional[T]{isPresent: false, isEmpty: true}
}

func OptionalOf[T any](t T) Optional[T] {
	return Optional[T]{value: &t, isPresent: true, isEmpty: false}
}

func (o Optional[T]) IsEmpty() bool {
	return o.isEmpty
}

func (o Optional[T]) IsPresent() bool {
	return o.isPresent
}

func (o Optional[T]) Get() *T {
	return o.value
}

func (o Optional[T]) GetOrDefault(t T) *T {
	if o.isPresent {
		return o.value
	}
	return &t
}
