package lexer

type TokenType int

const (
	TokenColon TokenType = iota
	TokenQuote
	TokenDash
	TokenSpace
	TokenValue
	TokenKey
	TokenEOF
)
