# Go HTML Link Parser

A simple HTML link parser that extracts all anchor tags (`<a>`) and their associated href attributes and text content from HTML documents.

## Installation

```bash
go get goHtmlLinkParser.bhavya.net
```

If you want to clone this git

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "strings"
    link "goHtmlLinkParser.bhavya.net"
)

exHTML = `
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
	fmt.Printf("%+v\n", links)
}
```

### Parse from File

```go
package main

import (
    "fmt"
    "os"
    link "goHtmlLinkParser.bhavya.net"
)

func main() {
    file, err := os.Open("ex1.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	links, err := link.Parse(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", links)
}
```

## Examples

The `examples/` directory contains several test cases:

- `ex1/`: Basic multiple links parsing
- `ex2/`: Complex nested HTML structures
- `ex3/`: Links with missing href attributes
- `ex4/`: Links with nested HTML elements

To run an example:

```bash
go run examples/ex1/main.go
```
