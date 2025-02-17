package types

// The identity map
func Id[T any](t T) T { return t }

// Return the given value, given any input of type In
func Return[In any, T any](t T) func(In) T { return func(_ In) T { return t } }

// Return the given value, given any input of type In
func Return0[T any](t T) func() T { return func() T { return t } }

func Vaule[T any]() (t T) { return t }

func Compose[T, U, V any](fn1 func(T) U, fn2 func(U) V) func(T) V {
	return func(t T) V { return fn2(fn1(t)) }
}
