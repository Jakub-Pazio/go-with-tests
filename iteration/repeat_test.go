package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkBigRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 1000)
	}
}

func BenchmarkBigSlowRepeat(b *testing.B) {
	for range b.N {
		SlowRepeat("a", 1000)
	}
}

func BenchmarkSmallRepeat(b *testing.B) {
	for range b.N {
		Repeat("a", 2)
	}
}

func BenchmarkSmallSlowRepeat(b *testing.B) {
	for range b.N {
		SlowRepeat("a", 2)
	}
}

func TestRepeat(t *testing.T) {
	cases := []struct {
		str    string
		number int
		want   string
	}{
		{"a", 1, "a"},
		{"a", 2, "aa"},
		{"a", 0, ""},
		{"a", -3, ""},
		{"b", 2, "bb"},
		{"works", 2, "worksworks"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s %d times", test.str, test.number), func(t *testing.T) {
			got := Repeat(test.str, test.number)
			assertCorrectString(t, test.want, got)
		})
	}
}

func assertCorrectString(t testing.TB, want, got string) {
	if want != got {
		t.Errorf("expected %q but got %q", want, got)
	}
}
