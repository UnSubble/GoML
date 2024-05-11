package parser

import (
	"errors"
	"strconv"
)

type YAMLNode struct {
	Key   string
	Next  *YAMLNode
	Child *YAMLNode
	value string
}

func (yamlNode *YAMLNode) HasNext() bool {
	return yamlNode.Next != nil
}

func (yamlNode *YAMLNode) HasChild() bool {
	return yamlNode.Child != nil
}

func (yamlNode *YAMLNode) GetValueAsString() (string, error) {
	if yamlNode.value == "" {
		return "", errors.New(string(nilValue))
	}

	return yamlNode.value, nil
}

func (yamlNode *YAMLNode) GetValueAsInt() (int, error) {
	if yamlNode.value == "" {
		return 0, errors.New(string(nilValue))
	}
	result, err := strconv.Atoi(yamlNode.value)

	if err != nil {
		return 0, err
	}

	return result, nil
}
