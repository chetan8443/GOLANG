package main

import "sync"

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(7)

	// URL list for testing different websites
	urlList := []string{
		"https://google.com",
		"https://facebook.com",
		"https://github.com",
		"https://go.dev",
	}

	// Calling GoRoutine functions on list of URL
	for _, url := range urlList {
		go getStatusCode(url, wg)
	}

	// Checking Mutex function on different go routine functions
	routine1(wg, &mutex)
	routine2(wg, &mutex)
	routine3(wg, &mutex)

	wg.Wait()
}
