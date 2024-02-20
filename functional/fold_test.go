package functional

import "testing"

func TestFold(t *testing.T) {
	array := []int{1, 2, 3}

	add := func(a, b int) int { return a + b }
	want := 6

	got := foldl(array, add, 0)

	if want != got {
		t.Errorf("expected %d but got %d", want, got)
	}
}
