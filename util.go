package t

func Cast[T any](value any) Option[T] {
	if t, ok := value.(T); ok {
		return Some(t)
	} else {
		return None[T]()
	}
}
