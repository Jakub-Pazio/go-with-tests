package generics

import "testing"

func assertEqual[T comparable](t testing.TB, a, b T) {
	t.Helper()
	if a != b {
		t.Error("value are not equal")
	}
}

func assertNotEqual[T comparable](t testing.TB, a, b T) {
	t.Helper()
	if a == b {
		t.Error("value are equal")
	}
}
