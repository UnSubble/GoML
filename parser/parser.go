package parser

type MagicParser interface {
}

type YAMLFileParser struct {
	path string
	root *YAMLNode
}

func NewYAMLFileParser() *MagicParser {
	return nil
}

func (yamlFileParser *YAMLFileParser) Parse() *YAMLNode {
	// TODO
	return yamlFileParser.root
}
