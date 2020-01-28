package main

import (
	"flag"
	"fmt"
	"io"
	link "my-go-code/html-link-parser"
	"net/http"
	"net/url"
	"strings"
)

/*
	1. GET the webpage
	2. parse all links on the page
	3. build proper urls with the links
	4. filter out any links w/ a diff domain
	5. find all pages (BFS, breadth first search)
	6. Print out XML
*/
/*
	Handle these types of URLS:
	/some-path
	https://gophercises.com/some-path

	SKIP these type:
	#fragment
	mailto:something@gmail.com
*/
func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	for _, page := range pages {
		fmt.Println(page)
	}
}

type empty struct{}

func bfs(urlStr string, maxDepth int) []string {
	// all the urls we have seen
	seen := make(map[string]empty)

	// queue
	var q map[string]empty
	// next queue
	nq := map[string]empty{
		urlStr: empty{},
	}

	for i := 0; i <= maxDepth; i++ {
		// q == nq -- and nq == make()
		q, nq = nq, make(map[string]empty)
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = empty{}
			for _, link := range get(url) {
				nq[link] = empty{}
			}
		}
	}

	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		// panic(err)
		return []string{}
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)
	var returnHrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			returnHrefs = append(returnHrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			returnHrefs = append(returnHrefs, l.Href)

		}
	}
	return returnHrefs
}

func filter(links []string, keepFunc func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFunc(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
