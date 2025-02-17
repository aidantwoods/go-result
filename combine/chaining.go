package combine

import "aidanwoods.dev/go-result/result"

type ApplyToResult[Out, T any] struct {
	result result.Result[T]
}

func Chain[Out, T any](result result.Result[T]) ApplyToResult[Out, T] {
	return ApplyToResult[Out, T]{
		result: result,
	}
}

func (a ApplyToResult[O, T]) AndThen(fn func(T) result.Result[O]) result.Result[O] {
	return result.AndThen(a.result, fn)
}

func (a ApplyToResult[O, T]) Map(fn func(T) O) result.Result[O] {
	return result.Map(a.result, fn)
}

type ApplyToResult2[O1, O2, T any] struct {
	result result.Result[T]
}

func Chain2[Out2, Out1, T any](result result.Result[T]) ApplyToResult2[Out1, Out2, T] {
	return ApplyToResult2[Out1, Out2, T]{
		result: result,
	}
}

func (a ApplyToResult2[O1, O2, T]) AndThen(fn func(T) result.Result[O1]) ApplyToResult[O2, O1] {
	return Chain[O2](result.AndThen(a.result, fn))
}

func (a ApplyToResult2[O1, O2, T]) Map(fn func(T) O1) ApplyToResult[O2, O1] {
	return Chain[O2](result.Map(a.result, fn))
}

type ApplyToResult3[O1, O2, O3, T any] struct {
	result result.Result[T]
}

func Chain3[Out3, Out2, Out1, T any](result result.Result[T]) ApplyToResult3[Out1, Out2, Out3, T] {
	return ApplyToResult3[Out1, Out2, Out3, T]{
		result: result,
	}
}

func (a ApplyToResult3[O1, O2, O3, T]) AndThen(fn func(T) result.Result[O1]) ApplyToResult2[O3, O2, O1] {
	return Chain2[O2, O3](result.AndThen(a.result, fn))
}

func (a ApplyToResult3[O1, O2, O3, T]) Map(fn func(T) O1) ApplyToResult2[O3, O2, O1] {
	return Chain2[O2, O3](result.Map(a.result, fn))
}

type ApplyToResult4[O1, O2, O3, O4, T any] struct {
	result result.Result[T]
}

func Chain4[Out4, Out3, Out2, Out1, T any](result result.Result[T]) ApplyToResult4[Out1, Out2, Out3, Out4, T] {
	return ApplyToResult4[Out1, Out2, Out3, Out4, T]{
		result: result,
	}
}

func (a ApplyToResult4[O1, O2, O3, O4, T]) AndThen(fn func(T) result.Result[O1]) ApplyToResult3[O4, O3, O2, O1] {
	return Chain3[O2, O3, O4](result.AndThen(a.result, fn))
}

func (a ApplyToResult4[O1, O2, O3, O4, T]) Map(fn func(T) O1) ApplyToResult3[O4, O3, O2, O1] {
	return Chain3[O2, O3, O4](result.Map(a.result, fn))
}

type ApplyToResult5[O1, O2, O3, O4, O5, T any] struct {
	result result.Result[T]
}

func Chain5[Out5, Out4, Out3, Out2, Out1, T any](result result.Result[T]) ApplyToResult5[Out1, Out2, Out3, Out4, Out5, T] {
	return ApplyToResult5[Out1, Out2, Out3, Out4, Out5, T]{
		result: result,
	}
}

func (a ApplyToResult5[O1, O2, O3, O4, O5, T]) AndThen(fn func(T) result.Result[O1]) ApplyToResult4[O5, O4, O3, O2, O1] {
	return Chain4[O2, O3, O4, O5](result.AndThen(a.result, fn))
}

func (a ApplyToResult5[O1, O2, O3, O4, O5, T]) Map(fn func(T) O1) ApplyToResult4[O5, O4, O3, O2, O1] {
	return Chain4[O2, O3, O4, O5](result.Map(a.result, fn))
}
