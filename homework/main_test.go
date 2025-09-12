package main_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avito-edu/goms-2025-practice-1/homework"
)

func TestAdd(t *testing.T) {
	got := main.Add(10, 5)
	want := 15.0
	require.Equal(t, got, want)
}

func TestDivide(t *testing.T) {
	got := main.Divide(10, 5)
	want := 2.0
	require.Equal(t, got, want)
}

func TestDivideByZero(t *testing.T) {
	require.Panics(
		t,
		func() { main.Divide(10, 0) },
		"Ожидалась паника при делении на ноль",
	)
}

func TestSubtract(t *testing.T) {
	got := main.Subtract(10, 5)
	want := 5.0
	require.Equal(t, got, want)
}

func TestMultiply(t *testing.T) {
	got := main.Multiply(10, 5)
	want := 50.0
	require.Equal(t, got, want)
}
