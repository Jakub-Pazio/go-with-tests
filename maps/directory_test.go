package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("word in dictionary", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		got, err := dict.Search("test")
		want := "this is just a test"

		assertNoError(t, err)
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
	t.Run("key does not exist yet", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("added", "new value")

		got, err := dict.Search("added")
		want := "new value"

		assertNoError(t, err)
		assertStrings(t, want, got)
	})

	t.Run("key is already in dictionary", func(t *testing.T) {
		dict := Dictionary{}
		key := "added"
		intialVal := "initial value"
		dict.Add(key, intialVal)
		err := dict.Add("added", "should not be added")

		want := ErrKeyAlreadyExists
		assertError(t, want, err)
		assertDefinition(t, dict, key, intialVal)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("positive update test", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		dict.Add(key, "first")

		newVal := "second"
		err := dict.Update("test", newVal)

		assertNoError(t, err)
		assertDefinition(t, dict, key, newVal)
	})
	t.Run("update fail, key does not exsist", func(t *testing.T) {
		dict := Dictionary{}
		key := "unpresent"

		want := ErrCantUpdate
		err := dict.Update(key, "does not matter")

		assertError(t, want, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("positive delete", func(t *testing.T) {
		dict := Dictionary{}
		key := "some"
		value := "value"
		dict.Add(key, value)

		assertDefinition(t, dict, key, value)

		dict.Delete(key)

		_, err := dict.Search(key)

		assertError(t, ErrKeyNotPresent, err)
	})
	t.Run("delete with key not in dictionary", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Delete("i dont exist")
		want := ErrNoDeleteKey

		assertError(t, want, err)
	})
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

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("expected to not error but did")
	}
}

// I dont like this helper method. It assumes how out dictionary is working
// It couples tests to Search method. I don't think I would write code like this myself
func assertDefinition(t testing.TB, d Dictionary, key, expected string) {
	t.Helper()
	val, err := d.Search(key)
	assertNoError(t, err)
	assertStrings(t, expected, val)
}
