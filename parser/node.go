package parser

type YAMLNode struct {
	entryCount int
	Key        string
	Value      *string
	Next       *YAMLNode
	Child      *YAMLNode
}

func (yamlNode *YAMLNode) HasNext() bool {
	return yamlNode.Next != nil
}

func (yamlNode *YAMLNode) HasChild() bool {
	return yamlNode.Child != nil
}
