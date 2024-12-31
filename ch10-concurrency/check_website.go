package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// Normally in Go when we call a function doSomething() we wait for it to return (even if it has no value to return, we still wait for it to finish).
// We say that this operation is blocking - it makes us wait for it to finish.
// An operation that does not block in Go will run in a separate process called a goroutine.

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
