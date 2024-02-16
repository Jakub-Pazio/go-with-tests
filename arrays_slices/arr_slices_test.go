package arraysslices

import (
	"slices"
	"testing"
)

func TestSumArray(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	// slice := arr[:] creates slice from an array
	// arr := []int{1, 2, 3, 4, 5} would be a SLICE initialization

	got := SumArray(arr)

	want := 15
	if want != got {
		t.Errorf("exprected %d but got %d", want, got)
	}
}

func TestSumSlice(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := SumSlice(numbers)

	want := 15
	if want != got {
		t.Errorf("exprected %d but got %d", want, got)
	}
}

func TestSumAll(t *testing.T) {
	arr := [][]int{{1, 2}, {3, 4}}
	got := SumAll(arr...) // Destructs array (of arrays) to its elements

	want := []int{3, 7}
	if !compareSlices(got, want) {
		t.Errorf("exprected %v but got %v", want, got)
	}
}

func TestSumTails(t *testing.T) {
	arr := [][]int{{1, 2, 5}, {10}, {3, 4, 9}, {}}
	got := SumAllTails(arr...)

	want := []int{7, 0, 13, 0}
	// Now we can compare slices with function from standar library
	// https://pkg.go.dev/slices#Equal
	// https://stackoverflow.com/a/71349708
	if !slices.Equal(got, want) {
		t.Errorf("exprected %v but got %v", want, got)
	}
}

// Not using reflect.DeepEqual for a reason
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range len(a) {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
