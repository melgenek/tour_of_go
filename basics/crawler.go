package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type work struct {
	url   string
	depth int
}

type result struct {
	urls  []string
	depth int
}

func worker(in chan work, res chan result, fetcher Fetcher) {
	for work := range in {
		body, urls, err := fetcher.Fetch(work.url)
		if err != nil {
			fmt.Println(err)
			res <- result{[]string{}, work.depth}
		} else {
			fmt.Printf("found: %d %s %q\n", work.depth, work.url, body)
			res <- result{urls, work.depth}
		}
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	in := make(chan work)
	res := make(chan result)

	for w := 1; w <= 3; w++ {
		go worker(in, res, fetcher)
	}

	in <- work{url, 0}
	total := 1
	cache := make(map[string]bool)
	cache[url] = true
	for result := range res {
		if result.depth < depth {
			for _, url := range result.urls {
				if !cache[url] {
					total++
					cache[url] = true
					go func(u string) { in <- work{u, result.depth + 1} }(url)
				}
			}
		}
		total--
		if total == 0 {
			break
		}
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
