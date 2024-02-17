package concurrency

import "net/http"

// Just check if you got any status code
func CheckWebsite(url string) bool {
	res, err := http.Get(url)
	if err != nil {
		return false
	}

	return res.StatusCode == 200
}
