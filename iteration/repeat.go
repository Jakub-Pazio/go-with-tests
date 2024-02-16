package iteration

import "strings"

func Repeat(s string, n int) string {
	var b strings.Builder
	for range n {
		b.WriteString(s)
	}
	return b.String()
}

func SlowRepeat(s string, n int) string {
	var result string
	for range n {
		result += s
	}
	return result
}
