package main

import (
	"fmt"
	"strings"

	link "goHtmlLinkParser.bhavya.net"
)

var exHTML = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exHTML)

	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", links)
}
