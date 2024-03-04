package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// This problem in teory is resolved with go 1.22 but for now you can resolve it by declaring a new variable or adding a param to the anonymous func
		// First solution
		//url := url
		//go func() {
		//	result[url] = wc(url)
		//}()

		// Second solution
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)

	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
