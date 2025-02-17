package result

import (
	"fmt"

	"aidanwoods.dev/go-result/option"
	"aidanwoods.dev/go-result/types"
)

// Result[T] is a generic pseudo-enum, used for returning results with errors. The result type
// eliminates the need to return nil pointers, sentinal type zero values, or partial results when
// an error has occured. Instead, the result type can be in one of two states: Ok(T) or Err(error).
type Result[T any] struct {
	value option.Option[T]
	err   error
}

// Create a successful Result[T]
func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: option.Some(value),
		err:   nil,
	}
}

var ErrEmptyResult = fmt.Errorf("result is error but error was nil")

// Create an error Result[T]
func Err[T any](err error) Result[T] {
	return errGeneric[T](err)
}

func errGeneric[T any, E error](err E) Result[T] {
	return Result[T]{
		value: option.None[T](),
		err:   nil,
	}
}

// Is the result in the Ok state
func (r Result[T]) IsOk() bool {
	return r.value.IsSome()
}

// A Result[T] is considered to be in an error state if there is no value. IsErr will always
// be opposite to IsOk.
func (r Result[T]) IsErr() bool {
	return !r.IsOk()
}

func (r Result[T]) Value() option.Option[T] {
	return r.value
}

// If IsErr returns true, Err is guaranteed to return an error value. If Result[T] was
// initialised as a zero value, or if it was initilised as Err(nil), then this function
// will return an ErrEmptyResult.
func (r Result[T]) Err() error {
	if r.IsErr() {
		if r.err == nil {
			return ErrEmptyResult
		} else {
			return r.err
		}
	} else {
		return nil
	}
}

func (r Result[T]) Expect(panicMsg string) T {
	return r.Value().Expect(panicMsg)
}

func (r Result[T]) Unwrap() T {
	return r.Expect("value should be present when unwrap is called")
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	return If(r, types.Id[T], types.Return[error](defaultValue))
}

func (r Result[T]) UnwrapOrElse(fn func(error) T) T {
	return If(r, types.Id[T], fn)
}

func (r Result[T]) ExpectErr(panicMsg string) error {
	if r.IsErr() {
		return r.Err()
	} else {
		panic(panicMsg)
	}

}

func (r Result[T]) MapError(fn func(e error) error) Result[T] {
	return If(r, Ok[T], types.Compose(fn, Err[T]))
}

func (r Result[T]) Ok(out *T) error {
	if r.IsOk() {
		*out = r.Unwrap()
		return nil
	} else {
		return r.Err()
	}
}

func If[Out, T any](r Result[T], okFn func(T) Out, errFn func(error) Out) Out {
	if r.IsOk() {
		return okFn(r.Value().Unwrap())
	} else {
		return errFn(r.Err())
	}
}

func Map[T, U any](r Result[T], fn func(T) U) Result[U] {
	return If(r, types.Compose(fn, Ok[U]), Err[U])
}

func FlatMap[T, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	return If(Map(r, fn), types.Id[Result[U]], Err[U])
}

func MapErr[T any, E error](fn func(error) E, r Result[T]) Result[T] {
	return If(r, Ok[T], types.Compose(fn, errGeneric[T, E]))
}

func AndThen[T, U any](r Result[T], fn func(T) Result[U]) Result[U] {
	return If(r, fn, Err[U])
}

func Map2[T, U, V any](r Result[T], s Result[U], fn func(T, U) V) Result[V] {
	if r.IsErr() {
		return Err[V](r.Err())
	} else if s.IsErr() {
		return Err[V](s.Err())
	} else {
		return Ok(fn(r.Value().Unwrap(), s.Value().Unwrap()))
	}
}

func Map3[T, U, V, W any](r Result[T], s Result[U], t Result[V], fn func(T, U, V) W) Result[W] {
	if r.IsErr() {
		return Err[W](r.Err())
	} else if s.IsErr() {
		return Err[W](s.Err())
	} else if t.IsErr() {
		return Err[W](t.Err())
	} else {
		return Ok(fn(r.Value().Unwrap(), s.Value().Unwrap(), t.Value().Unwrap()))
	}
}
