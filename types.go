package t

type Void struct{}

var void = Void{}

type Tuple[T, U any] struct {
	First  T
	Second U
}

func (t Tuple[T, U]) Destructure() (T, U) {
	return t.First, t.Second
}

func newTuple[T, U any](first T, second U) Tuple[T, U] {
	return Tuple[T, U]{
		First:  first,
		Second: second,
	}
}
