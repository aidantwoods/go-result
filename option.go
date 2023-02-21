package t

type Option[T any] struct {
	value *T
	some  bool
}

func (r Option[T]) Value() *T {
	return r.value
}

func (r Option[T]) Expect(panicMsg string) T {
	if r.value == nil {
		panic(panicMsg)
	} else {
		return *r.value
	}
}

func (r Option[T]) Unwrap() T {
	return r.Expect("value was not present in result")
}

func (r Option[T]) UnwrapOr(defaultValue T) T {
	if r.value == nil {
		return defaultValue
	} else {
		return *r.value
	}
}

func (r Option[T]) UnwrapOrElse(defaultValue T) T {
	if r.value == nil {
		return defaultValue
	} else {
		return *r.value
	}
}
