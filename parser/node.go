package parser

import (
	"strings"
)

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

func (yamlNode *YAMLNode) GetValueAsString() string {
	if yamlNode.Value != nil {
		return *yamlNode.Value
	}
	return ""
}

func (yamlNode *YAMLNode) GetValueAsSlice() []interface{} {
	if yamlNode.Value != nil {
		var slice []interface{}
		builder := &strings.Builder{}
		lastQuote := ' '
		isStr := false

		for _, val := range *yamlNode.Value {
			if val == lastQuote || lastQuote == ' ' {
				isStr = !isStr
				lastQuote = val
				if isStr {
					lastQuote = ' '
				}
			}
			if val == ' ' && !isStr {
				slice = append(slice, builder.String())
				builder.Reset()
			} else {
				builder.WriteRune(val)
			}

		}
		return slice
	}
	return nil
}
