package t_test

import (
	"fmt"
	"testing"

	. "aidanwoods.dev/go-result"
	"github.com/stretchr/testify/require"
)

func divideBy3(i int) Result[int] {
	if i%3 == 0 {
		return Ok(i / 3)
	} else {
		return Err[int](fmt.Errorf("number does not divide by 3"))
	}
}

func foo(i int) Result[int] {
	var divided int
	if err := divideBy3(i).Ok(&divided); err != nil {
		return Err[int](err)
	}

	if divided == 2 {
		return Ok(2)
	}

	return foo(divided + 9)
}

func TestResultExample(t *testing.T) {
	require.True(t, foo(-9).IsOk())
	require.True(t, foo(6).IsOk())
	require.True(t, foo(-54).IsOk())
}
