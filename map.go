package t

// the identity map
func Id[T any](t T) T { return t }

func Return[In any, T any](t T) func(In) T { return func(_ In) T { return t } }

func Compose[T, U, V any](fn1 func(T) U, fn2 func(U) V) func(T) V {
	return func(t T) V { return fn2(fn1(t)) }
}
