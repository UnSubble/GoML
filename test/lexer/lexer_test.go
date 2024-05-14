package lexer_test

import (
	"os"
	"testing"

	"github.com/unsubble/goml/lexer"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ExpectedTokenShouldEqualToActualToken(t *testing.T) {
	l := lexer.NewLexer()
	l.SetString("\"test 1\":\n\ttset")
	l.Lex()
	l.Print()
}
