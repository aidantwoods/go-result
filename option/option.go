package option

import "aidanwoods.dev/go-result/types"

type Option[T any] struct {
	isSome bool
	value  *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{
		isSome: true,
		value:  &value,
	}
}

// None is explicitly identical to the zero value of Option[T]
func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsSome() bool {
	return o.isSome
}

func (o Option[T]) IsNone() bool {
	return !o.IsSome()
}

func (o Option[T]) Expect(panicMsg string) T {
	if o.IsSome() {
		return *o.value
	} else {
		panic(panicMsg)
	}
}

func (o Option[T]) Unwrap() T {
	return o.Expect("value should be present when unwrap is")
}

func (o Option[T]) UnwrapOr(defaultValue T) T {
	return If(o, types.Id[T], types.Return0(defaultValue))
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	return If(o, types.Id[T], fn)
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

func Map[T, U any](r Option[T], fn func(T) U) Option[U] {
	return If(r, types.Compose(fn, Some[U]), None[U])
}

func FlatMap[T, U any](r Option[T], fn func(T) Option[U]) Option[U] {
	return If(Map(r, fn), types.Id[Option[U]], None[U])
}

func Cast[T any](value any) Option[T] {
	if t, ok := value.(T); ok {
		return Some(t)
	} else {
		return None[T]()
	}
}
