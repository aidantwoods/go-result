package t

type ApplyToResult[Out, T any] struct {
	result Result[T]
}

func Out[Out, T any](result Result[T]) ApplyToResult[Out, T] {
	return ApplyToResult[Out, T]{
		result: result,
	}
}

func (a ApplyToResult[O, T]) AndThen(fn func(T) Result[O]) Result[O] {
	return AndThen(a.result, fn)
}

type ApplyToResult2[O1, O2, T any] struct {
	result Result[T]
}

func Out2[Out2, Out1, T any](result Result[T]) ApplyToResult2[Out1, Out2, T] {
	return ApplyToResult2[Out1, Out2, T]{
		result: result,
	}
}

func (a ApplyToResult2[O1, O2, T]) AndThen(fn func(T) Result[O1]) ApplyToResult[O2, O1] {
	return Out[O2](AndThen(a.result, fn))
}

type ApplyToResult3[O1, O2, O3, T any] struct {
	result Result[T]
}

func Out3[Out3, Out2, Out1, T any](result Result[T]) ApplyToResult3[Out1, Out2, Out3, T] {
	return ApplyToResult3[Out1, Out2, Out3, T]{
		result: result,
	}
}

func (a ApplyToResult3[O1, O2, O3, T]) AndThen(fn func(T) Result[O1]) ApplyToResult2[O3, O2, O1] {
	return Out2[O2, O3](AndThen(a.result, fn))
}

type ApplyToResult4[O1, O2, O3, O4, T any] struct {
	result Result[T]
}

func Out4[Out4, Out3, Out2, Out1, T any](result Result[T]) ApplyToResult4[Out1, Out2, Out3, Out4, T] {
	return ApplyToResult4[Out1, Out2, Out3, Out4, T]{
		result: result,
	}
}

func (a ApplyToResult4[O1, O2, O3, O4, T]) AndThen(fn func(T) Result[O1]) ApplyToResult3[O4, O3, O2, O1] {
	return Out3[O2, O3, O4](AndThen(a.result, fn))
}

type ApplyToResult5[O1, O2, O3, O4, O5, T any] struct {
	result Result[T]
}

func Out5[Out5, Out4, Out3, Out2, Out1, T any](result Result[T]) ApplyToResult5[Out1, Out2, Out3, Out4, Out5, T] {
	return ApplyToResult5[Out1, Out2, Out3, Out4, Out5, T]{
		result: result,
	}
}

func (a ApplyToResult5[O1, O2, O3, O4, O5, T]) AndThen(fn func(T) Result[O1]) ApplyToResult4[O5, O4, O3, O2, O1] {
	return Out4[O2, O3, O4, O5](AndThen(a.result, fn))
}
