package main

import (
	"fmt"
	"os"

	link "goHtmlLinkParser.bhavya.net"
)

func main() {
	file, err := os.Open("ex3.html")
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
