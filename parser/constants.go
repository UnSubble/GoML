package parser

type ParserError struct {
	message string
}

func (e *ParserError) Error() string {
	return e.message
}
