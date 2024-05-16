package main

import (
	"github.com/unsubble/goml/parser"
)

func main() {
	p := parser.NewYAMLFileParser("test.yaml")
	node, err := p.ParseFile()
	if err != nil {
		panic(err)
	}
	parser.Print(node, "")
}
