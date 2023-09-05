package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	channel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// "send statment"
			channel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// "Receive expression"
		r := <-channel
		results[r.string] = r.bool
	}

	return results
}
