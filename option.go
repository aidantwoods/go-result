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
	return If(o, Id[T], Return0(defaultValue))
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	return If(o, Id[T], fn)
}

func (o Option[T]) Some(out *T) bool {
	if o.IsSome() {
		*out = o.Unwrap()
		return true
	} else {
		return false
	}
}

func If[Out, T any](r Option[T], someFn func(T) Out, noneFn func() Out) Out {
	if r.IsSome() {
		return someFn(r.Unwrap())
	} else {
		return noneFn()
	}
}
