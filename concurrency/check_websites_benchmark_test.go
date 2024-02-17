package concurrency

import (
	"testing"
	"time"
)

func stabCheckWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	size := 200
	urls := make([]string, size)
	for i := range size {
		urls[i] = "http://doesnotmatter.atall"
	}

	b.ResetTimer()
	for range b.N {
		CheckWebsites(stabCheckWebsite, urls)
	}
}

func BenchmarkFastCheckWebsites(b *testing.B) {
	size := 200
	urls := make([]string, size)
	for i := range size {
		urls[i] = "http://doesnotmatter.atall"
	}

	b.ResetTimer()
	for range b.N {
		FastCheckWebsites(stabCheckWebsite, urls)
	}
}
