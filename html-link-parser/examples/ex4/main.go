package main

import (
	"fmt"
	link "my-go-code/html-link-parser"
	"strings"
)

var examplehtml = `<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
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
