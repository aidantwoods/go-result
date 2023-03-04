package t

import "fmt"

// Result[T] is a generic pseudo-enum, used for returning results with errors. The result type
// eliminates the need to return nil pointers, sentinal type zero values, or partial results when
// an error has occured. Instead, the result type can be in one of two states: Ok(T) or Err(error).
type Result[T any] struct {
	value Option[T]
	err   Option[error]
}

// Create a successful Result[T]
func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: Some(value),
		err:   None[error](),
	}
}

// Create an error Result[T]
func Err[T any](err error) Result[T] {
	return errGeneric[T](err)
}

func errGeneric[T any, E error](err E) Result[T] {
	return Result[T]{
		value: None[T](),
		err:   Some[error](err),
	}
}

// func Wrap[T any](format string, err error) Result[T] {
// 	return Result[T]{
// 		value: None[T](),
// 		err:   Some(fmt.Errorf(format, err)),
// 	}
// }

// Is the result in the Ok state
func (r Result[T]) IsOk() bool {
	return !r.IsErr()
}

// Is the result in the Err state
func (r Result[T]) IsErr() bool {
	return r.err.IsSome()
}

func (r Result[T]) Value() Option[T] {
	return r.value
}

func (r Result[T]) Err() Option[error] {
	return r.err
}

func (r Result[T]) Expect(panicMsg string) T {
	return r.Value().Expect(panicMsg)
}

func (r Result[T]) Unwrap() T {
	return r.Expect("value was not present in result")
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	return Match(r, Id[T], Return[error](defaultValue))
}

func (r Result[T]) UnwrapOrElse(fn func(error) T) T {
	return Match(r, Id[T], fn)
}

func (r Result[T]) ExpectErr(panicMsg string) error {
	return r.Err().Expect(panicMsg)
}

func (r Result[T]) UnwrapErr() error {
	return r.ExpectErr("err was not present in result")
}

func (r Result[T]) UnwrapErrOr(defaultValue error) error {
	return Match(r, Return[T](defaultValue), Id[error])
}

func (r Result[T]) WrapErr(format string) Result[T] {
	return r.MapError(func(err error) error { return fmt.Errorf(format, err) })
}

func (r Result[T]) MapError(fn func(e error) error) Result[T] {
	return Match(r, Ok[T], Compose(fn, Err[T]))
}

func (r Result[T]) Ok(out *T) error {
	if r.IsOk() {
		*out = r.Unwrap()
		return nil
	} else {
		return r.UnwrapErr()
	}
}

func Match[Out, T any](r Result[T], okFn func(T) Out, errFn func(error) Out) Out {
	if r.IsOk() {
		return okFn(r.Value().Unwrap())
	} else {
		return errFn(r.Err().Unwrap())
	}
}

func Map[T, U any](r Result[T], fn func(T) U) Result[U] {
	return Match(r, Compose(fn, Ok[U]), Err[U])
}

func MapErr[T any, E error](fn func(error) E, r Result[T]) Result[T] {
	return Match(r, Ok[T], Compose(fn, errGeneric[T, E]))
}

func AndThen[T, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	return Match(r, fn, Err[U])
}

func Map2[T, U, V any](r Result[T], s Result[U], fn func(T, U) V) Result[V] {
	if r.IsErr() {
		return Err[V](r.Err().Unwrap())
	} else if s.IsErr() {
		return Err[V](s.Err().Unwrap())
	} else {
		return Ok(fn(r.Value().Unwrap(), s.Value().Unwrap()))
	}
}

func Map3[T, U, V, W any](r Result[T], s Result[U], t Result[V], fn func(T, U, V) W) Result[W] {
	if r.IsErr() {
		return Err[W](r.Err().Unwrap())
	} else if s.IsErr() {
		return Err[W](s.Err().Unwrap())
	} else if t.IsErr() {
		return Err[W](t.Err().Unwrap())
	} else {
		return Ok(fn(r.Value().Unwrap(), s.Value().Unwrap(), t.Value().Unwrap()))
	}
}

func maybeErrorToOption(err error) Option[error] {
	if err != nil {
		return Some(err)
	} else {
		return None[error]()
	}
}

func NewResult[T any](value T, err error) Result[T] {
	result := Result[T]{
		value: Some(value),
		err:   maybeErrorToOption(err),
	}

	return result
}

func NewPtrResult[T any](value *T, err error) Result[T] {
	result := Result[T]{
		value: Maybe(value),
		err:   maybeErrorToOption(err),
	}

	return result
}

func (r Result[T]) Results() (*T, error) {
	return r.Value().AsPtr(), r.UnwrapErrOr(nil)
}

func (r Result[T]) UnwrappedResults() (T, error) {
	return r.Value().Unwrap(), r.UnwrapErrOr(nil)
}
