package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	urlFlag := flag.String("url", "http://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	/*
		1. GET the webpage
		2. parse all links on the page
		3. build proper urls with the links
		4. filter out any links w/ a diff domain
		5. find all pages (BFS, breadth first search)
		6. Print out XML
	*/
}
