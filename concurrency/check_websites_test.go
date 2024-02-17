package concurrency

import (
	"reflect"
	"testing"
)

func mocWebsitesCheck(url string) bool {
	return url != "http://dourlslikethisevenexist.pl"
}

func TestCheckSites(t *testing.T) {
	urlRes := map[string]bool{
		"http://google.com":                 true,
		"http://archlinux.org/":             true,
		"http://dourlslikethisevenexist.pl": false,
	}

	sites := []string{
		"http://google.com",
		"http://archlinux.org/",
		"http://dourlslikethisevenexist.pl",
	}

	got := CheckWebsites(mocWebsitesCheck, sites)

	if !reflect.DeepEqual(urlRes, got) {
		t.Errorf("results are not correct\ngot: %v\n wanted: %v", urlRes, got)
	}
}
