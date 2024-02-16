package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Adder(2, 2)
	want := 4

	if got != want {
		t.Errorf("wanted %d but got %d", want, got)
	}
}

func ExampleAdder() {
	sum := Adder(1000, 337)
	fmt.Println(sum)
	// Output: 1337
}
