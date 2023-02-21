package t

import "fmt"

type Result[T any] struct {
	ResultE[T, error]
}

func (r Result[T]) Results() (*T, error) {
	r.assertState()

	return r.value, r.UnwrapErrOr(nil)
}

func (r Result[T]) UnwrappedResults() (T, error) {
	r.assertState()

	return *r.value, r.UnwrapErrOr(nil)
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		ResultE[T, error]{
			value: &value,
			err:   nil,
		},
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		ResultE[T, error]{
			value: nil,
			err:   &err,
		},
	}
}

func ErrMap[T any, E error](fn func(error) E) func(e error) Result[T] {
	return func(e error) Result[T] {
		var err error
		err = fn(e)
		return Result[T]{
			ResultE[T, error]{
				value: nil,
				err:   &err,
			},
		}
	}
}

func Wrap[T any](format string, err error) Result[T] {
	wrapped := fmt.Errorf(format, err)
	return Result[T]{
		ResultE[T, error]{
			value: nil,
			err:   &wrapped,
		},
	}
}

func NewResult[T any](value T, err error) Result[T] {
	result := Result[T]{
		ResultE[T, error]{
			value: &value,
			err:   &err,
		},
	}
	result.assertState()

	return result
}

func NewPtrResult[T any](value *T, err error) Result[T] {
	result := Result[T]{
		ResultE[T, error]{
			value: value,
			err:   &err,
		},
	}
	result.assertState()

	return result
}

func Map[T, U any](r Result[T], fn func(T) U) Result[U] {
	r.assertState()

	if r.IsErr() {
		return Err[U](*r.err)
	} else {
		return Ok(fn(*r.value))
	}
}

func MapErr[T any, E error](fn func(error) E, r Result[T]) Result[T] {
	r.assertState()

	if r.IsOk() {
		return Ok(*r.value)
	} else {
		return Err[T](fn(*r.err))
	}
}

func AndThen[T, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	r.assertState()

	if r.IsErr() {
		return Err[U](*r.err)
	} else {
		return fn(*r.value)
	}
}

func AndThenWith[T, U, V any](r Result[T], s Result[U], fn func(T, U) Result[V]) Result[V] {
	r.assertState()

	if r.IsErr() {
		return Err[V](*r.err)
	} else if s.IsErr() {
		return Err[V](*s.err)
	} else {
		return fn(*r.value, *s.value)
	}
}

func Map2[T, U, V any](r Result[T], s Result[U], fn func(T, U) V) Result[V] {
	r.assertState()

	if r.IsErr() {
		return Err[V](*r.err)
	} else if s.IsErr() {
		return Err[V](*s.err)
	} else {
		return Ok(fn(*r.value, *s.value))
	}
}

func Map3[T, U, V, W any](r Result[T], s Result[U], t Result[V], fn func(T, U, V) W) Result[W] {
	r.assertState()

	if r.IsErr() {
		return Err[W](*r.err)
	} else if s.IsErr() {
		return Err[W](*s.err)
	} else if t.IsErr() {
		return Err[W](*t.err)
	} else {
		return Ok(fn(*r.value, *s.value, *t.value))
	}
}

func Match[Out, T any](r Result[T], okFn func(T) Out, errFn func(error) Out) Out {
	if r.IsOk() {
		return okFn(*r.value)
	} else {
		return errFn(*r.err)
	}
}
