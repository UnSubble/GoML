package parser

import (
	"fmt"
	"io"
	"os"

	"github.com/unsubble/goml/lexer"
)

type MagicParser interface {
	ParseFile() (*YAMLNode, error)
}

type YAMLFileParser struct {
	path string
	root *YAMLNode
}

func NewYAMLFileParser(path string) MagicParser {
	return &YAMLFileParser{path: path}
}

func (yamlFileParser *YAMLFileParser) ParseFile() (*YAMLNode, error) {
	file, err := os.Open(yamlFileParser.path)

	if err != nil {
		return nil, err
	}

	l := lexer.NewLexer()

	defer file.Close()

	buffer := make([]byte, 1024)
	read := 1

	for read != 0 {
		read, err = file.Read(buffer)

		if err != nil && err != io.EOF {
			return nil, err
		}

		bufferedStr := string(buffer[:read])
		l.Src = bufferedStr
		l.Lex()
	}

	l.Print()

	yamlFileParser.root = ParseTokens(l.RootToken)

	return yamlFileParser.root, nil
}

func ParseTokens(token *lexer.Token) *YAMLNode {
	entryMap := make(map[int]*YAMLNode)

	var spaceCount int
	var last *YAMLNode = nil

	root := &YAMLNode{Key: "[ROOT]", entryCount: -1}
	entryMap[spaceCount] = root

	for token != nil {
		switch token.TokenValue {
		case lexer.TokenSpace:
			spaceCount++
		case lexer.TokenIdentifier:
			value := token.Value
			if token.Next.TokenValue == lexer.TokenIdentifier {
				spaceCount = 0
			}
			for token.Next.TokenValue == lexer.TokenIdentifier {
				token = token.Next
				if last.Value != nil {
					*last.Value = *last.Value + " " + value
				} else {
					last.Value = &value
				}
				value = token.Value
			}
			node := &YAMLNode{Key: value, entryCount: spaceCount}
			if token.Next.TokenValue == lexer.TokenColon {
				handleNode(entryMap, spaceCount, node, &last)
			} else {
				if last.Value != nil {
					*last.Value += node.Key
				} else {
					last.Value = &node.Key
				}
			}
			spaceCount = 0
		}

		token = token.Next
	}

	return root.Next
}

func handleNode(entryMap map[int]*YAMLNode, spaceCount int, node *YAMLNode, last **YAMLNode) {
	if spaceCount <= getParent(entryMap, spaceCount).entryCount {
		getParent(entryMap, spaceCount).Next = node
		// if spaceCount == 0 {
		clear(entryMap)
		// }
	} else if entryMap[spaceCount] != nil {
		entryMap[spaceCount].Next = node
	} else {
		getParent(entryMap, spaceCount).Child = node
	}
	entryMap[spaceCount] = node
	*last = node
}

func getParent(entryMap map[int]*YAMLNode, spaceCount int) *YAMLNode {
	for i := spaceCount - 1; i >= 0; i-- {
		if entryMap[i] != nil {
			return entryMap[i]
		}
	}
	return entryMap[spaceCount]
}

func clear(entryMap map[int]*YAMLNode) {
	for i := 0; i < len(entryMap); i++ {
		entryMap[i] = nil
	}
}

func Print(yamlNode *YAMLNode, space string) {
	root := yamlNode

	for root != nil {
		fmt.Print(space, root.Key)
		if root.Value != nil {
			fmt.Println(": { VALUE=", *root.Value, "}")
		} else {
			fmt.Println(":")
		}
		child := root.Child
		Print(child, space+"  ")
		root = root.Next
	}
}
