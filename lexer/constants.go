package lexer

type TokenType int

const (
	TokenColon TokenType = iota
	TokenDash
	TokenSpace
	TokenIdentifier
	TokenEOF
)
