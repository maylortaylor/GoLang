package main

import (
	"fmt"
	 "my-go-code/html-link-parser"
	"strings"
)


var examplehtml = `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
  <a href="/other-page2">
    <span> some span </span>
  </a>
</body>
</html>`

func main() {
	r := strings.NewReader(examplehtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
