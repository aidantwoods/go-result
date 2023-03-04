package t

type Option[T any] struct {
	value *T
	some  bool
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		value: &value,
		some:  true,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		value: nil,
		some:  false,
	}
}

func Maybe[T any](t *T) Option[T] {
	if t != nil {
		return Some(*t)
	} else {
		return None[T]()
	}
}

func (o Option[T]) IsSome() bool {
	return o.some
}

func (o Option[T]) IsNone() bool {
	return !o.IsSome()
}

func (o Option[T]) AsPtr() *T {
	return o.value
}

func (o Option[T]) Expect(panicMsg string) T {
	if o.IsSome() {
		return *o.value
	} else {
		panic(panicMsg)
	}
}

func (o Option[T]) Unwrap() T {
	return o.Expect("value was not present in option")
}

func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.IsSome() {
		return o.Unwrap()
	} else {
		return defaultValue
	}
}

func (o Option[T]) UnwrapOrElse(defaultValue T) T {
	if o.IsSome() {
		return o.Unwrap()
	} else {
		return defaultValue
	}
}

func (o Option[T]) Some(out *T) bool {
	if o.IsSome() {
		*out = o.Unwrap()
		return true
	} else {
		return false
	}
}
