# Go Result

Go result is a utility package for handling values that might succeed or fail. The types introduced
are heavily inspired by, and should feel familiar to their counterparts in Rust. There are some
minor deviations in this library which have been made to either suit the Go eco-system better, or
due to constraints in the Go language.

## The Result Type

`Result[T]` is a new type which can be used to express a value which can either be in an `Ok(T)`
state, or an `Err(error)` state. Because Go makes extensive use of wrapped errors, and already has
and error interface, the result type only takes one generic type (corresponding to the `Ok` state).
A result type will not indicate the exact type of a potential error.
