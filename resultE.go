package t

import "reflect"

type ResultE[T any, E error] struct {
	value *T
	err   *E
}

func OkE[T any, E error](value T) ResultE[T, E] {
	return ResultE[T, E]{
		value: &value,
		err:   nil,
	}
}

func ErrE[T any, E error](err E) ResultE[T, E] {
	return ResultE[T, E]{
		value: nil,
		err:   &err,
	}
}

func NewResultE[T any, E error](value T, err E) ResultE[T, E] {
	result := ResultE[T, E]{
		value: &value,
		err:   &err,
	}
	result.assertState()

	return result
}

func (r ResultE[T, E]) validateState() bool {
	// we allow both to be non nil to allow functions that return data and an error to map cleanly
	// to this result type. If an error is present, the result is considered to be in an error state.
	return r.value != nil || !reflect.ValueOf(r.err).IsNil()
}

func (r ResultE[T, E]) assertState() {
	if !r.validateState() {
		panic("bad internal state")
	}
}

func (r ResultE[T, E]) IsOk() bool {
	r.assertState()

	return !r.IsErr()
}

func (r ResultE[T, E]) IsErr() bool {
	r.assertState()

	if r.err == nil {
		return false
	} else if reflect.TypeOf(r.err).Kind() == reflect.Ptr {
		return !reflect.Indirect(reflect.ValueOf(r.err)).IsNil()
	}

	return true
}

func (r ResultE[T, E]) Value() *T {
	r.assertState()

	return r.value
}

func (r ResultE[T, E]) Err() *E {
	r.assertState()

	return r.err
}

func (r ResultE[T, E]) Expect(panicMsg string) T {
	r.assertState()

	if r.IsErr() {
		panic(panicMsg)
	} else {
		return *r.Value()
	}
}

func (r ResultE[T, E]) Unwrap() T {
	return r.Expect("value was not present in result")
}

func (r ResultE[T, E]) UnwrapOr(defaultValue T) T {
	r.assertState()

	if r.IsErr() {
		return defaultValue
	} else {
		return r.Unwrap()
	}
}

func (r ResultE[T, E]) UnwrapOrElse(fn func(E) T) T {
	r.assertState()

	if r.IsErr() {
		return fn(r.UnwrapErr())
	} else {
		return r.Unwrap()
	}
}

func (r ResultE[T, E]) ExpectErr(panicMsg string) E {
	r.assertState()

	if r.IsErr() {
		return *r.Err()
	} else {
		panic(panicMsg)
	}
}

func (r ResultE[T, E]) UnwrapErr() E {
	return r.ExpectErr("err was not present in result")
}

func (r ResultE[T, E]) UnwrapErrOr(defaultValue E) E {
	r.assertState()

	if r.IsOk() {
		return defaultValue
	} else {
		return r.UnwrapErr()
	}
}

func (r ResultE[T, E]) WrapErr(format string) Result[T] {
	r.assertState()

	if r.IsOk() {
		return Ok(r.Unwrap())
	} else {
		return Wrap[T](format, r.UnwrapErr())
	}
}

func (r ResultE[T, E]) MapError(fn func(e E) error) Result[T] {
	r.assertState()

	if r.IsOk() {
		return Ok(r.Unwrap())
	} else {
		return Err[T](fn(r.UnwrapErr()))
	}
}

func MapE[T, U any, E error](r ResultE[T, E], fn func(T) U) ResultE[U, E] {
	r.assertState()

	if r.IsErr() {
		return ErrE[U](r.UnwrapErr())
	} else {
		return OkE[U, E](fn(r.Unwrap()))
	}
}

func MapEErr[T any, E, E2 error](r ResultE[T, E], fn func(E) E2) ResultE[T, E2] {
	r.assertState()

	if r.IsOk() {
		return OkE[T, E2](r.Unwrap())
	} else {
		return ErrE[T](fn(r.UnwrapErr()))
	}
}

func ResultEThen[T, U any, E error](r ResultE[T, E], fn func(T) ResultE[U, E]) ResultE[U, E] {
	r.assertState()

	if r.IsErr() {
		return ErrE[U](r.UnwrapErr())
	} else {
		return fn(r.Unwrap())
	}
}
