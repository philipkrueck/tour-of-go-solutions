// Video explanation: https://www.youtube.com/watch?v=2yUyMtXKWlk

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type URLCache struct {
	mu      sync.Mutex
	visited map[string]bool
}

// Visit marks a url as visited
// and returns whether it was already in the cache.
func (cache *URLCache) Visit(url string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	alreadyVisited := cache.visited[url]
	cache.visited[url] = true
	return alreadyVisited
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *URLCache, wg *sync.WaitGroup) {
	if depth <= 0 {
		return
	}
	if cache.Visit(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Go(func() {
			Crawl(u, depth-1, fetcher, cache, wg)
		})
	}
}

func main() {
	var wg sync.WaitGroup
	cache := URLCache{visited: make(map[string]bool)}

	Crawl("https://golang.org/", 4, fetcher, &cache, &wg)
	wg.Wait()
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
