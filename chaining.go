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

func (a ApplyToResult[O, T]) Map(fn func(T) O) Result[O] {
	return Map(a.result, fn)
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

func (a ApplyToResult2[O1, O2, T]) Map(fn func(T) O1) ApplyToResult[O2, O1] {
	return Out[O2](Map(a.result, fn))
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

func (a ApplyToResult3[O1, O2, O3, T]) Map(fn func(T) O1) ApplyToResult2[O3, O2, O1] {
	return Out2[O2, O3](Map(a.result, fn))
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

func (a ApplyToResult4[O1, O2, O3, O4, T]) Map(fn func(T) O1) ApplyToResult3[O4, O3, O2, O1] {
	return Out3[O2, O3, O4](Map(a.result, fn))
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

func (a ApplyToResult5[O1, O2, O3, O4, O5, T]) Map(fn func(T) O1) ApplyToResult4[O5, O4, O3, O2, O1] {
	return Out4[O2, O3, O4, O5](Map(a.result, fn))
}

type ApplyToResult6[O1, O2, O3, O4, O5, O6, T any] struct {
	result Result[T]
}

func Out6[Out6, Out5, Out4, Out3, Out2, Out1, T any](result Result[T]) ApplyToResult6[Out1, Out2, Out3, Out4, Out5, Out6, T] {
	return ApplyToResult6[Out1, Out2, Out3, Out4, Out5, Out6, T]{
		result: result,
	}
}

func (a ApplyToResult6[O1, O2, O3, O4, O5, O6, T]) AndThen(fn func(T) Result[O1]) ApplyToResult5[O6, O5, O4, O3, O2, O1] {
	return Out5[O2, O3, O4, O5, O6](AndThen(a.result, fn))
}

func (a ApplyToResult6[O1, O2, O3, O4, O5, O6, T]) Map(fn func(T) O1) ApplyToResult5[O6, O5, O4, O3, O2, O1] {
	return Out5[O2, O3, O4, O5, O6](Map(a.result, fn))
}

type ApplyToResult7[O1, O2, O3, O4, O5, O6, O7, T any] struct {
	result Result[T]
}

func Out7[Out7, Out6, Out5, Out4, Out3, Out2, Out1, T any](result Result[T]) ApplyToResult7[Out1, Out2, Out3, Out4, Out5, Out6, Out7, T] {
	return ApplyToResult7[Out1, Out2, Out3, Out4, Out5, Out6, Out7, T]{
		result: result,
	}
}

func (a ApplyToResult7[O1, O2, O3, O4, O5, O6, O7, T]) AndThen(fn func(T) Result[O1]) ApplyToResult6[O7, O6, O5, O4, O3, O2, O1] {
	return Out6[O2, O3, O4, O5, O6, O7](AndThen(a.result, fn))
}

func (a ApplyToResult7[O1, O2, O3, O4, O5, O6, O7, T]) Map(fn func(T) O1) ApplyToResult6[O7, O6, O5, O4, O3, O2, O1] {
	return Out6[O2, O3, O4, O5, O6, O7](Map(a.result, fn))
}

type ApplyToResult8[O1, O2, O3, O4, O5, O6, O7, O8, T any] struct {
	result Result[T]
}

func Out8[Out8, Out7, Out6, Out5, Out4, Out3, Out2, Out1, T any](result Result[T]) ApplyToResult8[Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, T] {
	return ApplyToResult8[Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, T]{
		result: result,
	}
}

func (a ApplyToResult8[O1, O2, O3, O4, O5, O6, O7, O8, T]) AndThen(fn func(T) Result[O1]) ApplyToResult7[O8, O7, O6, O5, O4, O3, O2, O1] {
	return Out7[O2, O3, O4, O5, O6, O7, O8](AndThen(a.result, fn))
}

func (a ApplyToResult8[O1, O2, O3, O4, O5, O6, O7, O8, T]) Map(fn func(T) O1) ApplyToResult7[O8, O7, O6, O5, O4, O3, O2, O1] {
	return Out7[O2, O3, O4, O5, O6, O7, O8](Map(a.result, fn))
}

type ApplyToResult9[O1, O2, O3, O4, O5, O6, O7, O8, O9, T any] struct {
	result Result[T]
}

func Out9[Out9, Out8, Out7, Out6, Out5, Out4, Out3, Out2, Out1, T any](result Result[T]) ApplyToResult9[Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9, T] {
	return ApplyToResult9[Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9, T]{
		result: result,
	}
}

func (a ApplyToResult9[O1, O2, O3, O4, O5, O6, O7, O8, O9, T]) AndThen(fn func(T) Result[O1]) ApplyToResult8[O9, O8, O7, O6, O5, O4, O3, O2, O1] {
	return Out8[O2, O3, O4, O5, O6, O7, O8, O9](AndThen(a.result, fn))
}

func (a ApplyToResult9[O1, O2, O3, O4, O5, O6, O7, O8, O9, T]) Map(fn func(T) O1) ApplyToResult8[O9, O8, O7, O6, O5, O4, O3, O2, O1] {
	return Out8[O2, O3, O4, O5, O6, O7, O8, O9](Map(a.result, fn))
}

type ApplyToResult10[O1, O2, O3, O4, O5, O6, O7, O8, O9, O10, T any] struct {
	result Result[T]
}

func Out10[Out10, Out9, Out8, Out7, Out6, Out5, Out4, Out3, Out2, Out1, T any](result Result[T]) ApplyToResult10[Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9, Out10, T] {
	return ApplyToResult10[Out1, Out2, Out3, Out4, Out5, Out6, Out7, Out8, Out9, Out10, T]{
		result: result,
	}
}

func (a ApplyToResult10[O1, O2, O3, O4, O5, O6, O7, O8, O9, O10, T]) AndThen(fn func(T) Result[O1]) ApplyToResult9[O10, O9, O8, O7, O6, O5, O4, O3, O2, O1] {
	return Out9[O2, O3, O4, O5, O6, O7, O8, O9, O10](AndThen(a.result, fn))
}

func (a ApplyToResult10[O1, O2, O3, O4, O5, O6, O7, O8, O9, O10, T]) Map(fn func(T) O1) ApplyToResult9[O10, O9, O8, O7, O6, O5, O4, O3, O2, O1] {
	return Out9[O2, O3, O4, O5, O6, O7, O8, O9, O10](Map(a.result, fn))
}
