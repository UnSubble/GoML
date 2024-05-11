package parser

import "os"

type MagicParser interface {
	Parse() (*YAMLNode, error)
}

type YAMLFileParser struct {
	path string
	root *YAMLNode
}

func NewYAMLFileParser(path string) MagicParser {
	return &YAMLFileParser{path: path}
}

func (yamlFileParser *YAMLFileParser) Parse() (*YAMLNode, error) {
	file, err := os.Open(yamlFileParser.path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	buffer := make([]byte, 1024)

	return yamlFileParser.root, nil
}
