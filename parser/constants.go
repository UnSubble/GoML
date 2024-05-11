package parser

type parserError string

const (
	nilValue parserError = "the node does not have any values"
)
