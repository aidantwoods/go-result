package t

type Void struct{}

var void = Void{}

type Tuple[T, U any] struct {
	First  T
	Second U
}

func newTuple[T, U any](first T, second U) Tuple[T, U] {
	return Tuple[T, U]{
		First:  first,
		Second: second,
	}
}
