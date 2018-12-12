// This is an utility to generate documentation from go-ipfs commands
package main

import (
	"fmt"
	"os"

	docs "github.com/ipfs/http-api-docs"
)

func main() {
	args := os.Args
	if len(args) < 2 || len(args) > 2 {
		fmt.Println("Usage: " + args[0] + " <formatter>")
		os.Exit(1)
	}

	arg := args[1]
	if arg == "markdown" {
		run(new(docs.MarkdownFormatter))
	} else if arg == "huskall" {
		run(new(docs.HuskallFormatter))
	} else if arg == "json" {
		run(new(docs.JsonFormatter))
	} else {
		fmt.Println("Unknown formatter: " + arg)
	}

}

func run(formatter docs.Formatter) {
	endpoints := docs.AllEndpoints()
	fmt.Println(docs.GenerateDocs(endpoints, formatter))
}
