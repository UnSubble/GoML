package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type Token struct {
	Value      string
	TokenValue TokenType
	Next       *Token
}

type YAMLLexer struct {
	Src       string
	RootToken *Token
	lastToken *Token
	lastChar  rune
	isStr     bool
}

func NewLexer() *YAMLLexer {
	return &YAMLLexer{RootToken: &Token{TokenValue: TokenEOF, Value: "[ROOT]"}}
}

func (l *YAMLLexer) Lex() error {
	if l.lastToken == nil {
		l.lastToken = l.RootToken
	}

	builder := &strings.Builder{}

	for _, char := range l.Src {
		l.updateStrMode(char)
		if l.shouldEmitToken(char) {
			if builder.Len() > 0 {
				l.lastToken = emitToken(l.lastToken, builder.String())
			}
			l.lastToken = emitToken(l.lastToken, string(char))
			builder.Reset()
		} else if char == '\n' && !l.isStr && builder.Len() > 0 {
			l.lastToken = emitToken(l.lastToken, builder.String())
			builder.Reset()
		} else if !unicode.IsSpace(char) || l.isStr {
			builder.WriteRune(char)
		}
	}

	if builder.Len() > 0 {
		l.lastToken = emitToken(l.lastToken, builder.String())
	}

	tokenizeAndEmit(l.lastToken, TokenEOF, "[EOF]")

	return nil
}

func emitToken(curr *Token, val string) *Token {
	switch val {
	case ":":
		return tokenizeAndEmit(curr, TokenColon, "[COLON]")
	case "-":
		return tokenizeAndEmit(curr, TokenDash, "[DASH]")
	case " ", "\r":
		return tokenizeAndEmit(curr, TokenSpace, "[SPACE]")
	case "\t":
		return tokenizeAndEmitSpace(curr, 4)
	default:
		return tokenizeAndEmit(curr, TokenIdentifier, val)
	}
}

func tokenizeAndEmit(curr *Token, tokenVal TokenType, val string) *Token {
	curr.Next = &Token{TokenValue: tokenVal, Value: val}
	return curr.Next
}

func tokenizeAndEmitSpace(curr *Token, count int) *Token {
	for i := 0; i < count; i++ {
		curr = tokenizeAndEmit(curr, TokenSpace, "[SPACE]")
	}
	return curr
}

func (l *YAMLLexer) shouldEmitToken(char rune) bool {
	return (char == ':' || char == '-' || char == '\t' || char == '\r' || char == ' ') && !l.isStr
}

func (l *YAMLLexer) updateStrMode(char rune) {
	if l.lastChar == 0 && (char == '"' || char == '\'') {
		l.lastChar = char
		l.isStr = true
	} else if l.lastChar == char {
		l.isStr = false
		l.lastChar = 0
	}
}

// TODO: remove this later
func (yamlLexer *YAMLLexer) Print() {
	curr := yamlLexer.RootToken
	for curr != nil {
		fmt.Print(curr.Value, "(", curr.TokenValue, ")", "-> ")
		curr = curr.Next
	}
	fmt.Println("<END>")
}
