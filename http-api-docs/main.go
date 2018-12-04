// This is an utility to generate documentation from go-ipfs commands
package main

import (
	"fmt"
	"github.com/ipfs/http-api-docs"
)

func main() {
	endpoints := docs.AllEndpoints()
	//formatter := new(docs.MarkdownFormatter)
	formatter := new(docs.HuskallFormatter)
	fmt.Println(docs.GenerateDocs(endpoints, formatter))
}
