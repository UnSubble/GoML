package parser

import (
	"io"
	"os"
)

type MagicParser interface {
	ParseNodes() (*YAMLNode, error)
}

type YAMLFileParser struct {
	path string
	root *YAMLNode
}

func NewYAMLFileParser(path string) MagicParser {
	return &YAMLFileParser{path: path}
}

func (yamlFileParser *YAMLFileParser) ParseNodes() (*YAMLNode, error) {
	file, err := os.Open(yamlFileParser.path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	buffer := make([]byte, 1024)
	read := 1

	for read != 0 {
		read, err = file.Read(buffer)

		if err != nil && err != io.EOF {
			return nil, err
		}

		// bufferedStr := string(buffer[:read])

		if err != nil {
			return nil, err
		}
	}

	return yamlFileParser.root, nil
}
