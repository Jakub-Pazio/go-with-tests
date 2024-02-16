package hello

import (
	"fmt"
	"io"
	"os"
)

type Lang string

const (
	engPrefix Lang = "Hello, "
	spaPrefix Lang = "Ola, "
	polPrefix Lang = "Cześć, "
)

// This is the way how i would do it in some library
// To not export internal function (lower-case name)
// And then just expose the version that prints to stdout
// In C and other languages you could do names like: _privatef() and publicf()
func greeter(w io.Writer, lang Lang, s string) {
	if s == "" {
		newHello := lang[:len(lang)-2]
		fmt.Fprintf(w, "%s", newHello)
	} else {
		fmt.Fprintf(w, "%s%s!", lang, s)
	}
}

func Greeter(lang Lang, s string) {
	greeter(os.Stdout, lang, s)
}
