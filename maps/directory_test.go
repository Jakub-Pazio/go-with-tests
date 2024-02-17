package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("word in dictionary", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, want, got)
	})

	t.Run("word in not dictionary", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		_, err := dict.Search("meNoHere")
		want := ErrKeyNotPresent

		assertError(t, want, err)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	dict.Add("added", "new value")

	got, _ := dict.Search("added")
	want := "new value"

	assertStrings(t, want, got)
}

func assertStrings(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("expected %s but got %s", want, got)
	}
}

func assertError(t testing.TB, want, got error) {
	t.Helper()
	if got == nil {
		t.Error("expected to throw an error but did not")
	}
	if want.Error() != got.Error() {
		t.Errorf("expected %s but got %s", want, got)
	}
}
