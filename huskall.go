package docs

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var ArgTypes = map[string]string{
	"bool":   "Bool",
	"file":   "ByteString",
	"int":    "Int",
	"string": "Text",
	"uint":   "Word",
}

type HuskallFormatter struct{}

var endpointName = regexp.MustCompile("[-/][A-Za-z0-9]")

func (hs *HuskallFormatter) GenerateIntro() string {
	buf := new(bytes.Buffer)
	fmt.Fprintln(buf, "{-# LANGUAGE DataKinds #-}")
	fmt.Fprintln(buf, "{-# LANGUAGE TypeOperators #-}")
	fmt.Fprintln(buf, "module IPFS.API where")
	fmt.Fprintln(buf)
	fmt.Fprintln(buf, "import Data.Aeson (Value)")
	fmt.Fprintln(buf, "import Data.ByteString (ByteString)")
	fmt.Fprintln(buf, "import Data.Text (Text)")
	fmt.Fprintln(buf, "import Servant.API")
	fmt.Fprintln(buf, "import Servant.Multipart")
	fmt.Fprintln(buf)

	return buf.String()
}

func (hs *HuskallFormatter) GenerateEndpointBlock(endp *Endpoint) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "type %s = ", endpointName.ReplaceAllStringFunc(endp.Name, func(s string) string {
		return strings.ToTitle(strings.Replace(strings.Replace(s, "/", "", -1), "-", "", -1))
	}))

	for _, frag := range strings.Split(endp.Name, "/") {
		if frag != "" {
			fmt.Fprintf(buf, "\"%s\" :> ", frag)
		}
	}

	var haveMultipart = false

	for _, arg := range endp.Arguments {
		if arg.Type == "file" {
			haveMultipart = true
			fmt.Fprintf(buf, "MultipartForm Mem (MultipartData Mem) :> ")
		} else {
			fmt.Fprintf(buf, "QueryParam' '[Required, Strict] \"%s\" %s :> ", arg.Name, ArgTypes[arg.Type])
		}
	}

	for _, opt := range endp.Options {
		if opt.Type == "file" {
			haveMultipart = true
			fmt.Fprintf(buf, "MultipartForm Mem (MultipartData Mem) :> ")
		} else {
			fmt.Fprintf(buf, "QueryParam \"%s\" %s :> ", opt.Name, ArgTypes[opt.Type])
		}
	}

	if haveMultipart {
		fmt.Fprintf(buf, "Stream 'POST 200 NoFraming PlainText ByteString")
	} else {
		fmt.Fprintf(buf, "Get '[JSON] Value")
	}

	fmt.Fprintln(buf)
	fmt.Fprintln(buf)

	return buf.String()
}

// Unneeded stuff

func (hs *HuskallFormatter) GenerateArgumentsBlock(reqs []*Argument, opts []*Argument) string {
	return ""
}

func (hs *HuskallFormatter) GenerateIndex(enps []*Endpoint) string {
	return ""
}

func (hs *HuskallFormatter) GenerateExampleBlock(endp *Endpoint) string {
	return ""
}

func (hs *HuskallFormatter) GenerateBodyBlock(args []*Argument) string {
	return ""
}

func (hs *HuskallFormatter) GenerateResponseBlock(response string) string {
	return ""
}
