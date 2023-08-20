package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	testCode := "let variable = 543 34 awayFrom 5"
	tokens := Tokenizer(testCode)
	PrintTokens(tokens)
}
