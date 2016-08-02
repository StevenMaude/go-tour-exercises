package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	Urls map[string]bool
	mux  sync.Mutex
}

var Cached = Cache{Urls: make(map[string]bool)}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	Cached.mux.Lock()
	_, ok := Cached.Urls[url]
	Cached.mux.Unlock()

	if !ok {
		var wg sync.WaitGroup
		body, urls, err := fetcher.Fetch(url)

		// Debatable where you put this according to the provided spec.
		// Should it cache or retry on failure if the same url is
		// encountered?
		Cached.mux.Lock()
		Cached.Urls[url] = true
		Cached.mux.Unlock()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		// The approach used in the ideal solution to wait until each
		// Crawl is finished is to use channels. Each anonymous
		// goroutine Crawls as below, but then sends on the channel.
		// By looping over the same range urls, it's possible to then
		// wait for a done value to be sent over the channel, which
		// blocks until the same number have returned a done value.

		// Using WaitGroups is simple, but wasn't covered in the Tour.
		// Similar example here: https://golang.org/pkg/sync/#WaitGroup

		// Also note that the wg.Add needs to be outside of the
		// go func, as otherwise wg.Wait may be reached before any
		// wg.Add occurs, and the function instantly exits.
		for _, u := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				Crawl(url, depth-1, fetcher)
			}(u)
		}
		wg.Wait()
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
