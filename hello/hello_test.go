package hello

import (
	"bytes"
	"testing"
)

func TestPrintHelloWorld(t *testing.T) {
	cases := []struct {
		testName string
		name     string
		lang     Lang
		want     string
	}{
		{"Polish greeting", "Jakub", polPrefix, "Cześć, Jakub!"},
		{"Spanish greeting", "Tod", spaPrefix, "Ola, Tod!"},
		{"English greeting", "Ed", engPrefix, "Hello, Ed!"},
		{"Handle empty name english", "", engPrefix, "Hello"},
		{"Handle empty name polish", "", polPrefix, "Cześć"},
	}

	for _, test := range cases {
		buf := bytes.Buffer{}
		t.Run(test.testName, func(t *testing.T) {
			greeter(&buf, test.lang, test.name) // This we are testing
			assertCorrectMessage(test.want, buf.String(), t)
		})
	}
}

func assertCorrectMessage(want, got string, t testing.TB) {
	if got != want {
		t.Errorf("wanted %s got %s", want, got)
	}
}
