package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type Token struct {
	TokenValue TokenType
	Next       *Token
	Value      string
}

type helper struct {
	isStr      bool
	isInnerKey bool
}

type YAMLLexer struct {
	src       string
	rootToken *Token
}

func NewLexer() *YAMLLexer {
	return &YAMLLexer{rootToken: &Token{}}
}

func (yamlLexer *YAMLLexer) SetString(str string) {
	yamlLexer.src = str
}

func (yamlLexer *YAMLLexer) Lex() error {
	root := Token{}
	curr := &root

	h := helper{isStr: false, isInnerKey: false}

	builder := &strings.Builder{}

	for _ /* TODO: replace it i(index) */, char := range yamlLexer.src {
		if char == '"' || char == '\'' {
			h.isStr = !h.isStr
			curr = emit(TokenQuote, "\"", curr)
		} else if !h.isStr && char == ':' {
			curr = emit(TokenKey, builder.String(), curr)
			builder.Reset()
		} else if !h.isStr && (char == '\n') {
			curr = emit(TokenValue, builder.String(), curr)
			builder.Reset()
		} else if !unicode.IsSpace(char) || h.isStr {
			builder.WriteRune(char)
		}
	}

	curr.Next = &Token{TokenValue: TokenEOF, Value: "[EOF]"}

	*yamlLexer.rootToken = *root.Next
	return nil
}

func emit(tokenVal TokenType, val string, curr *Token) *Token {
	curr.Next = &Token{TokenValue: tokenVal, Value: val}
	return curr.Next
}

func (yamlLexer *YAMLLexer) Print() {
	curr := yamlLexer.rootToken
	for curr != nil {
		fmt.Print(curr.Value, "-> ")
		curr = curr.Next
	}
	fmt.Println("<END>")
}
