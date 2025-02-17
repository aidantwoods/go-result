# Go Result

Go result is a utility package for handling values that might succeed or fail. The types introduced
are heavily inspired by, and should feel familiar to their counterparts in Rust. There are some
minor deviations in this library which have been made to either suit the Go eco-system better, or
due to constraints in the Go language.

## The Option Type
`option.Option[T]` is a type which can be used to express a value or absence of a value. Use of this type is distinct mostly distinct from a pointer `*T` due to the absence of shared mutability.

## The Result Type
`result.Result[T]` is a new type which can be used to express a value which can either be in an `result.Ok(T)`
state, or an `result.Err(error)` state. Because Go makes extensive use of wrapped errors, and already has
and error interface, the result type only takes one generic type (corresponding to the `Ok` state).
A result type will not indicate the exact type of a potential error.

The result type cannot simultaneously have no `Ok` value and no error. If the
result type is explicitly instantiated with `result.Err(nil)` or is instantiated as a
zero value of `result.Result[T]`, then `result.Result[T]` is defined to be in an error state,
and will have a value of `result.ErrEmptyResult` when queried.
