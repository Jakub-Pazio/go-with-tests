package concurrency

type WebSiteChecker func(url string) bool

func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	resMap := make(map[string]bool)

	for _, url := range urls {
		ok := wc(url)
		resMap[url] = ok
	}

	return resMap
}

func FastCheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	resMap := make(map[string]bool)
	c := make(chan struct {
		string
		bool
	})

	for _, url := range urls {
		go func(url string) {
			res := wc(url)
			c <- struct {
				string
				bool
			}{url, res}
		}(url)
	}

	for range len(urls) {
		res := <-c
		resMap[res.string] = res.bool
	}

	return resMap
}
