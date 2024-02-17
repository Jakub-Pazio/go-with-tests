package concurrency

import (
	"testing"
)

func TestCheckWebsite(t *testing.T) {
	got := CheckWebsite("http://google.com")

	if got != true {
		t.Errorf("Either my code is wrong or google.com does not work... (or both :P)")
	}
}
