package main

import "github.com/unsubble/goml/parser"

func main() {
	parser := parser.NewYAMLFileParser("")
	parser.Parse()
}
